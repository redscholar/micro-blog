package main

import (
	"gateway/option"
	"gateway/router"
)

func main() {
	svc := option.InitService()
	router.GinHttp(svc)
}
