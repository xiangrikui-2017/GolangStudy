package routes

import (
	"GolangStudy/blog/common/middleware"
	"github.com/gin-gonic/gin"
)

func InitRoutes(r *gin.Engine) {
	// 统一前缀
	apiGroup := r.Group("/api")
	apiGroup.Use(middleware.GlobalException())
	// 公共服务
	publicGroup := apiGroup.Group("")
	// 权限服务
	privateGroup := apiGroup.Group("", middleware.JwtAuth())

	// 用户模块
	initUserRoutes(publicGroup, privateGroup)
	// 帖子模块
	initPostRoutes(publicGroup, privateGroup)
	// 评论模块
	initCommentRoutes(publicGroup, privateGroup)
}
