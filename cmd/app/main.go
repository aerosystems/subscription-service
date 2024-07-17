package main

import (
	"context"
	"golang.org/x/sync/errgroup"
)

// @title Subscription Service
// @version 1.0.0
// @description A part of microservice infrastructure, who responsible for user subscriptions

// @contact.name Artem Kostenko
// @contact.url https://github.com/aerosystems

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
// @description Should contain Access JWT Token, with the Bearer started

// @host gw.verifire.dev/subs
// @schemes https
// @BasePath /
func main() {
	app := InitApp()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	group, ctx := errgroup.WithContext(ctx)

	group.Go(func() error {
		return app.httpServer.Run()
	})

	group.Go(func() error {
		return app.rpcServer.Run()
	})

	group.Go(func() error {
		return app.handleSignals(ctx, cancel)
	})

	if err := group.Wait(); err != nil {
		app.log.Errorf("error occurred: %v", err)
	}
}
