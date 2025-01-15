package main

import (
	"github.com/aerosystems/subscription-service/internal/common/config"
	GRPCServer "github.com/aerosystems/subscription-service/internal/presenters/grpc"
	HttpServer "github.com/aerosystems/subscription-service/internal/presenters/http"
	"github.com/sirupsen/logrus"
)

type App struct {
	log        *logrus.Logger
	cfg        *config.Config
	httpServer *HttpServer.Server
	grpcServer *GRPCServer.Server
}

func NewApp(
	log *logrus.Logger,
	cfg *config.Config,
	httpServer *HttpServer.Server,
	grpcServer *GRPCServer.Server,
) *App {
	if log == nil {
		panic("log is required")
	}
	if cfg == nil {
		panic("cfg is required")
	}
	if httpServer == nil {
		panic("httpServer is required")
	}
	if grpcServer == nil {
		panic("grpcServer is required")
	}
	return &App{
		log:        log,
		cfg:        cfg,
		httpServer: httpServer,
		grpcServer: grpcServer,
	}
}
