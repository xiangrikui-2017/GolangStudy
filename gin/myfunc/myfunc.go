package myfunc

import "github.com/gin-gonic/gin"

func HelloWorld(ctx *gin.Context) {
	ctx.String(200, "hello world")
}
