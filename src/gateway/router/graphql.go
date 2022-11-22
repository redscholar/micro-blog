package router

import (
	"gateway/graph"
	"gateway/graph/generated"
	graphHandler "github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gin-gonic/gin"
	"go-micro.dev/v4"
)

func graphqlRoute(r *gin.Engine, svc micro.Service) *gin.Engine {
	r.POST("/query", func(c *gin.Context) {
		graphHandler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{Context: c, Service: svc}})).ServeHTTP(c.Writer, c.Request)
	})
	r.GET("/", func(c *gin.Context) {
		playground.Handler("GraphQL", "/query").ServeHTTP(c.Writer, c.Request)
	})
	return r
}
