package micro

import (
	"github.com/go-micro/plugins/v4/config/source/url"
	"github.com/google/uuid"
	"go-micro.dev/v4/config"
	"net/http"
)

var cfg, _ = config.NewConfig()

func initConfig() {
	err := cfg.Load(url.NewSource())
	//err := cfg.Load(etcd.NewSource(etcd.WithAddress(etcdAddr...)))
	if err != nil {
		return
	}
	defer cfg.Sync()
	if cfg.Get("auth", "publicKey").String("") == "" {
		cfg.Set(uuid.New().String(), "auth", "publicKey")
	}
	if cfg.Get("auth", "privateKey").String("") == "" {
		cfg.Set(uuid.New().String(), "auth", "privateKey")
	}
	if cfg.Get("auth", "whiteList").StringSlice(nil) == nil {
		cfg.Set([]string{
			"/login",
		}, "auth", "whiteList")
	}
}

func HttpConfig() {
	handler := http.NewServeMux()
	handler.HandleFunc("/config", func(writer http.ResponseWriter, request *http.Request) {
		writer.Header().Set("Content-Type", "Application/json")
		writer.Write([]byte("{}"))
	})
	http.ListenAndServe(":8080", handler)
}

type Configuration struct {
	Auth struct {
		PublicKey  string   `json:"publicKey"`
		PrivateKey string   `json:"privateKey"`
		WhiteList  []string `json:"whiteList"`
	}
}
