package mongo

import (
	"article/micro"
	log "go-micro.dev/v4/logger"
	"go.mongodb.org/mongo-driver/mongo"
	"util"
)

var mongoStore = new(mongoCollection)

type mongoCollection struct {
	Client            *mongo.Client
	ArticleCollection *mongo.Collection
}

func InitMongo() error {
	var err error
	if mongoStore.Client, err = util.MongoConnection(micro.MongoUser, micro.MongoPassword, micro.MongoUrl); err != nil {
		log.Errorf("connect mongo error:%v", err)
		return err
	}
	db := mongoStore.Client.Database("article")

	if err = util.CreateCollection("articles", db); err != nil {
		log.Errorf("create articles collection error:%v", err)
		return err
	}

	mongoStore.ArticleCollection = db.Collection("articles")
	return nil
}
