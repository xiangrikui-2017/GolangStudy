package main

import (
	"GolangStudy/gin/myfunc"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/", myfunc.HelloWorld)
	router.Run(":8080")
}
