// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"cloud.google.com/go/firestore"
	"context"
	"firebase.google.com/go/v4/auth"
	"github.com/aerosystems/common-service/clients/gcpclient"
	"github.com/aerosystems/common-service/logger"
	"github.com/aerosystems/common-service/presenters/grpcserver"
	"github.com/aerosystems/common-service/presenters/httpserver"
	"github.com/aerosystems/subscription-service/internal/adapters"
	"github.com/aerosystems/subscription-service/internal/ports/grpc"
	"github.com/aerosystems/subscription-service/internal/ports/http"
	"github.com/aerosystems/subscription-service/internal/usecases"
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
	firestoreClient := ProvideFirestoreClient(config)
	invoiceRepo := ProvideInvoiceRepo(firestoreClient)
	priceRepo := ProvidePriceRepo()
	acquiring := ProvideMonobankAcquiring(config)
	monobankAcquiring := ProvideMonobankStrategy(acquiring, config)
	paymentUsecase := ProvidePaymentUsecase(invoiceRepo, priceRepo, monobankAcquiring)
	subscriptionRepo := ProvideSubscriptionRepo(firestoreClient)
	subscriptionUsecase := ProvideSubscriptionUsecase(subscriptionRepo)
	handler := ProvideHandler(paymentUsecase, subscriptionUsecase)
	server := ProvideHTTPServer(config, logrusLogger, firebaseAuth, handler)
	subscriptionService := ProvideSubscriptionService(subscriptionUsecase)
	grpcServerServer := ProvideGRPCServer(config, logrusLogger, subscriptionService)
	app := ProvideApp(logrusLogger, config, server, grpcServerServer)
	return app
}

func ProvideApp(log *logrus.Logger, cfg *Config, httpServer *HTTPServer.Server, gpcServer *GRPCServer.Server) *App {
	app := NewApp(log, cfg, httpServer, gpcServer)
	return app
}

func ProvideLogger() *logger.Logger {
	loggerLogger := logger.NewLogger()
	return loggerLogger
}

func ProvideConfig() *Config {
	config := NewConfig()
	return config
}

func ProvideHandler(paymentUsecase HTTPServer.PaymentUsecase, subscriptionUsecase HTTPServer.SubscriptionUsecase) *HTTPServer.Handler {
	handler := HTTPServer.NewHandler(subscriptionUsecase, paymentUsecase)
	return handler
}

func ProvideSubscriptionService(subscriptionUsecase GRPCServer.SubscriptionUsecase) *GRPCServer.SubscriptionService {
	subscriptionService := GRPCServer.NewSubscriptionService(subscriptionUsecase)
	return subscriptionService
}

func ProvideSubscriptionUsecase(subscriptionRepo usecases.SubscriptionRepository) *usecases.SubscriptionUsecase {
	subscriptionUsecase := usecases.NewSubscriptionUsecase(subscriptionRepo)
	return subscriptionUsecase
}

func ProvidePriceRepo() *adapters.PriceRepo {
	priceRepo := adapters.NewPriceRepo()
	return priceRepo
}

func ProvideSubscriptionRepo(client *firestore.Client) *adapters.SubscriptionRepo {
	subscriptionRepo := adapters.NewSubscriptionRepo(client)
	return subscriptionRepo
}

func ProvideInvoiceRepo(client *firestore.Client) *adapters.InvoiceRepo {
	invoiceRepo := adapters.NewInvoiceRepo(client)
	return invoiceRepo
}

// wire.go:

func ProvideHTTPServer(cfg *Config, log *logrus.Logger, firebaseAuth *HTTPServer.FirebaseAuth, handler *HTTPServer.Handler) *HTTPServer.Server {
	return HTTPServer.NewHTTPServer(&HTTPServer.Config{
		Config: httpserver.Config{
			Host: cfg.Host,
			Port: cfg.Port,
		},
		Mode: cfg.Mode,
	}, log, firebaseAuth, handler)
}

func ProvideGRPCServer(cfg *Config, log *logrus.Logger, subscriptionService *GRPCServer.SubscriptionService) *GRPCServer.Server {
	return GRPCServer.NewGRPCServer(&grpcserver.Config{
		Host: cfg.Host,
		Port: cfg.Port,
	}, log, subscriptionService)
}

func ProvideLogrusLogger(log *logger.Logger) *logrus.Logger {
	return log.Logger
}

func ProvidePaymentUsecase(invoiceRepo usecases.InvoiceRepository, priceRepo usecases.PriceRepository, monobankStrategy usecases.AcquiringOperations) *usecases.PaymentUsecase {
	return usecases.NewPaymentUsecase(invoiceRepo, priceRepo, monobankStrategy)
}

func ProvideMonobankStrategy(acquiring *monobank.Acquiring, cfg *Config) *usecases.MonobankAcquiring {
	return usecases.NewMonobankAcquiring(acquiring, cfg.MonobankRedirectUrl, cfg.MonobankWebHookUrl, monobank.USD)
}

func ProvideMonobankAcquiring(cfg *Config) *monobank.Acquiring {
	return monobank.NewAcquiring(cfg.MonobankToken)
}

func ProvideFirestoreClient(cfg *Config) *firestore.Client {
	ctx := context.Background()
	client, err := firestore.NewClient(ctx, cfg.GcpProjectId)
	if err != nil {
		panic(err)
	}
	return client
}

func ProvideFirebaseAuthClient(cfg *Config) *auth.Client {
	client, err := gcpclient.NewFirebaseClient(cfg.GcpProjectId, cfg.GoogleApplicationCredentials)
	if err != nil {
		panic(err)
	}
	return client
}

func ProvideFirebaseAuthMiddleware(client *auth.Client) *HTTPServer.FirebaseAuth {
	return HTTPServer.NewFirebaseAuth(client)
}
