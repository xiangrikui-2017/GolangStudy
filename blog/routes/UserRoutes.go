package routes

import (
	UserService "GolangStudy/blog/user/service"
	"github.com/gin-gonic/gin"
)

func initUserRoutes(publicGroup *gin.RouterGroup, privateGroup *gin.RouterGroup) {
	// 公共资源
	publicGroup.Group("/user")
	{
		publicGroup.POST("/register", UserService.Register)
		publicGroup.POST("/login", UserService.Login)
	}
	// 权限资源
	privateGroup.Group("/user")
	{

	}
}
