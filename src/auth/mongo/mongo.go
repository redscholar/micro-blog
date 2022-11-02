package mongo

import (
	"auth/micro"
	log "go-micro.dev/v4/logger"
	"go.mongodb.org/mongo-driver/mongo"
	"util"
)

var mongoStore = new(mongoCollection)

type mongoCollection struct {
	Client         *mongo.Client
	UserCollection *mongo.Collection
}

func InitMongo() error {
	var err error
	if mongoStore.Client, err = util.MongoConnection(micro.MongoUser, micro.MongoPassword, micro.MongoUrl); err != nil {
		log.Errorf("connect mongo error:%v", err)
		return err
	}
	db := mongoStore.Client.Database("auth")

	if err = util.CreateCollection("users", db); err != nil {
		log.Errorf("create user collection error:%v", err)

		return err
	}
	if err = util.CreateIndex("users", "username", db); err != nil {
		log.Errorf("create user index error:%v", err)
		return err
	}
	mongoStore.UserCollection = db.Collection("users")
	return nil
}
