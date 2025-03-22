package xapp

import (
	"github.com/gin-gonic/gin"
)

type App struct {
	GRPCServer *server.Server
	WebServer  *gin.Engine
}
