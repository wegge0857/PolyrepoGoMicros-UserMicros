package biz_test

import (
	"context"
	"errors"
	"testing"

	"userMicros/internal/biz"
	"userMicros/internal/biz/mocks" // 假设你的 Mock 在这里

	"github.com/go-kratos/kratos/v2/log"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestUserUseCase_Get(t *testing.T) {
	// 这里的 MockUserRepo 是 mockery 自动生成的类名
	mockRepo := new(mocks.UserRepo)
	logger := log.DefaultLogger

	userId := int64(100)
	expectedUser := &biz.User{
		Id:    userId,
		Name:  "Gemini",
		Email: "gemini@example.com",
	}

	// 此时类型必然匹配，因为都在同一个包内
	mockRepo.On("FindByID", mock.Anything, userId).Return(expectedUser, nil)

	uc := biz.NewUserUseCase(mockRepo, logger)

	res, err := uc.Get(context.Background(), userId)

	assert.NoError(t, err)
	assert.Equal(t, expectedUser.Name, res.Name)
	mockRepo.AssertExpectations(t)
}

func TestUserUseCase_Get_Error(t *testing.T) {
	// 1. 注意：类名 mocks.UserRepo
	mockRepo := new(mocks.UserRepo)

	// 2. 注入 Mock 对象
	uc := biz.NewUserUseCase(mockRepo, log.DefaultLogger)

	userId := int64(404)

	// 3. 模拟数据库查询不到报错
	// 注意：Return 的第一个参数必须匹配接口定义的 *User 类型（即使是 nil）
	// Return( (*User)(nil), ... ) 这种写法可以防止某些情况下的类型模糊
	mockRepo.On("FindByID", mock.Anything, userId).Return((*biz.User)(nil), errors.New("not found"))

	// 执行
	res, err := uc.Get(context.Background(), userId)

	// 断言
	assert.Error(t, err)
	assert.Nil(t, res) // 建议使用 assert.Nil(t, res) 传入 t
	assert.Equal(t, "not found", err.Error())

	// 验证调用
	mockRepo.AssertExpectations(t)
}
