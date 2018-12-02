package main

import (
	"github.com/dy-platform/user-api-passport/handler"
	"github.com/dy-platform/user-api-passport/idl/platform/user/api-passport"
	"github.com/dy-gopkg/kit"
	"github.com/dy-platform/user-srv-passport/util/config"
)

func main() {
	kit.Init()

	// 初始化业务配置
	uconfig.InitBusinessConfig()

	//TODO 初始化缓存
	//cache.CacheInit()

	platform_user_api_passport.RegisterPassportHandler(kit.Server(), &handler.Handler{})

	kit.Run()
}