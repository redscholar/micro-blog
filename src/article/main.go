package main

import (
	"article/handler"
	pb "article/proto"
	"go-micro.dev/v4"

	log "go-micro.dev/v4/logger"
)

var (
	service = "system"
	version = "latest"
)

func main() {
	// Create service
	srv := micro.NewService(
		micro.Name(service),
		micro.Version(version),
	)
	srv.Init()

	// Register handler
	pb.RegisterSystemHandler(srv.Server(), new(handler.System))
	// Run service
	if err := srv.Run(); err != nil {
		log.Fatal(err)
	}
}