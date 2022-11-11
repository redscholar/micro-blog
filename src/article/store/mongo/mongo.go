package mongo

import (
	log "go-micro.dev/v4/logger"
	"go.mongodb.org/mongo-driver/mongo"
	"sync"
	"util"
)

var mongoStore = new(mongoCollection)
var mongoOnce = sync.Once{}

type mongoCollection struct {
	client            *mongo.Client
	articleCollection *mongo.Collection
}

func initMongoStore() *mongoCollection {
	mongoOnce.Do(func() {
		var err error
		if mongoStore.client, err = util.MongoConnection(util.ComParam.MongoUser, util.ComParam.MongoPassword, util.ComParam.MongoUrl); err != nil {
			log.Errorf("connect mongo error:%v", err)
		}
		db := mongoStore.client.Database("article")

		if err = util.CreateCollection("articles", db); err != nil {
			log.Errorf("create articles collection error:%v", err)
		}
		mongoStore.articleCollection = db.Collection("articles")
	})

	return mongoStore
}
