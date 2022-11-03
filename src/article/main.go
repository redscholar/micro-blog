package main

import (
	"article/handler"
	"article/micro"
	"article/mongo"
	pb "article/proto"
	log "go-micro.dev/v4/logger"
)

func main() {
	// Create service
	micro.InitService()
	// Connect db
	mongo.InitMongo()

	//// Register handler
	pb.RegisterArticleHandler(micro.Service.Server(), &handler.Article{ArticleStore: mongo.NewArticleStore()})
	// Run service
	if err := micro.Service.Run(); err != nil {
		log.Fatal(err)
	}
}
