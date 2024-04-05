package db

import (
	"testing"
	"user-service/config"
	"user-service/internal/model"
)

func TestGetDB(t *testing.T) {
	config.Init()
	//log.InitLog()
	var user model.User
	if err := GetDB().Where("name = ?", "ft").First(&user).Error; err != nil {
		t.Errorf("failed to open database")
	}
	if user.Email != "ft@123.com" {
		t.Errorf("query error")
	}
}
