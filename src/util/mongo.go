package util

import (
	"context"
	"go-micro.dev/v4/logger"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"strings"
	"time"
)

// MongoConnection creates a connection to the mongo
func MongoConnection(dbuser, dbpassword, dburl string) (*mongo.Client, error) {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	mongoCredentials := options.Credential{
		Username: dbuser,
		Password: dbpassword,
	}
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(dburl).SetAuth(mongoCredentials))
	if err != nil {
		return nil, err
	}

	return client, nil
}

// CreateIndex creates a unique index for the given field in the collectionName
func CreateIndex(collectionName string, field string, db *mongo.Database) error {
	mod := mongo.IndexModel{
		Keys:    bson.M{field: 1},
		Options: options.Index().SetUnique(true),
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	collection := db.Collection(collectionName)
	_, err := collection.Indexes().CreateOne(ctx, mod)
	if err != nil {
		logger.Error(err)
		return err
	}
	return nil
}

// CreateCollection creates a new mongo collection if it does not exist
func CreateCollection(collectionName string, db *mongo.Database) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err := db.CreateCollection(ctx, collectionName)
	if err != nil {
		if strings.Contains(err.Error(), "already exists") {
			logger.Info(collectionName + "'s collection already exists, continuing with the existing mongo collection")
			return nil
		} else {
			return err
		}
	}

	logger.Info(collectionName + "'s mongo collection created")
	return nil
}
