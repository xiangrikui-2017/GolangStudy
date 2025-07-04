package entity

import (
	"GolangStudy/blog/user/model/entity"
	"gorm.io/gorm"
)

type Post struct {
	gorm.Model
	Title   string `gorm:"type:varchar(128);not null"`
	Content string `gorm:"type:varchar(1024);not null"`
	UserID  uint   `gorm:"not null"`
	User    entity.User
}
