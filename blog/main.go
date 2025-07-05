package main

import (
	CommentEntity "GolangStudy/blog/comment/model/entity"
	"GolangStudy/blog/common/config"
	PostEntity "GolangStudy/blog/post/model/entity"
	"GolangStudy/blog/routes"
	UserEntity "GolangStudy/blog/user/model/entity"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	// 初始化配置
	config.InitConfig("blog/common/config/config.yaml")
	// 初始化路由
	routes.InitRoutes(router)
	// 初始化数据库
	db := config.InitMysqlDB()
	db.AutoMigrate(&UserEntity.User{}, &PostEntity.Post{}, &CommentEntity.Comment{})
	// 初始化日志
	config.InitLogrus()

	router.Run(":9090")
}
