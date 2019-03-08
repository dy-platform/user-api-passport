package main

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/dy-platform/user-api-passport/router"
	"github.com/dy-gopkg/kit/web"
)

func main() {
	// Create service
	fmt.Println("aaaaaaaaaaaaaaaaaaaaaaaaa")
	web.Init()
	fmt.Println("bbbbbbbbbbbbbbbbbbb")

	r := router.Init()

	fmt.Println("ccccccccccccccccccccc")

	// Register Handler
	web.DefaultService.Handle("/", r)

	fmt.Println("ddddddddddddddddddd")


	logrus.Debugf("aaaaaaaaaaaaaaaaaaaaa")
	// Run server
	web.Run()
}