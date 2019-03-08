package main

import (
	"github.com/sirupsen/logrus"
	"github.com/dy-platform/user-api-passport/router"
	"github.com/dy-gopkg/kit/web"
)

func main() {
	// Create service
	web.Init()

	r := router.Init()
	// Register Handler
	web.DefaultService.Handle("/", r)

	logrus.Debugf("aaaaaaaaaaaaaaaaaaaaa")
	// Run server
	web.Run()
}