package router

import (
	"github.com/gin-gonic/gin"
)



func Init() *gin.Engine {
	r := gin.Default()

	r.POST("/wechat_signin", WeChatSignIn)


	return  r
}