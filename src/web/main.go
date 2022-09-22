package main

import (
	"web/config"
	"web/micro"
	"web/router"
)

func main() {
	go config.HttpConfig()
	micro.InitService()
	router.GinRouter()
}
