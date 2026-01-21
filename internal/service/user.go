package service

import (
	"context"
	"userMicros/internal/biz"

	userV1 "github.com/wegge0857/PolyrepoGoMicros-ApiLink/user/v1"
)

type UserService struct {
	// 当前你在 proto 里定义的接口，如果你没写，就会自动走 UnimplementedUserServer 的逻辑，就不会报错
	userV1.UnimplementedUserServer
	// 类型嵌入（Type Embedding） 或 匿名组合
	// 即UserService 自动获得了 UnimplementedUserServer 的所有方法。UserService实现了UserServer接口的所有方法

	// 添加这一行，以便在方法中使用 s.uc
	uc *biz.UserUseCase
}

// 在参数中加入 uc *biz.UserUseCase
func NewUserService(uc *biz.UserUseCase) *UserService {
	return &UserService{
		uc: uc,
	}
}

// GetUser
func (s *UserService) GetUser(ctx context.Context, req *userV1.GetUserRequest) (*userV1.GetUserReply, error) {
	// 调用 biz 层
	u, err := s.uc.Get(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	return &userV1.GetUserReply{
		Id:    u.Id,
		Name:  u.Name,
		Email: u.Email,
	}, nil
}

// UserStarRecord
func (s *UserService) UserStarRecord(ctx context.Context, req *userV1.UserStarRecordRequest) (*userV1.UserStarRecordReply, error) {
	// 调用 biz 层
	err := s.uc.UserStarRecord(ctx, req.UserId, req.EtfId)
	if err != nil {
		return nil, err
	}
	return &userV1.UserStarRecordReply{}, nil
}

func (s *UserService) CreateUser(ctx context.Context, req *userV1.CreateUserRequest) (*userV1.CreateUserReply, error) {
	return &userV1.CreateUserReply{}, nil
}
func (s *UserService) UpdateUser(ctx context.Context, req *userV1.UpdateUserRequest) (*userV1.UpdateUserReply, error) {
	return &userV1.UpdateUserReply{}, nil
}
func (s *UserService) DeleteUser(ctx context.Context, req *userV1.DeleteUserRequest) (*userV1.DeleteUserReply, error) {
	return &userV1.DeleteUserReply{}, nil
}
func (s *UserService) ListUser(ctx context.Context, req *userV1.ListUserRequest) (*userV1.ListUserReply, error) {
	return &userV1.ListUserReply{}, nil
}
