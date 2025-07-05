package routes

import (
	PostService "GolangStudy/blog/post/service"
	"github.com/gin-gonic/gin"
)

func initPostRoutes(publicGroup *gin.RouterGroup, privateGroup *gin.RouterGroup) {
	// 公共资源
	postPubGroup := publicGroup.Group("/post")
	{
		postPubGroup.POST("/list", PostService.ListPost)
	}
	// 权限资源
	postPrivtGroup := privateGroup.Group("/post")
	{
		postPrivtGroup.GET("/get/:id", PostService.GetPostById)
		postPrivtGroup.POST("/create", PostService.CreatePost)
		postPrivtGroup.POST("/update", PostService.UpdatePost)
		postPrivtGroup.POST("/delete", PostService.DeletePost)
	}
}
