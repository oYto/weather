package service

import (
	"WeatherQuery/user-service/internal/model"
	"WeatherQuery/user-service/internal/util"
	"WeatherQuery/user-service/proto"
	"context"
	"errors"
	"fmt"
	"log"
)

type UserService struct {
	proto.UnimplementedUserServiceServer
}

// RegisterUser 用户注册
func (u UserService) RegisterUser(ctx context.Context, req *proto.RegisterUserRequest) (*proto.UserResponse, error) {
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
	//uuid, err := util.GenerateUUID()
	if err != nil {
		log.Println("uuid failed")
	}
	newUser := &model.User{
		Name:     req.Name,
		Email:    req.Email,
		Password: req.Password,
	}
	err = util.AddUser(newUser)
	if err != nil {
		return nil, err
	}
	return &proto.UserResponse{
		User: &proto.UserInfo{Name: newUser.Name, Email: newUser.Email},
	}, nil
}

// LoginUser 用户登录
func (u UserService) LoginUser(ctx context.Context, req *proto.LoginUserRequest) (*proto.UserResponse, error) {

	return &proto.UserResponse{}, nil
}

// UpdateUserInfo 更新用户信息
func (u UserService) UpdateUserInfo(ctx context.Context, req *proto.UpdateUserInfoRequest) (*proto.UserResponse, error) {

	return &proto.UserResponse{}, nil
}

func (u UserService) SetDefaultCity(ctx context.Context, req *proto.SetDefaultCityRequest) (*proto.UserResponse, error) {

	return &proto.UserResponse{}, nil
}
