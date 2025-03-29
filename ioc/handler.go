package ioc

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"market-service/handler"
	"market-service/pkg/xgin"
)

func InitHttpServer(
	marketHdl *handler.MarketHandler, mdls ...gin.HandlerFunc) *xgin.Server {
	engine := gin.Default()
	engine.Use(mdls...)
	marketHdl.RegisterRoutes(engine)
	addr := viper.GetString("http.addr")
	return &xgin.Server{
		Addr:   addr,
		Engine: engine,
	}
}
