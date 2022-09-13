package main

import (
	"context"
	"crypto/tls"
	"fmt"
	"github.com/go-micro/plugins/v4/registry/etcd"
	"go-micro.dev/v4"
	"go-micro.dev/v4/auth"
	"go-micro.dev/v4/broker"
	"go-micro.dev/v4/client"
	"go-micro.dev/v4/codec"
	"go-micro.dev/v4/codec/bytes"
	"go-micro.dev/v4/debug/trace"
	log "go-micro.dev/v4/logger"
	"go-micro.dev/v4/registry"
	"go-micro.dev/v4/selector"
	"go-micro.dev/v4/server"
	"go-micro.dev/v4/transport"
	"io"
	"os"
	"time"
	"web/handler"
	pb "web/proto"
)

const (
	service    = "web"
	version    = "latest"
	authHeader = "token"
)

func main() {
	caCert, _ := tls.LoadX509KeyPair(os.Getenv("CA_CERT_FILE"), os.Getenv("CA_KEY_FILE"))
	transportCert, _ := tls.LoadX509KeyPair(os.Getenv("TRANSPORT_CERT_FILE"), os.Getenv("TRANSPORT_KEY_FILE"))
	// Create service
	srv := micro.NewService(
		micro.Server(
			server.NewServer(
				server.Name(service),
				server.Id(server.DefaultId),
				server.Version(version),
				server.Address(server.DefaultAddress),
				//server.Advertise(server.DefaultAddress),
				server.Broker(broker.DefaultBroker),
				server.Codec("application/text", func(closer io.ReadWriteCloser) codec.Codec {
					return &bytes.Codec{}
				}),
				server.Context(context.Background()),
				server.Registry(registry.DefaultRegistry),
				//server.Tracer(trace.DefaultTracer),
				server.Metadata(map[string]string{"description": "web ui and route"}),
				server.RegisterTTL(server.DefaultRegisterTTL),
				server.Context(context.WithValue(context.Background(), "isReady", true)),
				server.RegisterCheck(func(ctx context.Context) error {
					if !ctx.Value("isReady").(bool) {
						return fmt.Errorf("server not ready to registry")
					}
					return nil
				}),
				server.RegisterInterval(server.DefaultRegisterInterval),
				server.TLSConfig(&tls.Config{Certificates: []tls.Certificate{transportCert}}),
				server.WithRouter(server.DefaultRouter),
				server.Wait(nil),
				server.WrapHandler(func(handlerFunc server.HandlerFunc) server.HandlerFunc {
					return func(ctx context.Context, req server.Request, rsp interface{}) error { // 支持链路追踪
						newCtx, s := trace.DefaultTracer.Start(ctx, "web")
						s.Type = trace.SpanTypeRequestInbound
						defer trace.DefaultTracer.Finish(s)
						return handlerFunc(newCtx, req, rsp)
					}
				}),
				server.WrapHandler(func(handlerFunc server.HandlerFunc) server.HandlerFunc { // 支持auth认证
					return func(ctx context.Context, req server.Request, rsp interface{}) error {
						token := req.Header()[authHeader]
						account, err := auth.DefaultAuth.Inspect(token)
						if err != nil {
							return err
						}
						return handlerFunc(context.WithValue(ctx, "account", account), req, rsp)
					}
				}),
				server.WrapSubscriber(func(subscriberFunc server.SubscriberFunc) server.SubscriberFunc {
					return func(ctx context.Context, msg server.Message) error {
						newCtx, s := trace.DefaultTracer.Start(ctx, "web")
						s.Type = trace.SpanTypeRequestInbound
						defer trace.DefaultTracer.Finish(s)
						return subscriberFunc(newCtx, msg)
					}
				}),
			),
		),
		micro.Name(service),
		micro.Version(version),
		micro.Client(
			client.NewClient(
				client.Broker(broker.DefaultBroker),
				client.Codec("application/text", func(closer io.ReadWriteCloser) codec.Codec {
					return &bytes.Codec{}
				}),
				client.ContentType(client.DefaultContentType),
				client.PoolSize(0),
				client.Selector(selector.NewSelector(selector.SetStrategy(selector.RoundRobin))),
				client.Registry(registry.DefaultRegistry),
				client.Wrap(func(c client.Client) client.Client { // 构造客户端
					if c.Options().Registry == nil {
						return nil
					}
					return c
				}),
				client.WrapCall(func(callFunc client.CallFunc) client.CallFunc {
					return func(ctx context.Context, node *registry.Node, req client.Request, rsp interface{}, opts client.CallOptions) error { // 设置请求链路追踪
						newCtx, s := trace.DefaultTracer.Start(ctx, "web")
						s.Type = trace.SpanTypeRequestInbound
						defer trace.DefaultTracer.Finish(s)
						return callFunc(newCtx, node, req, rsp, opts)
					}
				}),
				client.Backoff(func(ctx context.Context, req client.Request, attempts int) (time.Duration, error) { // 打印请求
					log.Infof("attempts %v, the req is %v", attempts, req.Body())
					return 0, nil
				}),
				client.Retry(func(ctx context.Context, req client.Request, retryCount int, err error) (bool, error) {
					return err != nil, nil
				}),
				client.Retries(client.DefaultRetries),
				client.RequestTimeout(client.DefaultRequestTimeout),
				client.StreamTimeout(client.DefaultRequestTimeout),
				client.DialTimeout(transport.DefaultDialTimeout),
				client.WithRouter(nil),
			),
		),
		micro.Selector(selector.DefaultSelector),

		micro.Registry(
			etcd.NewRegistry( // 设置etcd注册中心
				registry.Addrs(),                 // etcd 地址。默认127.0.0.1:2379
				registry.Timeout(10*time.Second), // 超时时间
				registry.Secure(true),            // 是否启用tls
				registry.TLSConfig(&tls.Config{Certificates: []tls.Certificate{caCert}}), // tls设置
			),
		),
		//micro.Auth(jwt.NewAuth()),
		//micro.Profile(http.NewProfile()),
		//micro.Server(server.NewServer()), // 默认mutp
		//micro.Client(client.NewClient()), // 默认mutp
		//micro.Selector(),
		//micro.Transport(),
		//micro.Broker(),
		//micro.Cache(),
		//micro.Config(),
		//micro.Context(ctx), // 默认context.Background()
	)
	srv.Init()

	//Register handler
	pb.RegisterWebHandler(srv.Server(), new(handler.Web))
	// Run service
	if err := srv.Run(); err != nil {
		log.Fatal(err)
	}
}
