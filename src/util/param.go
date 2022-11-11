package util

import (
	"github.com/urfave/cli/v2"
	_ "go-micro.dev/v4/util/cmd"
	"strings"
)

var ComParam = new(struct {
	CaCertFile        string
	CaKeyFile         string
	EtcdCertFile      string
	EtcdKeyFile       string
	TransportCertFile string
	TransportKeyFile  string
	BrokerCertFile    string
	BrokerKeyFile     string

	MongoUser     string
	MongoPassword string
	MongoUrl      string

	EtcdAddr  []string
	RedisAddr []string
})
var ComFlags = []cli.Flag{&cli.StringFlag{
	Name:    "ca-cert-file",
	Usage:   "ca cert file path",
	EnvVars: []string{"CA_CERT_FILE"},
	Value:   "/Users/liujian/work/000-learn/code/micro-blog/cert/ca.crt",
}, &cli.StringFlag{
	Name:    "ca-key-file",
	Usage:   "ca key file path",
	EnvVars: []string{"CA_KEY_FILE"},
	Value:   "/Users/liujian/work/000-learn/code/micro-blog/cert/ca.key",
}, &cli.StringFlag{
	Name:    "etcd-cert-file",
	Usage:   "etcd cert file path",
	EnvVars: []string{"ETCD_CERT_FILE"},
	Value:   "/Users/liujian/work/000-learn/code/micro-blog/cert/etcd.crt",
}, &cli.StringFlag{
	Name:    "etcd-key-file",
	Usage:   "etcd key file path",
	EnvVars: []string{"ETCD_KEY_FILE"},
	Value:   "/Users/liujian/work/000-learn/code/micro-blog/cert/etcd.key",
}, &cli.StringFlag{
	Name:    "transport-cert-file",
	Usage:   "transport cert file path",
	EnvVars: []string{"TRANSPORT_CERT_FILE"},
	Value:   "/Users/liujian/work/000-learn/code/micro-blog/cert/transport.crt",
}, &cli.StringFlag{
	Name:    "transport-key-file",
	Usage:   "transport key file path",
	EnvVars: []string{"TRANSPORT_KEY_FILE"},
	Value:   "/Users/liujian/work/000-learn/code/micro-blog/cert/transport.key",
}, &cli.StringFlag{
	Name:    "broker-cert-file",
	Usage:   "broker cert file path",
	EnvVars: []string{"BROKER_CERT_FILE"},
	Value:   "/Users/liujian/work/000-learn/code/micro-blog/cert/broker.crt",
}, &cli.StringFlag{
	Name:    "broker-key-file",
	Usage:   "broker key file path",
	EnvVars: []string{"BROKER_KEY_FILE"},
	Value:   "/Users/liujian/work/000-learn/code/micro-blog/cert/broker.key",
}, &cli.StringFlag{
	Name:    "etcd-addr",
	Usage:   "the address to connect etcd. e.g: 127.0.0.1:2379,127.0.0.2:2379",
	EnvVars: []string{"ETCD_ADDR"},
	Value:   "127.0.0.1:12379",
}, &cli.StringFlag{
	Name:    "redis-addr",
	Usage:   "the address to connect redis. e.g: redis://127.0.0.1:6379,redis://127.0.0.2:6379",
	EnvVars: []string{"REDIS_ADDR"},
	Value:   "redis://127.0.0.1:6379",
}, &cli.StringFlag{
	Name:    "mongo-user",
	Usage:   "the user to access mongo. e.g: admin",
	EnvVars: []string{"MONGO_USER"},
	Value:   "admin",
}, &cli.StringFlag{
	Name:    "mongo-password",
	Usage:   "the password to access mongo. e.g: 1234",
	EnvVars: []string{"MONGO_PASSWORD"},
	Value:   "123456",
}, &cli.StringFlag{
	Name:    "mongo-url",
	Usage:   "the url to connect mongo. e.g: mongodb://mongo-service:27017",
	EnvVars: []string{"MONGO_URL"},
	Value:   "mongodb://127.0.0.1:27017",
}}

func InitComParam(c *cli.Context) error {
	ComParam.CaCertFile = c.String("ca-cert-file")
	ComParam.CaKeyFile = c.String("ca-key-file")
	ComParam.EtcdCertFile = c.String("etcd-cert-file")
	ComParam.EtcdKeyFile = c.String("etcd-key-file")
	ComParam.TransportCertFile = c.String("transport-cert-file")
	ComParam.TransportKeyFile = c.String("transport-key-file")
	ComParam.BrokerCertFile = c.String("broker-cert-file")
	ComParam.BrokerKeyFile = c.String("broker-key-file")
	ComParam.EtcdAddr = strings.Split(c.String("etcd-addr"), ",")
	ComParam.RedisAddr = strings.Split(c.String("redis-addr"), ",")

	ComParam.MongoUser = c.String("mongo-user")
	ComParam.MongoPassword = c.String("mongo-password")
	ComParam.MongoUrl = c.String("mongo-url")
	return nil
}
