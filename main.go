package main

import (
	"github.com/dy-gopkg/kit/web"
	"github.com/dy-platform/user-api-passport/router"
)

func main() {
	// Create service
	web.Init()

	r := router.Init()

	// Register Handler
	web.DefaultService.Handle("/", r)

	// Run server
	web.Run()
}