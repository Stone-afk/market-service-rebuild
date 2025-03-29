package main

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"log"
	"market-service/grpc"
	"market-service/handler"
	"market-service/ioc"
	"market-service/pkg/xapp"
	"market-service/repository"
	"market-service/repository/dao"
	"market-service/service"
)

func initViperWatch() {
	cfile := pflag.String("config",
		"config/config.yaml", "配置文件路径")
	pflag.Parse()
	// 直接指定文件路径
	viper.SetConfigFile(*cfile)
	// 实时监听配置变更
	viper.WatchConfig()
	viper.OnConfigChange(func(in fsnotify.Event) {
		fmt.Println(in.Name, in.Op)
		fmt.Println(viper.GetString("db.dsn"))
	})
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
}

func main() {
	initViperWatch()
	app := InitApp()
	go func() {
		err := app.HTTPServer.Start()
		if err != nil {
			log.Println(err)
			panic(err)
		}
	}()
	err := app.GRPCServer.Serve()
	if err != nil {
		log.Println(err)
		panic(err)
	}
}

func InitApp() *xapp.App {
	db := ioc.InitDB("mysql")
	logger := ioc.InitLogger()
	etcdClient := ioc.InitEtcdClient()
	marketPriceDAO := dao.NewMarketPriceDAO(db)
	officialCoinRateDAO := dao.NewOfficialCoinRateDAO(db)
	marketPriceRepo := repository.NewMarketPriceRepository(marketPriceDAO)
	officialCoinRateRepo := repository.NewOfficialCoinRateRepository(officialCoinRateDAO)
	marketPricesService := service.NewMarketPricesService(marketPriceRepo)
	officialCoinRateService := service.NewOfficialCoinRateService(officialCoinRateRepo)
	marketHandler := handler.NewMarketHandler(marketPricesService, officialCoinRateService)
	httpServer := ioc.InitHttpServer(marketHandler)
	marketServiceServer := grpc.NewMarketServiceServer(marketPricesService, officialCoinRateService)
	gRPCServer := ioc.InitGRPCServer(logger, etcdClient, marketServiceServer)
	return &xapp.App{
		HTTPServer: httpServer,
		GRPCServer: gRPCServer,
	}
}
