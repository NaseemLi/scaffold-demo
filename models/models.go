// 模型层
package models

import (
	"scaffold-demo/config"
	"scaffold-demo/utils"
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID        uint           `gorm:"primaryKey" json:"-"`
	Username  string         `gorm:"size:64;not null;uniqueIndex" json:"username"`
	Password  string         `gorm:"size:128;not null" json:"-"`
	CreatedAt time.Time      `json:"-"`
	UpdatedAt time.Time      `json:"-"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}

func AutoMigrate() error {
	return config.DB.AutoMigrate(&User{})
}

func InitAdminUser() error {
	var count int64
	if err := config.DB.Model(&User{}).Count(&count).Error; err != nil {
		return err
	}
	if count > 0 {
		return nil
	}

	hashedPassword, err := utils.HashPassword(config.Password)
	if err != nil {
		return err
	}

	admin := User{
		Username: config.Username,
		Password: hashedPassword,
	}
	return config.DB.Create(&admin).Error
}
