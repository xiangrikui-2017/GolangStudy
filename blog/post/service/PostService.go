package service

import (
	"GolangStudy/blog/common/config"
	"GolangStudy/blog/common/result"
	PostDto "GolangStudy/blog/post/model/dto"
	PostEntity "GolangStudy/blog/post/model/entity"
	"errors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
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

func ListPost(ctx *gin.Context) {
	var listPost PostDto.PostListDto
	if err := ctx.ShouldBindQuery(&listPost); err != nil {
		ctx.JSON(http.StatusBadRequest, result.Error(err.Error()))
		return
	}
	// 设置分页信息
	if listPost.Page.PageNum == 0 {
		listPost.Page.PageNum = 1
	}
	if listPost.Page.PageSize == 0 {
		listPost.Page.PageSize = 10
	}
	var total int64
	var posts []PostEntity.Post
	baseQuery := config.DB.Omit("content").
		Count(&total).
		Limit(listPost.Page.PageSize).
		Offset((listPost.Page.PageNum - 1) * listPost.Page.PageSize).
		Order("create_at desc")
	if listPost.ID != 0 {
		baseQuery = baseQuery.Where("id = ?", listPost.ID)
	}
	if listPost.UserID != 0 {
		baseQuery = baseQuery.Where("user_id = ?", listPost.UserID)
	}
	if listPost.Title != "" {
		baseQuery = baseQuery.Where("title like ?", "%"+listPost.Title+"%")
	}
	if err := baseQuery.Find(&posts).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, result.Error(err.Error()))
	}
	ctx.JSON(http.StatusOK, result.Success(posts))
}

func GetPostById(ctx *gin.Context) {
	id := ctx.Param("id")

	var post PostEntity.Post
	err := config.DB.First(&post, id).Error
	if err != nil {
		ctx.JSON(http.StatusBadRequest, result.Error("获取文章详情失败"))
		return
	}
	ctx.JSON(http.StatusOK, result.Success(post))
}

func DeletePost(ctx *gin.Context) {
	var getPost PostDto.PostGetOrDelDto
	if err := ctx.ShouldBindJSON(&getPost); err != nil {
		ctx.JSON(http.StatusBadRequest, result.Error(err.Error()))
		return
	}

	var post PostEntity.Post
	err := config.DB.Select("id,user_id").Take(&post, getPost.ID).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			ctx.JSON(http.StatusBadRequest, result.Error("未查询到要删除的文章"))
			return
		}
	}
	if post.UserID != ctx.GetUint("UserId") {
		ctx.JSON(http.StatusBadRequest, result.Error("文章无删除权限"))
		return
	}

	if err := config.DB.Delete(&PostEntity.Post{}, getPost.ID).Error; err != nil {
		ctx.JSON(http.StatusOK, result.Error("删除失败"))
	}

	ctx.JSON(http.StatusOK, result.Success("删除成功"))
}
