package ioc

import (
	"github.com/spf13/viper"
	etcdv3 "go.etcd.io/etcd/client/v3"
	"google.golang.org/grpc"
	grpc2 "market-service/grpc"
	"market-service/pkg/logger"
	"market-service/pkg/xgrpc/interceptors/logging"
	"market-service/pkg/xgrpc/server"
)

func InitGRPCServer(l logger.Logger, etcdClient *etcdv3.Client, weServer *grpc2.MarketServiceServer) *server.Server {
	type Config struct {
		Port    int   `yaml:"port"`
		EtcdTTL int64 `yaml:"etcdTTL"`
	}
	var cfg Config
	err := viper.UnmarshalKey("grpc.server", &cfg)
	if err != nil {
		panic(err)
	}
	grpcSrv := grpc.NewServer(grpc.ChainUnaryInterceptor(
		logging.NewLoggerInterceptorBuilder(l).Build(),
	))
	weServer.Register(grpcSrv)
	return server.NewGRPCXServer(grpcSrv, etcdClient, l, cfg.Port, "market", cfg.EtcdTTL)
}
