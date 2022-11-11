package main

import (
	"web/router"
)

func main() {
	svc := option.InitService()
	router.GinHttp(svc)
}
