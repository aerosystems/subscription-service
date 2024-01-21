package main

import (
	"fmt"
	middleware "github.com/aerosystems/subs-service/internal/middleware"
	"github.com/aerosystems/subs-service/internal/models"
	"github.com/aerosystems/subs-service/internal/presenters/rest"
	RPCServer "github.com/aerosystems/subs-service/internal/presenters/rpc"
	"github.com/aerosystems/subs-service/internal/repository"
	"github.com/aerosystems/subs-service/internal/services"
	"github.com/aerosystems/subs-service/internal/services/payment"
	"github.com/aerosystems/subs-service/internal/validators"
	"github.com/aerosystems/subs-service/pkg/gorm_postgres"
	"github.com/aerosystems/subs-service/pkg/logger"
	"github.com/aerosystems/subs-service/pkg/monobank"
	"github.com/go-playground/validator/v10"
	"github.com/sirupsen/logrus"
	"net/rpc"
	"os"
)

const (
	webPort             = 80
	rpcPort             = 5001
	monobankRedirectUrl = "https://verifire.app/payment/success"
	monobankWebHookUrl  = "https://gw.verifire.app/subs/v1/webhook/monobank"
)

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
	clientGORM.AutoMigrate(models.Subscription{}, models.Invoice{})

	monobankAcquiring := monobank.NewAcquiring("m40_51qVxGawjd3FiJdjhlw")
	monobankStrategy := payment.NewMonobankStrategy(*monobankAcquiring, monobankRedirectUrl, monobankWebHookUrl, monobank.USD)
	paymentMap := map[models.PaymentMethod]payment.AcquiringOperations{
		models.MonobankPaymentMethod: monobankStrategy,
	}

	priceRepo := repository.NewPriceRepo()

	subscriptionRepo := repository.NewSubscriptionRepo(clientGORM)
	subscriptionService := services.NewSubsServiceImpl(subscriptionRepo, priceRepo)

	invoiceRepo := repository.NewInvoiceRepo(clientGORM)
	paymentService := payment.NewServiceImpl(invoiceRepo, priceRepo, paymentMap)

	baseHandler := rest.NewBaseHandler(os.Getenv("APP_ENV"), log.Logger, subscriptionService, paymentService)
	rpcServer := RPCServer.NewSubsServer(rpcPort, log.Logger, subscriptionService)

	accessTokenService := services.NewAccessTokenServiceImpl(os.Getenv("ACCESS_SECRET"))

	oauthMiddleware := middleware.NewOAuthMiddlewareImpl(accessTokenService)
	basicAuthMiddleware := middleware.NewBasicAuthMiddlewareImpl(os.Getenv("BASIC_AUTH_DOCS_USERNAME"), os.Getenv("BASIC_AUTH_DOCS_PASSWORD"))

	app := NewConfig(baseHandler, oauthMiddleware, basicAuthMiddleware)

	e := app.NewRouter()
	middleware.AddLog(e, log.Logger)

	validator := validator.New()
	e.Validator = &validators.CustomValidator{Validator: validator}

	errChan := make(chan error)

	go func() {
		log.Infof("starting subs-service HTTP server on port %d\n", webPort)
		errChan <- e.Start(fmt.Sprintf(":%d", webPort))
	}()

	go func() {
		log.Infof("starting subs-service RPC server on port %d\n", rpcPort)
		errChan <- rpc.Register(rpcServer)
		errChan <- rpcServer.Listen(rpcPort)
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
