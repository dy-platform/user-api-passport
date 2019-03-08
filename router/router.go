package router

import (
	"github.com/gin-gonic/gin"
)



func Init() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)

	r := gin.Default()

	r.POST("/wechat_signin", WeChatSignIn)


	return  r
}