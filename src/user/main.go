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
	pb.RegisterAuthHandler(micro.Service.Server(), new(handler.Auth))
	// Run service
	if err := micro.Service.Run(); err != nil {
		log.Fatal(err)
	}
}
