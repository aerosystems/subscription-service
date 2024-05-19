package main

import (
	"github.com/aerosystems/subscription-service/internal/config"
	HttpServer "github.com/aerosystems/subscription-service/internal/presenters/http"
	RpcServer "github.com/aerosystems/subscription-service/internal/presenters/rpc"
	"github.com/sirupsen/logrus"
)

type App struct {
	log        *logrus.Logger
	cfg        *config.Config
	httpServer *HttpServer.Server
	rpcServer  *RpcServer.Server
}

func NewApp(
	log *logrus.Logger,
	cfg *config.Config,
	httpServer *HttpServer.Server,
	rpcServer *RpcServer.Server,
) *App {
	return &App{
		log:        log,
		cfg:        cfg,
		httpServer: httpServer,
		rpcServer:  rpcServer,
	}
}
