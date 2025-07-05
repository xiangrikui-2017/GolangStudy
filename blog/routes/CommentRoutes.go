package routes

import (
	CommentService "GolangStudy/blog/comment/service"
	"github.com/gin-gonic/gin"
)

func initCommentRoutes(publicGroup *gin.RouterGroup, privateGroup *gin.RouterGroup) {
	// 公共资源
	publicGroup.Group("/commnet")
	{

	}
	// 权限资源
	commentPrivGroup := privateGroup.Group("/commnet")
	{
		commentPrivGroup.POST("/create", CommentService.CreatComment)
		commentPrivGroup.GET("/delete/:id", CommentService.DeleteComment)
		commentPrivGroup.POST("/list", CommentService.CommentList)
	}
}
