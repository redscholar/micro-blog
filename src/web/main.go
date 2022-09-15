package main

import (
	log "go-micro.dev/v4/logger"
	"web/handler"
	"web/micro"
	pb "web/proto"
)

func main() {
	go micro.HttpConfig()
	go micro.GinRouter()
	// Create service
	srv := micro.Service
	//srv.Init()
	//Register handler
	pb.RegisterWebHandler(srv.Server(), new(handler.Web))
	// Run service
	if err := srv.Run(); err != nil {
		log.Fatal(err)
	}
}
