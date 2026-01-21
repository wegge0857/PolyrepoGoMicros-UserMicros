package data

import (
	"context"
	"database/sql"
	"userMicros/internal/biz"

	"github.com/go-kratos/kratos/v2/log"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"github.com/dtm-labs/client/dtmcli/dtmimp"
	"github.com/dtm-labs/client/dtmgrpc"
)

type UserRepo struct {
	data *Data //  数据源句柄: 指向 internal/data/data.go 中定义的 Data 结构体的指针
	log  *log.Helper
}

// 注入阶段（启动时）： wire 运行 NewUserRepo。
// 生成的对象被塞进biz层的接口里，最终通过 NewUserUseCase 注入给 uc.repo。
// Wire 把 Data 层的实例“注入”到了 Biz 层 （依赖注入）
func NewUserRepo(data *Data, logger log.Logger) biz.UserRepo { // 这里的参数由Wire去传入
	return &UserRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

// 调用阶段（请求时）： 当用户访问接口，Service 调用 uc.Get(ctx, id)。
// 在 uc内部，uc.repo.FindByID(ctx, id) ---> uc.repo 实际上是一个接口 触发 已实现的 FindByID方法
func (r *UserRepo) FindByID(ctx context.Context, id int64) (*biz.User, error) {
	var u User
	db := r.data.db.WithContext(ctx) // WithContext(ctx) 确保遵循链路追踪和超时控制
	// 使用 GORM 的 First 方法按主键查询
	if err := db.First(&u, id).Error; err != nil {
		return nil, gorm.ErrRecordNotFound
	}

	// 将 Data 层的模型转换为 Biz 层的业务实体
	return &biz.User{
		Id:    u.ID,
		Name:  u.Username,
		Email: u.Email,
	}, nil
}

// 记录用户收藏，调用dtm的屏障
func (r *UserRepo) UserStarRecord(ctx context.Context, userId int64, etfId int64) error {

	log.Info("--------->begin")

	barrier, err := dtmgrpc.BarrierFromGrpc(ctx)
	if err != nil {
		return err
	}

	// 特殊情况 判断是否是回退逻辑
	if barrier.Op == dtmimp.OpCompensate || barrier.Op == dtmimp.OpRollback {
	}

	// 正常加的逻辑
	sqlDB, err := r.data.db.DB()
	if err != nil {
		return err
	}
	log.Infof("--------->UserRepo.UserStarRecord.userId: %d, etfId: %d", userId, etfId)

	err = barrier.CallWithDB(sqlDB, func(sTx *sql.Tx) error {
		// 1. 【关键】将原生 sTx 包装进 GORM
		// 这样 gdb 就是一个已经开启了事务、且使用了 DTM 屏障连接的 GORM 对象
		gdb, err := gorm.Open(
			mysql.New(mysql.Config{Conn: sTx}),
			&gorm.Config{})
		if err != nil {
			return err
		}

		// 使用 ctx 保证链路追踪
		tx := gdb.WithContext(ctx)

		// 2. 准备 SQL 表达式
		record := UserStarRecord{
			UserId: userId,
			EtfId:  etfId,
		}

		// 3. 执行业务逻辑 (完全使用 tx 对象)
		// 添加 star 记录
		return tx.Create(&record).Error

	})
	if err != nil {
		return err
	}

	return nil
}
