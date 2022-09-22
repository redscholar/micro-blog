package router

import (
	"github.com/gin-gonic/gin"
	"go-micro.dev/v4/auth"
	"html/template"
	"net/http"
	"strconv"
	"time"
	"web/micro"
)

func GinRouter() {
	router := gin.Default()
	router.SetFuncMap(template.FuncMap{
		//模板中的两个变量相加，实现分页
		"add": func(x, y int) int {
			return x + y
		},
		//将返回到前端的字符串转换为html代码，如数据库字段为 <p>xx</p>,转换为前端的p标签
		"tran": func(code string) template.HTML {
			return template.HTML(code)
		},
	})

	router.LoadHTMLGlob("template/**/*")
	router.StaticFS("static/", http.Dir("./static"))
	router.Use(func(c *gin.Context) {
		for _, s := range micro.Service.Options().Config.Get("auth", "whiteList").StringSlice(make([]string, 0)) {
			if c.Request.URL.Path == s {
				c.Next()
				return
			}
		}
		token := c.GetHeader(micro.AuthHeader)
		if token == "" {
			token, _ = c.GetQuery("token")
		}
		account, err := micro.Service.Options().Auth.Inspect(token)
		if err != nil {
			c.Redirect(http.StatusMovedPermanently, "/login")
			c.Abort()
		}
		expireAt, _ := strconv.ParseInt(account.Metadata["expireAt"], 10, 0)
		if expireAt-time.Now().Unix() < int64(micro.Service.Options().Config.Get("auth", "refreshTime").Int(0)) {
			// 生成新的token
			expireTime := micro.Service.Options().Config.Get("auth", "expireTime").Int(0)
			newAccount, err := micro.Service.Options().Auth.Generate(account.ID, auth.WithType("user"),
				auth.WithMetadata(map[string]string{
					"createAt": strconv.FormatInt(time.Now().Unix(), 10),
					"expireAt": strconv.FormatInt(time.Now().Add(time.Second*time.Duration(expireTime)).Unix(), 10),
				}))
			if err != nil {
				c.Redirect(http.StatusMovedPermanently, "/login")
				c.Abort()
			}
			newToken, err := micro.Service.Options().Auth.Token(auth.WithExpiry(time.Second*time.Duration(expireTime)), auth.WithCredentials(newAccount.ID, newAccount.Secret))
			if err != nil {
				c.Redirect(http.StatusMovedPermanently, "/login")
				c.Abort()
			}
			token = newToken.AccessToken
		}
		c.Set(micro.AuthHeader, token)
		if err != nil {
			c.Redirect(http.StatusMovedPermanently, "/login")
			c.Abort()
		}
	})
	router = Route(router)
	err := router.Run(":5001")
	if err != nil {
		return
	}
}
