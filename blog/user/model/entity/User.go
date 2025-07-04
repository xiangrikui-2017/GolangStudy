package entity

import "gorm.io/gorm"

type User struct {
	gorm.Model
	UserName string `gorm:"type:varchar(20);unique;not null"`
	Password string `gorm:"type:varchar(128);not null"`
	Email    string `gorm:"type:varchar(128);unique;not null"`
}
