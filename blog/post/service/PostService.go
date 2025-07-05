package service

import (
	"GolangStudy/blog/common/config"
	"GolangStudy/blog/common/result"
	PostDto "GolangStudy/blog/post/model/dto"
	PostEntity "GolangStudy/blog/post/model/entity"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreatePost(ctx *gin.Context) {
	var createPost PostDto.PostCreateDto
	if err := ctx.ShouldBindJSON(&createPost); err != nil {
		ctx.JSON(http.StatusBadRequest, result.Error(err.Error()))
		return
	}
	userId := ctx.GetUint("UserID")
	post := PostEntity.Post{
		Title:   createPost.Title,
		Content: createPost.Content,
		UserID:  userId,
	}
	if err := config.DB.Create(&post).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, result.Error("文章创建失败"))
	}
	ctx.JSON(http.StatusOK, result.Success(post))
}

func UpdatePost(ctx *gin.Context) {
	var updatePost PostDto.PostUpdateDto
	if err := ctx.ShouldBindJSON(&updatePost); err != nil {
		ctx.JSON(http.StatusBadRequest, result.Error(err.Error()))
		return
	}
	if updatePost.Title == "" || updatePost.Content == "" {
		ctx.JSON(http.StatusBadRequest, result.Error("文章标题或内容不能为空"))
		return
	}
	var post PostEntity.Post
	err := config.DB.Where("id = ?", updatePost.ID).First(&post)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, result.Error("更新文章查询无数据"))
		return
	}
	// 校验当前请求人是否是文章所有者
	if post.UserID != ctx.GetUint("UserID") {
		ctx.JSON(http.StatusInternalServerError, result.Error("无权更新此文章"))
		return
	}
	post.Title = updatePost.Title
	post.Content = updatePost.Content
	err = config.DB.Where("id = ?", post.ID).Updates(post)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, "文章更新失败")
		return
	}
	ctx.JSON(http.StatusOK, result.Success(post))
}

func GetPost(ctx *gin.Context) {
	var getPost PostDto.PostGetOrDelDto
	if err := ctx.ShouldBindJSON(&getPost); err != nil {
		ctx.JSON(http.StatusBadRequest, result.Error(err.Error()))
		return
	}
}

func DeletePost(ctx *gin.Context) {

}
