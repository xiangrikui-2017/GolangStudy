package dto

import (
	CommonDto "GolangStudy/blog/common/dto"
)

type CreateComment struct {
	Content string `json:"content" binding:"required"`
	PostID  uint   `json:"post_id" binding:"required"`
}

type ListComment struct {
	PostID uint `json:"post_id" binding:"required"`
	CommonDto.PageDTO
}
