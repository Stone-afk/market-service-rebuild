package xapp

import (
	"market-service/pkg/xgin"
	"market-service/pkg/xgrpc/server"
)

type App struct {
	GRPCServer *server.Server
	HTTPServer *xgin.Server
}
