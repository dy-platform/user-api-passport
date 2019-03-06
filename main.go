package main

import (
	"github.com/sirupsen/logrus"
	"github.com/dy-platform/user-api-passport/handler"
	"github.com/dy-platform/user-api-passport/idl/platform/user/api-passport"
	"github.com/dy-gopkg/kit"
)

func main() {
	kit.Init()

	err := platform_user_api_passport.RegisterPassportHandler(kit.Server(), &handler.Handler{})
	if err != nil {
		logrus.Fatalf("RegisterPassportHandler error:%v", err)
	}

	kit.Run()
}