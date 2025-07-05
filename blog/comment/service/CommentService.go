package service

import (
	"GolangStudy/blog/comment/model/dto"
	"GolangStudy/blog/comment/model/entity"
	"GolangStudy/blog/common/config"
	"GolangStudy/blog/common/result"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreatComment(ctx *gin.Context) {
	var createComment dto.CreateComment
	if err := ctx.ShouldBindJSON(&createComment); err != nil {
		ctx.JSON(http.StatusBadRequest, result.Error(err.Error()))
		return
	}
	userId := ctx.GetUint("UserId")
	comment := entity.Comment{
		Content: createComment.Content,
		UserID:  userId,
		PostID:  createComment.PostID,
	}
	if err := config.DB.Create(&comment).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, result.Error("评论失败"))
		return
	}
	ctx.JSON(http.StatusOK, result.Success(comment))
}

func CommentList(ctx *gin.Context) {
	var listComment dto.ListComment
	if err := ctx.ShouldBindJSON(&listComment); err != nil {
		ctx.JSON(http.StatusBadRequest, result.Error(err.Error()))
		return
	}
	if listComment.PageSize == 0 {
		listComment.PageSize = 10
	}
	if listComment.PageNum == 0 {
		listComment.PageNum = 1
	}
	var total int64
	var comments []entity.Comment

	baseQuery := config.DB.Model(&entity.Comment{})
	if listComment.PostID != 0 {
		baseQuery.
			Where("post_id = ?", listComment.PostID)
	}
	baseQuery.Count(&total)
	baseQuery.Limit(listComment.PageSize).Offset((listComment.PageNum - 1) * listComment.PageSize).Find(&comments)
	ctx.JSON(http.StatusOK, result.PageSuccess(comments, listComment.PageNum, listComment.PageSize, int(total)))
}

func DeleteComment(ctx *gin.Context) {
	commentId := ctx.Param("id")

	var comment entity.Comment
	config.DB.Select("id, user_id").Where("id = ?", commentId).Delete(&comment)
	if comment.UserID != ctx.GetUint("UserId") {
		ctx.JSON(http.StatusInternalServerError, result.Error("无权删除他人评论"))
	}

	err := config.DB.Delete(&entity.Comment{}, commentId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, result.Error("评论删除失败"))
	}
	ctx.JSON(http.StatusOK, result.Success("评论删除成功"))
}
