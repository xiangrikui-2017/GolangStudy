package routes

import "github.com/gin-gonic/gin"

func initPostRoutes(publicGroup *gin.RouterGroup, privateGroup *gin.RouterGroup) {
	// 公共资源
	publicGroup.Group("/post")
	{

	}
	// 权限资源
	privateGroup.Group("/post")
	{

	}
}
