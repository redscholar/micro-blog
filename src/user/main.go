package main

import (
	"context"
	"crypto/tls"
	"github.com/go-micro/plugins/v4/registry/etcd"
	"go-micro.dev/v4"
	"go-micro.dev/v4/client"
	"go-micro.dev/v4/debug/trace"
	log "go-micro.dev/v4/logger"
	"go-micro.dev/v4/metadata"
	"go-micro.dev/v4/registry"
	"go-micro.dev/v4/transport"
	"os"
	"time"
	webprotoc "web/proto"
)

var (
	service = "user"
	version = "latest"
)

func main() {
	caCert, _ := tls.LoadX509KeyPair(os.Getenv("CA_CERT_FILE"), os.Getenv("CA_KEY_FILE"))

	// Create service
	srv := micro.NewService(
		micro.Name(service),
		micro.Version(version),
		micro.Client(client.NewClient(client.Transport(transport.NewHTTPTransport(transport.Secure(true), transport.TLSConfig(&tls.Config{Certificates: []tls.Certificate{caCert}}))))),

		micro.Registry(
			etcd.NewRegistry( // 设置etcd注册中心
				registry.Addrs(),                 // etcd 地址。默认127.0.0.1:2379
				registry.Timeout(10*time.Second), // 超时时间
				registry.Secure(true),            // 是否启用tls
				registry.TLSConfig(&tls.Config{Certificates: []tls.Certificate{caCert}}), // tls设置
			),
		),
	)
	srv.Init()
	rsp := new(webprotoc.CallResponse)
	ctx, span := trace.DefaultTracer.Start(context.Background(), "web")
	span.Type = trace.SpanTypeRequestInbound
	defer trace.DefaultTracer.Finish(span)
	srv.Client()
	md, _ := metadata.FromContext(ctx)
	log.Infof("trace:%v", md["Micro-Trace-Id"])
	log.Infof("span:%v", md["Micro-Span-Id"])
	err := srv.Client().Call(ctx, srv.Client().NewRequest("web", "Web.Call", &webprotoc.CallRequest{Name: "aaa"}), rsp)
	//err = srv.Client().Call(ctx, srv.Client().NewRequest("web", "Web.Call", &webprotoc.CallRequest{Name: "aaa"}), rsp)
	if err != nil {
		log.Fatal(err)
	}
	//Register handler
	//pb.RegisterUserHandler(srv.Server(), new(handler.User))
	//// Run service
	//if err := srv.Run(); err != nil {
	//	log.Fatal(err)
	//}
}
