package main

import (
	"auth/handler"
	"auth/micro"
	"auth/mongo"
	pb "auth/proto"
	log "go-micro.dev/v4/logger"
)

func main() {
	// Create service
	micro.InitService()
	// Connect db
	mongo.InitMongo()
	//Register handler
	pb.RegisterAuthHandler(micro.Service.Server(), &handler.Auth{UserStore: mongo.NewUserStore()})
	// Run service
	if err := micro.Service.Run(); err != nil {
		log.Fatal(err)
	}
}
