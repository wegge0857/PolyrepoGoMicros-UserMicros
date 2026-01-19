package data

import (
	"time"
	"userMicros/internal/conf"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewUserRepo)

type Data struct {
	db *gorm.DB // 添加 GORM 句柄
}

type User struct {
	ID        int64 `gorm:"primarykey"`
	Username  string
	Email     string
	CreatedAt time.Time
	UpdatedAt time.Time
}

// NewData 负责初始化数据库连接
func NewData(c *conf.Data, logger log.Logger) (*Data, func(), error) {
	// 这里使用 gorm.Open 连接数据库
	db, err := gorm.Open(mysql.Open(c.Database.Source), &gorm.Config{})
	if err != nil {
		log.NewHelper(logger).Error("userMicros 数据库连接失败...")
		return nil, nil, err
	}

	// 自动迁移表结构（可选，方便测试）
	db.AutoMigrate(&User{})
	log.NewHelper(logger).Info("userMicros 数据库连接成功...")
	cleanup := func() {
		log.NewHelper(logger).Info("userMicros 关闭mysql数据资源...") //Ctrl + C 或 Kill
		// 如果需要显式关闭 DB，可以在这里获取 sql.DB 并 Close
	}
	return &Data{db: db}, cleanup, nil
}
