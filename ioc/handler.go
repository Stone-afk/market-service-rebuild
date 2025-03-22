package ioc

import (
	"github.com/gin-gonic/gin"
	"market-service/handler"
)

func InitWebServer(mdls []gin.HandlerFunc,
	marketHdl *handler.MarketHandler) *gin.Engine {
	server := gin.Default()
	server.Use(mdls...)
	marketHdl.RegisterRoutes(server)
	return server
}
