package main

import (
	"fmt"
	"github.com/aerosystems/subs-service/internal/handlers"
	"github.com/aerosystems/subs-service/internal/models"
	"github.com/aerosystems/subs-service/internal/repository"
	"github.com/aerosystems/subs-service/pkg/gorm_postgres"
	"github.com/aerosystems/subs-service/pkg/logger"
	"github.com/sirupsen/logrus"
	"os"
)

const webPort = 80
const rpcPort = 5001

// @title Subscription Service
// @version 1.0.0
// @description A part of microservice infrastructure, who responsible for user subscriptions

// @contact.name Artem Kostenko
// @contact.url https://github.com/aerosystems

// @license.name Apache 2.0
// @license.url https://www.apache.org/licenses/LICENSE-2.0.html

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
// @description Should contain Access JWT Token, with the Bearer started

// @host gw.verifire.com/subs
// @schemes https
// @BasePath /
func main() {
	log := logger.NewLogger(os.Getenv("HOSTNAME"))

	clientGORM := GormPostgres.NewClient(logrus.NewEntry(log.Logger))
	clientGORM.AutoMigrate(models.Subscription{})

	subscriptionRepo := repository.NewSubscriptionRepo(clientGORM)

	baseHandler := handlers.NewBaseHandler(subscriptionRepo)

	app := Config{
		BaseHandler: baseHandler,
	}

	e := app.NewRouter()
	app.AddMiddleware(e, log.Logger)

	errChan := make(chan error)

	go func() {
		log.Infof("starting subs-service HTTP server on port %d\n", webPort)
		errChan <- e.Start(fmt.Sprintf(":%d", webPort))
	}()

	for {
		select {
		case err := <-errChan:
			if err != nil {
				log.Fatal(err)
			}
		}
	}
}
