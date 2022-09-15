package micro

import (
	"context"
	"crypto/tls"
	"fmt"
	"github.com/go-micro/plugins/v4/auth/jwt"
	brokerGrpc "github.com/go-micro/plugins/v4/broker/grpc"
	cacheRedis "github.com/go-micro/plugins/v4/cache/redis"
	registryEtcd "github.com/go-micro/plugins/v4/registry/etcd"
	"github.com/go-micro/plugins/v4/store/redis"
	"go-micro.dev/v4"
	"go-micro.dev/v4/auth"
	"go-micro.dev/v4/broker"
	"go-micro.dev/v4/cache"
	"go-micro.dev/v4/client"
	"go-micro.dev/v4/debug/profile"
	"go-micro.dev/v4/debug/profile/http"
	"go-micro.dev/v4/debug/trace"
	log "go-micro.dev/v4/logger"
	"go-micro.dev/v4/registry"
	"go-micro.dev/v4/runtime"
	"go-micro.dev/v4/server"
	"go-micro.dev/v4/store"
	"go-micro.dev/v4/transport"
	"time"
)

var Service micro.Service

const (
	service    = "web"
	version    = "latest"
	address    = ":37100"
	authHeader = "Authorization"
)

type scheduler struct {
	stop chan bool
}

func (s scheduler) Notify() (<-chan runtime.Event, error) {
	var a = make(chan runtime.Event)
	go func() {
		ticker := time.NewTicker(1 * time.Minute)

		for {
			select {
			case <-ticker.C:
				a <- runtime.Event{}
			case <-s.stop:
				return
			}
		}
	}()
	return a, nil
}

func (s scheduler) Close() error {
	s.stop <- true
	return nil
}

