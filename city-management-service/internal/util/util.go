package util

import (
	"city/internal/db"
	"city/internal/model"
	"city/proto"
	"errors"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"log"
)

func IsNotExist(req *proto.AddCityRequest) (bool, error) {
	if err := db.GetDB().Where("name = ? and country = ?", req.Name, req.Country).First(&model.City{}).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return true, nil
		}
		return false, err
	}
	return false, nil
}

func AddCity(req *proto.AddCityRequest) (*model.City, error) {
	uuid, err := GenerateUUID()
	if err != nil {
		return nil, errors.New("generateUUID failed")
	}

	city := &model.City{
		UUID:      uuid,
		Name:      req.Name,
		Country:   req.Country,
		Longitude: req.Longitude,
		Latitude:  req.Latitude,
	}
	if err := db.GetDB().Create(&city).Error; err != nil {
		log.Printf("failed to insert city: %v", err)
	}
	return city, nil
}

func SearchCityByName(name string) ([]model.City, error) {
	var cities []model.City
	if err := db.GetDB().Where("name = ?", name).Find(&cities).Error; err != nil {
		return nil, err
	}

	return cities, nil
}

func GetListCity() ([]model.City, error) {
	var cities []model.City
	if err := db.GetDB().Find(&cities).Error; err != nil {
		return nil, err
	}
	return cities, nil
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
