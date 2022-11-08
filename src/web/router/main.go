package router

import (
	"context"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"go-micro.dev/v4/debug/trace"
	"go-micro.dev/v4/metadata"
	"net/http"
	"strconv"
	"time"
	"util"
	"web/micro"
)

func GinRouter() {
	router := gin.Default()
	// 跨域
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowHeaders:     []string{"*"},
		AllowCredentials: true,
	}))
	// 链路追踪
	router.Use(func(c *gin.Context) {
		newCtx, s := trace.DefaultTracer.Start(context.Background(), "web")
		s.Type = trace.SpanTypeRequestInbound
		defer trace.DefaultTracer.Finish(s)
		md, _ := metadata.FromContext(newCtx)
		c.Set(micro.ClientCtx, newCtx)
		c.Header("Trace-Id", md["Micro-Trace-Id"])
	})
	// token校验
	router.Use(func(c *gin.Context) {
		for _, s := range micro.Service.Options().Config.Get("auth", "whiteList").StringSlice(make([]string, 0)) {
			if c.Request.URL.Path == s {
				c.Next()
				return
			}
		}
		token := c.GetHeader(micro.AuthHeader)
		if token == "" {
			token = "eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJ0eXBlIjoidXNlciIsInNjb3BlcyI6bnVsbCwibWV0YWRhdGEiOnsiY3JlYXRlQXQiOiIxNjY3ODE1MTYzIiwiZXhwaXJlQXQiOiIxNjY3ODE2MzYzIn0sImV4cCI6MTY2NzgxNjM2MywiaXNzIjoiYmxvZyIsInN1YiI6IjVhZmQ4YjQ4LTRjM2QtNGRmMC04Nzk2LTJhMjk5ODdmNzYyZSJ9.GdObmavloX6gVEQi5zMSh972LAM3vUk1ci10fMni65gwHvCzGclo7_ILSDrMU5yvPzEOSDrUg4agBQ5tIZ8RXlVX1iOBNaSQee0ldQU3m3jL3bAXnGNmnL8lduQhRGqnm0VQcTh6RVmH6LNB5YbSDJRNYa2qy06GQ4LASIeaGdrggLhEm9O1oFX2HJCZ31qmuangs_ko-kkWEIusCKdcEW_0-ekXTWzLQHMZubWqYwuwMubnviK8lh6V4qJmNhQ3GIj1mud2-32d-ILdlfWZFwAtQ7BVF0us6yUUl_W88ACVRbymUgARIirirBoeT5gmftiLc4wnPqSiK2jSxglN7A"
		}
		account, err := micro.Service.Options().Auth.Inspect(token)
		if err != nil {
			c.JSON(http.StatusForbidden, gin.H{
				"code": -1,
				"msg":  err.Error(),
			})
			c.Abort()
			return
		}
		expireAt, _ := strconv.ParseInt(account.Metadata["expireAt"], 10, 0)
		if expireAt-time.Now().Unix() < int64(micro.Service.Options().Config.Get("auth", "refreshTime").Int(0)) {
			// 生成新的token
			if newToken, err := util.GenToken(micro.Service, ""); err != nil {
				c.JSON(http.StatusForbidden, gin.H{
					"code": -1,
					"msg":  err.Error(),
				})
				c.Abort()
				return
			} else {
				c.Header(micro.AuthHeader, newToken)
			}

		}
		c.Set(micro.AuthHeader, token)
		if cctx, exist := c.Get(micro.ClientCtx); exist {
			c.Set(micro.ClientCtx, metadata.NewContext(cctx.(context.Context), map[string]string{
				micro.AuthHeader: token,
			}))
		} else {
			c.Set(micro.ClientCtx, metadata.NewContext(context.Background(), map[string]string{
				micro.AuthHeader: token,
			}))
		}
	})
	router = Route(router)

	err := router.Run(":5001")
	if err != nil {
		return
	}
}