func init() {
	initCmd()
	initConfig()
	etcdCert, _ := tls.LoadX509KeyPair(etcdCertFile, etcdKeyFile)
	transportCert, _ := tls.LoadX509KeyPair(transportCertFile, transportKeyFile)
	brokerCert, _ := tls.LoadX509KeyPair(brokerCertFile, brokerKeyFile)
	auth := jwt.NewAuth(
		auth.Addrs(),
		auth.Namespace("blog"),
		auth.PublicKey(cfg.Get("auth", "publicKey").String("")),
		auth.PrivateKey(cfg.Get("auth", "privateKey").String("")),
		//auth.Credentials("root", "123"),
		//auth.ClientToken(&auth.Token{}),
	)
	Service = micro.NewService(
		micro.Server(
			server.NewServer(
				server.Name(service),
				//server.Id(server.DefaultId),
				server.Version(version),
				server.Address(address),
				//server.Advertise(server.DefaultAddress),
				//server.Broker(broker.DefaultBroker),
				//server.Codec("application/text", func(closer io.ReadWriteCloser) codec.Codec {
				//	return &bytes.Codec{}
				//}),
				//server.Registry(registry.DefaultRegistry),
				//server.Tracer(trace.DefaultTracer),
				server.Metadata(map[string]string{"description": "web ui and route"}),
				//server.RegisterTTL(server.DefaultRegisterTTL),
				server.Context(context.WithValue(context.Background(), "isReady", true)),
				server.RegisterCheck(func(ctx context.Context) error {
					if !ctx.Value("isReady").(bool) {
						return fmt.Errorf("server not ready to registry")
					}
					return nil
				}),
				//server.RegisterInterval(server.DefaultRegisterInterval),
				//server.TLSConfig(&tls.Config{Certificates: []tls.Certificate{transportCert}}),
				//server.WithRouter(server.DefaultRouter),
				server.Wait(nil),
				server.WrapHandler(func(handlerFunc server.HandlerFunc) server.HandlerFunc {
					return func(ctx context.Context, req server.Request, rsp interface{}) error { // 支持链路追踪
						newCtx, s := trace.DefaultTracer.Start(ctx, "web")
						s.Type = trace.SpanTypeRequestInbound
						defer trace.DefaultTracer.Finish(s)
						return handlerFunc(newCtx, req, rsp)
					}
				}),
				//server.WrapHandler(func(handlerFunc server.HandlerFunc) server.HandlerFunc { // 支持auth认证
				//	return func(ctx context.Context, req server.Request, rsp interface{}) error {
				//		token := req.Header()[authHeader]
				//		account, err := auth.Inspect(token)
				//		if err != nil {
				//			return handlerFunc(ctx, req, rsp)
				//		}
				//		return handlerFunc(context.WithValue(ctx, "account", account), req, rsp)
				//	}
				//}),
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
		//micro.Name(service),
		//micro.Version(version),
		//micro.Address(":7001"),
		//micro.Tracer(trace.DefaultTracer),
		//micro.RegisterTTL(10*time.Second),
		//micro.RegisterInterval(10*time.Second),
		//micro.Metadata(map[string]string{"description2": "web ui and route"}),
		//micro.WrapHandler(),
		//micro.WrapSubscriber(),
		micro.Client(
			client.NewClient(
				//client.Broker(broker.DefaultBroker),
				//client.Codec("application/text", func(closer io.ReadWriteCloser) codec.Codec {
				//	return &bytes.Codec{}
				//}),
				//client.ContentType(client.DefaultContentType),
				//client.PoolSize(client.DefaultPoolSize),
				//client.Selector(selector.NewSelector(selector.SetStrategy(selector.Random))),
				//client.Registry(registry.DefaultRegistry),
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
		//micro.Selector(selector.DefaultSelector),
		//micro.WrapClient(),
		//micro.WrapCall(),
		micro.Broker(
			brokerGrpc.NewBroker(
				broker.Addrs(),
				//broker.Codec(json.Marshaler{}),
				//broker.ErrorHandler(func(event broker.Event) error {
				//	return nil
				//}),
				//broker.Registry(registry.DefaultRegistry),
				broker.Secure(true),
				broker.TLSConfig(&tls.Config{Certificates: []tls.Certificate{brokerCert}}),
			),
		),
		micro.Registry(
			registryEtcd.NewRegistry( // 设置etcd注册中心
				registry.Addrs(etcdAddr...), // etcd 地址。默认127.0.0.1:2379
				//registry.Timeout(10*time.Second), // 超时时间
				registry.Secure(true), // 是否启用tls
				registry.TLSConfig(&tls.Config{Certificates: []tls.Certificate{etcdCert}}), // tls设置
			),
		),
		micro.Transport(
			transport.NewHTTPTransport(
				transport.Addrs(),
				transport.Codec(nil),
				transport.Timeout(transport.DefaultDialTimeout),
				transport.Secure(true),
				transport.TLSConfig(&tls.Config{Certificates: []tls.Certificate{transportCert}}),
			),
		),

		micro.Auth(auth),
		micro.Cache(
			cacheRedis.NewCache(
				//cache.Expiration(10*time.Second),
				//cache.Items(nil),
				cache.WithAddress(redisAddr[0]),
				//cache.WithContext(context.Background()),
			),
		),
		micro.Store(
			redis.NewStore(
				store.Nodes(redisAddr...),
				//store.Database("blog"),
				store.Table("web"),
				//store.WithContext(context.Background()),
				//store.WithClient(nil),
			),
		),
		micro.Config(cfg),
		micro.Runtime(
			runtime.NewRuntime(
				runtime.WithSource("blog"),
				runtime.WithScheduler(&scheduler{}),
				runtime.WithType("service"),
				runtime.WithImage("web:1.0"),
				runtime.WithClient(nil),
			),
		),

		micro.Profile(http.NewProfile(profile.Name("web"))),

		//micro.BeforeStart(func() error {
		//	log.Info("before start 1")
		//	return nil
		//}),
		//micro.AfterStart(func() error {
		//	log.Info("after start 1")
		//	return nil
		//}),
		//micro.BeforeStart(func() error {
		//	log.Info("before stop 1")
		//	return nil
		//}),
		//micro.BeforeStart(func() error {
		//	log.Info("after stop 1")
		//	return nil
		//}),
		//micro.HandleSignal(true),
		micro.Registry(
			registryEtcd.NewRegistry( // 设置etcd注册中心
				registry.Addrs(etcdAddr...), // etcd 地址。默认127.0.0.1:2379
				//registry.Timeout(10*time.Second), // 超时时间
				registry.Secure(true), // 是否启用tls
				//registry.TLSConfig(&tls.Config{Certificates: []tls.Certificate{etcdCert}}), // tls设置
			),
		),
		micro.Transport(
			transport.NewHTTPTransport(
				transport.Addrs(),
				transport.Codec(nil),
				transport.Timeout(transport.DefaultDialTimeout),
				transport.Secure(true),
				//transport.TLSConfig(&tls.Config{Certificates: []tls.Certificate{transportCert}}),
			),
		),
	)

}
