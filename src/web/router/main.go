package router

import (
	"context"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"go-micro.dev/v4"
	"go-micro.dev/v4/debug/trace"
	"go-micro.dev/v4/metadata"
	"net/http"
	"strconv"
	"time"
	"util"
	"web/option"
)

func GinHttp(svc micro.Service) {
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
		c.Set(option.ClientCtx, newCtx)
		c.Header("Trace-Id", md["Micro-Trace-Id"])
	})
	// token校验
	router.Use(func(c *gin.Context) {
		for _, s := range svc.Options().Config.Get("auth", "whiteList").StringSlice(make([]string, 0)) {
			if c.Request.URL.Path == s {
				c.Next()
				return
			}
		}
		token := c.GetHeader(option.AuthHeader)
		account, err := svc.Options().Auth.Inspect(token)
		if err != nil {
			c.JSON(http.StatusForbidden, gin.H{
				"code": -1,
				"msg":  err.Error(),
			})
			c.Abort()
			return
		}
		expireAt, _ := strconv.ParseInt(account.Metadata["expireAt"], 10, 0)
		if expireAt-time.Now().Unix() < int64(svc.Options().Config.Get("auth", "refreshTime").Int(0)) {
			// 生成新的token
			if newToken, err := util.GenToken(svc, ""); err != nil {
				c.JSON(http.StatusForbidden, gin.H{
					"code": -1,
					"msg":  err.Error(),
				})
				c.Abort()
				return
			} else {
				c.Header(option.AuthHeader, newToken)
			}

		}
		c.Set(option.AuthHeader, token)
		if cctx, exist := c.Get(option.ClientCtx); exist {
			c.Set(option.ClientCtx, metadata.NewContext(cctx.(context.Context), map[string]string{
				option.AuthHeader: token,
			}))
		} else {
			c.Set(option.ClientCtx, metadata.NewContext(context.Background(), map[string]string{
				option.AuthHeader: token,
			}))
		}
	})
	router = signRoute(router, svc)
	router = graphqlRoute(router)

	err := router.Run(":5001")
	if err != nil {
		return
	}
}
