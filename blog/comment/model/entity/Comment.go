package entity

import (
	PostEntity "GolangStudy/blog/post/model/entity"
	UserEntity "GolangStudy/blog/user/model/entity"
	"gorm.io/gorm"
)

type Comment struct {
	gorm.Model
	Content string `gorm:"type:varchar(128);not null"`
	UserID  uint
	User    UserEntity.User
	PostID  uint
	Post    PostEntity.Post
}
