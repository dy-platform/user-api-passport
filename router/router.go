package router

import (
	"github.com/gin-gonic/gin"
)



func Init() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)

	r := gin.Default()

	r.POST("/user/passport/wechat-signin", WeChatSignIn)


	return  r
}