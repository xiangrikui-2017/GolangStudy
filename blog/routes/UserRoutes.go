package routes

import (
	UserService "GolangStudy/blog/user/service"
	"github.com/gin-gonic/gin"
)

func initUserRoutes(publicGroup *gin.RouterGroup, privateGroup *gin.RouterGroup) {
	// 公共资源
	userPubGroup := publicGroup.Group("/user")
	{
		userPubGroup.POST("/register", UserService.Register)
		userPubGroup.POST("/login", UserService.Login)
		userPubGroup.GET("/hello", UserService.Hello)
	}
	// 权限资源
	privateGroup.Group("/user")
	{

	}
}
