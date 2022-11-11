package main

import (
	"auth/handler"
	"auth/option"
	pb "auth/proto"
	log "go-micro.dev/v4/logger"
)

func main() {
	// Create service
	svc := option.InitService()

	//Register handler
	pb.RegisterAuthHandler(svc.Server(), handler.WireAuthHandler(svc))
	// Run service
	if err := svc.Run(); err != nil {
		log.Fatal(err)
	}
}
