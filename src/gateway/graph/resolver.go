package graph

import (
	"github.com/gin-gonic/gin"
	"go-micro.dev/v4"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	*gin.Context
	micro.Service
}
