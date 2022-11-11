package main

import (
	"article/handler"
	"article/option"
	pb "article/proto/article"
	log "go-micro.dev/v4/logger"
)

func main() {
	// Create service
	svc := option.InitService()

	//// Register handler
	pb.RegisterArticleHandler(svc.Server(), handler.WireArticleHandler(svc))
	// Run service
	if err := svc.Run(); err != nil {
		log.Fatal(err)
	}
}
