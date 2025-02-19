package main

import (
	GRPCServer "github.com/aerosystems/subscription-service/internal/ports/grpc"
	HTTPServer "github.com/aerosystems/subscription-service/internal/ports/http"
	"github.com/sirupsen/logrus"
)

type App struct {
	log        *logrus.Logger
	cfg        *Config
	httpServer *HTTPServer.Server
	grpcServer *GRPCServer.Server
}

func NewApp(
	log *logrus.Logger,
	cfg *Config,
	httpServer *HTTPServer.Server,
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
