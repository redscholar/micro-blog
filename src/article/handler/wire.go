//go:build wireinject
// +build wireinject

package handler

import (
	"article/store/mongo"
	"github.com/google/wire"
	"go-micro.dev/v4"
)

func WireArticleHandler(service micro.Service) *articleHandler {
	wire.Build(NewArticle, mongo.WireMongoArticleStore)
	return &articleHandler{}
}
