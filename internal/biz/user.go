package biz

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
)

// User 是业务领域对象
type User struct {
	Id    int64
	Name  string
	Email string
}

// Data 层是 Biz 层接口的具体实现者 （依赖倒置，而非通常的业务依赖数据库）
// 实现逻辑：在data层 返回格式 必须为biz.UserRepo
type UserRepo interface {
	FindByID(ctx context.Context, id int64) (*User, error)
	UserStarRecord(ctx context.Context, userId int64, etfId int64) error
}

// UserUseCase 是业务逻辑主体
type UserUseCase struct {
	repo UserRepo
	log  *log.Helper
}

func NewUserUseCase(repo UserRepo, logger log.Logger) *UserUseCase {
	return &UserUseCase{repo: repo, log: log.NewHelper(logger)}
}

// Get 获取用户业务逻辑
func (uc *UserUseCase) Get(ctx context.Context, id int64) (*User, error) {
	return uc.repo.FindByID(ctx, id)
}

// UserStarRecord func
func (uc *UserUseCase) UserStarRecord(ctx context.Context, userId int64, etfId int64) error {
	return uc.repo.UserStarRecord(ctx, userId, etfId)
}
