package router

import (
	"context"
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
	router.Use(func(c *gin.Context) {
		method := c.Request.Method
		origin := c.Request.Header.Get("Origin")
		if origin != "" {
			c.Header("Access-Control-Allow-Origin", "*") // 可将将 * 替换为指定的域名
			c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
			c.Header("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept, Authorization")
			c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Cache-Control, Content-Language, Content-Type")
			c.Header("Access-Control-Allow-Credentials", "true")
		}
		if method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		}
		c.Next()
	})
	// 链路追踪
	router.Use(func(c *gin.Context) {
		newCtx, s := trace.DefaultTracer.Start(context.Background(), "web")
		s.Type = trace.SpanTypeRequestInbound
		defer trace.DefaultTracer.Finish(s)
		md, _ := metadata.FromContext(newCtx)
		c.Set("clientCtx", newCtx)
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
		account, err := micro.Service.Options().Auth.Inspect(token)
		if err != nil {
			c.JSON(http.StatusForbidden, gin.H{
				"code": -1,
				"msg":  err.Error(),
			})
			c.Abort()
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
			} else {
				c.Header(micro.AuthHeader, newToken)
			}

		}
		c.Set(micro.AuthHeader, token)
	})
	router = Route(router)
	err := router.Run(":5001")
	if err != nil {
		return
	}
}
