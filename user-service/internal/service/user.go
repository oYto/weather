package service

import (
	"context"
	"errors"
	"fmt"
	"log"
	"user-service/internal/model"
	"user-service/internal/util"
)

type UserService struct {
	pkg.UnimplementedUserServiceServer
}

// RegisterUser 用户注册
func (u UserService) RegisterUser(ctx context.Context, req *pkg.RegisterUserRequest) (*pkg.UserResponse, error) {
	// 输入验证
	if req.Name == "" || req.Email == "" || req.Password == "" {
		return nil, errors.New("name, email, and password are required")
	}

	existingUser, err := util.GetUserByEmail(req.Email)
	if err != nil {
		return nil, err
	}
	if existingUser != nil {
		fmt.Println("existingUser")
		return nil, errors.New("email is already registered")
	}
	uuid, err := util.GenerateUUID()
	if err != nil {
		log.Println("uuid failed")
	}
	newUser := &model.User{
		Uuid:     uuid,
		Name:     req.Name,
		Email:    req.Email,
		Password: req.Password,
	}
	err = util.AddUser(newUser)
	if err != nil {
		return nil, err
	}
	return &pkg.UserResponse{
		User: &pkg.UserInfo{Name: newUser.Name, Email: newUser.Email},
	}, nil
}

// LoginUser 用户登录
func (u UserService) LoginUser(ctx context.Context, req *pkg.LoginUserRequest) (*pkg.UserResponse, error) {

	return &pkg.UserResponse{}, nil
}

// UpdateUserInfo 更新用户信息
func (u UserService) UpdateUserInfo(ctx context.Context, req *pkg.UpdateUserInfoRequest) (*pkg.UserResponse, error) {
	if req.Name == "" || req.Email == "" {
		return nil, errors.New("name, email are required")
	}

	exitsUser, err := util.UserIsExist(req.Uuid)
	// 要么不存在，要么查询出错
	if exitsUser == nil {
		return nil, err
	}

	// 存在
	user, err := util.UserUpdateInfo(req)
	if err != nil {
		return nil, err
	}
	return &pkg.UserResponse{
		User: &pkg.UserInfo{Name: user.Name, Email: user.Email},
	}, nil
}

// SetDefaultCity 修改默认地址
func (u UserService) SetDefaultCity(ctx context.Context, req *pkg.SetDefaultCityRequest) (*pkg.UserResponse, error) {
	if req.DefaultCity == "" {
		return nil, errors.New("city is required")
	}

	exitsUser, err := util.UserIsExist(req.Uuid)
	// 要么不存在，要么查询出错
	if exitsUser == nil {
		return nil, err
	}

	// 存在
	user, err := util.SetDefaultCity(req)
	if err != nil {
		return nil, err
	}
	return &pkg.UserResponse{
		User: &pkg.UserInfo{Uuid: req.Uuid, DefaultCity: user.DefaultCity},
	}, nil
	return &pkg.UserResponse{}, nil
}
