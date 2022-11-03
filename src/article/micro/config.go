package micro

import (
	"github.com/go-micro/plugins/v4/config/source/url"
	"go-micro.dev/v4/config"
	log "go-micro.dev/v4/logger"
	"time"
)

var cfg, _ = config.NewConfig()

func initConfig() {
	for {
		err := cfg.Load(url.NewSource())
		//err := cfg.Load(etcd.NewSource(etcd.WithAddress(etcdAddr...)))
		if err != nil {
			log.Info("load config error the error is:%v , try again ", err)
			time.Sleep(time.Second)
			continue
		}
		break
	}
}
