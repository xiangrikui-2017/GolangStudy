package routes

import "github.com/gin-gonic/gin"

func initCommentRoutes(publicGroup *gin.RouterGroup, privateGroup *gin.RouterGroup) {
	// 公共资源
	publicGroup.Group("/commnet")
	{

	}
	// 权限资源
	privateGroup.Group("/commnet")
	{

	}
}
