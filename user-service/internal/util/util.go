package util

import (
	"WeatherQuery/user-service/internal/db"
	"WeatherQuery/user-service/internal/model"
	"errors"
	"gorm.io/gorm"
)

func GetUserByEmail(email string) (*model.User, error) {
	var user model.User
	if err := db.GetDB().Where("email = ?", email).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			// 找不到用户，不算错误
			return nil, nil
		}
		// 查询出错
		return nil, err
	}
	return &user, nil
}

func AddUser(user *model.User) error {
	if err := db.GetDB().Create(user).Error; err != nil {
		return err
	}
	return nil
}

// GenerateUUID 生成一个UUID并返回其字符串表示形式
//func GenerateUUID() (string, error) {
//	// 生成UUID
//	id, err := uuid.NewRandom()
//	if err != nil {
//		return "", err
//	}
//	// 将UUID转换为字符串
//	uuidStr := id.String()
//	return uuidStr, nil
//}
