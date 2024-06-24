// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"cloud.google.com/go/firestore"
	"context"
	"firebase.google.com/go/auth"
	"github.com/aerosystems/subscription-service/internal/config"
	"github.com/aerosystems/subscription-service/internal/infrastructure/repository/fire"
	"github.com/aerosystems/subscription-service/internal/infrastructure/repository/memory"
	"github.com/aerosystems/subscription-service/internal/models"
	"github.com/aerosystems/subscription-service/internal/presenters/http"
	"github.com/aerosystems/subscription-service/internal/presenters/http/handlers"
	"github.com/aerosystems/subscription-service/internal/presenters/http/handlers/payment"
	"github.com/aerosystems/subscription-service/internal/presenters/http/handlers/subscription"
	"github.com/aerosystems/subscription-service/internal/presenters/http/middleware"
	"github.com/aerosystems/subscription-service/internal/presenters/rpc"
	"github.com/aerosystems/subscription-service/internal/usecases"
	"github.com/aerosystems/subscription-service/pkg/firebase"
	"github.com/aerosystems/subscription-service/pkg/logger"
	"github.com/aerosystems/subscription-service/pkg/monobank"
	"github.com/sirupsen/logrus"
)

// Injectors from wire.go:

//go:generate wire
func InitApp() *App {
	logger := ProvideLogger()
	logrusLogger := ProvideLogrusLogger(logger)
	config := ProvideConfig()
	client := ProvideFirebaseAuthClient(config)
	firebaseAuth := ProvideFirebaseAuthMiddleware(client)
	xApiKeyAuth := ProvideXAPiKeyMiddleware(config)
	baseHandler := ProvideBaseHandler(logrusLogger, config)
	firestoreClient := ProvideFirestoreClient(config)
	subscriptionRepo := ProvideSubscriptionRepo(firestoreClient)
	subscriptionUsecase := ProvideSubscriptionUsecase(subscriptionRepo)
	handler := ProvideSubscriptionHandler(baseHandler, subscriptionUsecase)
	invoiceRepo := ProvideInvoiceRepo(firestoreClient)
	priceRepo := ProvidePriceRepo()
	acquiring := ProvideMonobankAcquiring(config)
	monobankStrategy := ProvideMonobankStrategy(acquiring, config)
	v := ProvidePaymentMap(monobankStrategy)
	paymentUsecase := ProvidePaymentUsecase(invoiceRepo, priceRepo, v)
	paymentHandler := ProvidePaymentHandler(baseHandler, paymentUsecase)
	server := ProvideHttpServer(logrusLogger, firebaseAuth, xApiKeyAuth, handler, paymentHandler)
	rpcServerServer := ProvideRpcServer(logrusLogger, subscriptionUsecase)
	app := ProvideApp(logrusLogger, config, server, rpcServerServer)
	return app
}

func ProvideApp(log *logrus.Logger, cfg *config.Config, httpServer *HttpServer.Server, rpcServer *RpcServer.Server) *App {
	app := NewApp(log, cfg, httpServer, rpcServer)
	return app
}

func ProvideLogger() *logger.Logger {
	loggerLogger := logger.NewLogger()
	return loggerLogger
}

func ProvideConfig() *config.Config {
	configConfig := config.NewConfig()
	return configConfig
}

func ProvideRpcServer(log *logrus.Logger, subscriptionUsecase RpcServer.SubscriptionUsecase) *RpcServer.Server {
	server := RpcServer.NewServer(log, subscriptionUsecase)
	return server
}

func ProvidePaymentHandler(baseHandler *handlers.BaseHandler, paymentUsecase handlers.PaymentUsecase) *payment.Handler {
	handler := payment.NewPaymentHandler(baseHandler, paymentUsecase)
	return handler
}

func ProvideSubscriptionHandler(baseHandler *handlers.BaseHandler, subscriptionUsecase handlers.SubscriptionUsecase) *subscription.Handler {
	handler := subscription.NewSubscriptionHandler(baseHandler, subscriptionUsecase)
	return handler
}

func ProvidePaymentUsecase(invoiceRepo usecases.InvoiceRepository, priceRepo usecases.PriceRepository, strategies map[models.PaymentMethod]usecases.AcquiringOperations) *usecases.PaymentUsecase {
	paymentUsecase := usecases.NewPaymentUsecase(invoiceRepo, priceRepo, strategies)
	return paymentUsecase
}

func ProvideSubscriptionUsecase(subscriptionRepo usecases.SubscriptionRepository) *usecases.SubscriptionUsecase {
	subscriptionUsecase := usecases.NewSubscriptionUsecase(subscriptionRepo)
	return subscriptionUsecase
}

func ProvidePriceRepo() *memory.PriceRepo {
	priceRepo := memory.NewPriceRepo()
	return priceRepo
}

func ProvideSubscriptionRepo(client *firestore.Client) *fire.SubscriptionRepo {
	subscriptionRepo := fire.NewSubscriptionRepo(client)
	return subscriptionRepo
}

func ProvideInvoiceRepo(client *firestore.Client) *fire.InvoiceRepo {
	invoiceRepo := fire.NewInvoiceRepo(client)
	return invoiceRepo
}

// wire.go:

func ProvideHttpServer(log *logrus.Logger, firebaseAuthMiddleware *middleware.FirebaseAuth, xApiKeyAuthMiddleware *middleware.XApiKeyAuth, subscriptionHandler *subscription.Handler, paymentHandler *payment.Handler) *HttpServer.Server {
	return HttpServer.NewServer(log, firebaseAuthMiddleware, xApiKeyAuthMiddleware, subscriptionHandler, paymentHandler)
}

func ProvideLogrusLogger(log *logger.Logger) *logrus.Logger {
	return log.Logger
}

func ProvideBaseHandler(log *logrus.Logger, cfg *config.Config) *handlers.BaseHandler {
	return handlers.NewBaseHandler(log, cfg.Mode)
}

func ProvidePaymentMap(monobankStrategy *usecases.MonobankStrategy) map[models.PaymentMethod]usecases.AcquiringOperations {
	return map[models.PaymentMethod]usecases.AcquiringOperations{models.MonobankPaymentMethod: monobankStrategy}
}

func ProvideMonobankStrategy(acquiring *monobank.Acquiring, cfg *config.Config) *usecases.MonobankStrategy {
	return usecases.NewMonobankStrategy(acquiring, cfg.MonobankRedirectUrl, cfg.MonobankWebHookUrl, monobank.USD)
}

func ProvideMonobankAcquiring(cfg *config.Config) *monobank.Acquiring {
	return monobank.NewAcquiring(cfg.MonobankToken)
}

func ProvideFirestoreClient(cfg *config.Config) *firestore.Client {
	ctx := context.Background()
	client, err := firestore.NewClient(ctx, cfg.GcpProjectId)
	if err != nil {
		panic(err)
	}
	return client
}

func ProvideFirebaseAuthMiddleware(client *auth.Client) *middleware.FirebaseAuth {
	return middleware.NewFirebaseAuth(client)
}

func ProvideXAPiKeyMiddleware(cfg *config.Config) *middleware.XApiKeyAuth {
	xApiKeyAuthMiddleware, err := middleware.NewXApiKeyAuth(cfg.ApiKey)
	if err != nil {
		panic(err)
	}
	return xApiKeyAuthMiddleware
}

func ProvideFirebaseAuthClient(cfg *config.Config) *auth.Client {
	app, err := firebaseApp.NewApp(cfg.GcpProjectId, cfg.GoogleApplicationCredentials)
	if err != nil {
		panic(err)
	}
	return app.Client
}
