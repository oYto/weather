package util

import (
	"errors"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"user-service/internal/db"
	"user-service/internal/model"
	pkg "user-service/pkg/pb"
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

func UserIsExist(uuid string) (*model.User, error) {
	user := &model.User{}
	if err := db.GetDB().Where("uuid = ?", uuid).First(user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.New("user is not exist")
		}
		return nil, err
	}
	return user, errors.New("user is exist")
}

func UserUpdateInfo(req *pkg.UpdateUserInfoRequest) (*model.User, error) {
	user := &model.User{}
	if err := db.GetDB().Model(&user).Where("uuid = ?", req.Uuid).
		Updates(map[string]interface{}{"email": req.Email, "name": req.Name}).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func SetDefaultCity(req *pkg.SetDefaultCityRequest) (*model.User, error) {
	user := &model.User{}
	if err := db.GetDB().Model(&user).Where("uuid = ?", req.Uuid).
		Update("default_city", req.DefaultCity).Error; err != nil {
		return nil, err
	}
	return user, nil
}
func AddUser(user *model.User) error {
	if err := db.GetDB().Create(user).Error; err != nil {
		return err
	}
	return nil
}

// GenerateUUID 生成一个UUID并返回其字符串表示形式
func GenerateUUID() (string, error) {
	// 生成UUID
	id, err := uuid.NewRandom()
	if err != nil {
		return "", err
	}
	// 将UUID转换为字符串
	uuidStr := id.String()
	return uuidStr, nil
}
