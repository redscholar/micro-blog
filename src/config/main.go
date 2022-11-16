package main

import (
	_ "embed"
	"encoding/base64"
	"github.com/gin-gonic/gin"
	"log"
)

//go:embed public_key.pem
var publicKey []byte

//go:embed private_key.pem
var privateKey []byte

func main() {
	router := gin.Default()
	config := defaultConfig()
	router.POST("config", func(c *gin.Context) {
		err := c.Bind(config)
		if err != nil {
			c.JSON(200, gin.H{
				"code": -1,
				"msg":  "failed to get new config",
			})
			c.Abort()
		}
		c.JSON(200, gin.H{
			"code": 0,
			"msg":  "success",
		})
	})

	router.GET("config", func(c *gin.Context) {
		c.JSON(200, config)
	})

	err := router.Run(":8080")
	if err != nil {
		log.Fatalf("start config failed, the error is %v", err)
	}
}

func defaultConfig() *Configuration {
	c := new(Configuration)
	c.Auth.WhiteList = []string{
		"/login",
	}
	c.Auth.PublicKey = base64.StdEncoding.EncodeToString(publicKey)
	c.Auth.PrivateKey = base64.StdEncoding.EncodeToString(privateKey)
	c.Auth.ExpireTime = 1200
	c.Auth.RefreshTime = 120
	return c
}

type Configuration struct {
	Auth struct {
		PublicKey   string   `json:"publicKey"`
		PrivateKey  string   `json:"privateKey"`
		WhiteList   []string `json:"whiteList"`
		ExpireTime  int      `json:"expireTime"`
		RefreshTime int      `json:"refreshTime"`
	} `json:"auth"`
}
