//go:build wireinject
// +build wireinject

package handler

import (
	"auth/store/mongo"
	"github.com/google/wire"
	"go-micro.dev/v4"
)

func WireAuthHandler(service micro.Service) *authHandler {
	wire.Build(NewAuthHandler, mongo.WireMongoUserStore)
	return &authHandler{}
}
