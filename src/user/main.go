package main

import (
	log "go-micro.dev/v4/logger"
	"user/handler"
	"user/micro"
	pb "user/proto"
)

func main() {
	// Create service
	micro.InitService()

	//Register handler
	pb.RegisterUserHandler(micro.Service.Server(), new(handler.User))
	// Run service
	if err := micro.Service.Run(); err != nil {
		log.Fatal(err)
	}
}
