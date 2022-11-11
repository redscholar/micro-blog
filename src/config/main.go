package config

import (
	_ "embed"
	"encoding/base64"
	"encoding/json"
	log "go-micro.dev/v4/logger"
	"net/http"
)

//go:embed public_key.pem
var publicKey []byte

//go:embed private_key.pem
var privateKey []byte

func HttpConfig() {
	c := new(Configuration)
	c.Auth.WhiteList = []string{
		"/signIn",
		"/signUp",
	}
	c.Auth.PublicKey = base64.StdEncoding.EncodeToString(publicKey)
	c.Auth.PrivateKey = base64.StdEncoding.EncodeToString(privateKey)
	c.Auth.ExpireTime = 1200
	c.Auth.RefreshTime = 120
	handler := http.NewServeMux()

	handler.HandleFunc("/config", func(writer http.ResponseWriter, request *http.Request) {
		writer.Header().Set("Content-Type", "Application/json")
		data, _ := json.Marshal(c)
		writer.Write(data)
	})
	if err := http.ListenAndServe(":8080", handler); err != nil {
		log.Errorf("start config failed, the error is %v", err)
	}
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
