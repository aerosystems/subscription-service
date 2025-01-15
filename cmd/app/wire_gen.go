// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"cloud.google.com/go/firestore"
	"context"
	"firebase.google.com/go/v4/auth"
	"github.com/aerosystems/subscription-service/internal/adapters"
	"github.com/aerosystems/subscription-service/internal/common/config"
	"github.com/aerosystems/subscription-service/internal/common/custom_errors"
	"github.com/aerosystems/subscription-service/internal/models"
	"github.com/aerosystems/subscription-service/internal/presenters/grpc"
	"github.com/aerosystems/subscription-service/internal/presenters/http"
	"github.com/aerosystems/subscription-service/internal/usecases"
	"github.com/aerosystems/subscription-service/pkg/gcp"
	"github.com/aerosystems/subscription-service/pkg/logger"
	"github.com/aerosystems/subscription-service/pkg/monobank"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

// Injectors from wire.go:

//go:generate wire
func InitApp() *App {
	logger := ProvideLogger()
	logrusLogger := ProvideLogrusLogger(logger)
	config := ProvideConfig()
	httpErrorHandler := ProvideCustomErrorHandler(config)
	client := ProvideFirebaseAuthClient(config)
	firebaseAuth := ProvideFirebaseAuthMiddleware(client)
	baseHandler := ProvideBaseHandler(logrusLogger, config)
	firestoreClient := ProvideFirestoreClient(config)
	subscriptionRepo := ProvideSubscriptionRepo(firestoreClient)
	subscriptionUsecase := ProvideSubscriptionUsecase(subscriptionRepo)
	subscriptionHandler := ProvideSubscriptionHandler(baseHandler, subscriptionUsecase)
	invoiceRepo := ProvideInvoiceRepo(firestoreClient)
	priceRepo := ProvidePriceRepo()
	acquiring := ProvideMonobankAcquiring(config)
	monobankStrategy := ProvideMonobankStrategy(acquiring, config)
	v := ProvidePaymentMap(monobankStrategy)
	paymentUsecase := ProvidePaymentUsecase(invoiceRepo, priceRepo, v)
	paymentHandler := ProvidePaymentHandler(baseHandler, paymentUsecase)
	server := ProvideHttpServer(config, logrusLogger, httpErrorHandler, firebaseAuth, subscriptionHandler, paymentHandler)
	grpcServerSubscriptionHandler := ProvideGRPCSubscriptionHandler(subscriptionUsecase)
	grpcServerServer := ProvideGRPCServer(config, logrusLogger, grpcServerSubscriptionHandler)
	app := ProvideApp(logrusLogger, config, server, grpcServerServer)
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

func ProvideGRPCSubscriptionHandler(subscriptionUsecase GRPCServer.SubscriptionUsecase) *GRPCServer.SubscriptionHandler {
	subscriptionHandler := GRPCServer.NewSubscriptionHandler(subscriptionUsecase)
	return subscriptionHandler
}

func ProvidePaymentHandler(baseHandler *HTTPServer.BaseHandler, paymentUsecase HTTPServer.PaymentUsecase) *HTTPServer.PaymentHandler {
	paymentHandler := HTTPServer.NewPaymentHandler(baseHandler, paymentUsecase)
	return paymentHandler
}

func ProvideSubscriptionHandler(baseHandler *HTTPServer.BaseHandler, subscriptionUsecase HTTPServer.SubscriptionUsecase) *HTTPServer.SubscriptionHandler {
	subscriptionHandler := HTTPServer.NewSubscriptionHandler(baseHandler, subscriptionUsecase)
	return subscriptionHandler
}

func ProvidePaymentUsecase(invoiceRepo usecases.InvoiceRepository, priceRepo usecases.PriceRepository, strategies map[models.PaymentMethod]usecases.AcquiringOperations) *usecases.PaymentUsecase {
	paymentUsecase := usecases.NewPaymentUsecase(invoiceRepo, priceRepo, strategies)
	return paymentUsecase
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

func ProvideApp(log *logrus.Logger, cfg *config.Config, httpServer *HTTPServer.Server, gpcServer *GRPCServer.Server) *App {
	if log == nil {
		panic("log is nil")
	}
	if cfg == nil {
		panic("config is nil")
	}
	if httpServer == nil {
		panic("HTTP server is nil")
	}
	if gpcServer == nil {
		panic("GRPC server is nil")
	}
	return NewApp(log, cfg, httpServer, gpcServer)
}

func ProvideHttpServer(cfg *config.Config, log *logrus.Logger, errorHandler *echo.HTTPErrorHandler, firebaseAuthMiddleware *HTTPServer.FirebaseAuth, subscriptionHandler *HTTPServer.SubscriptionHandler, paymentHandler *HTTPServer.PaymentHandler) *HTTPServer.Server {
	if cfg == nil {
		panic("config is nil")
	}
	if log == nil {
		panic("log is nil")
	}
	if errorHandler == nil {
		panic("error handler is nil")
	}
	if firebaseAuthMiddleware == nil {
		panic("Firebase Auth middleware is nil")
	}
	if subscriptionHandler == nil {
		panic("subscription handler is nil")
	}
	if paymentHandler == nil {
		panic("payment handler is nil")
	}
	return HTTPServer.NewServer(cfg.Port, log, errorHandler, firebaseAuthMiddleware, subscriptionHandler, paymentHandler)
}

func ProvideGRPCServer(cfg *config.Config, log *logrus.Logger, grpcSubscriptionHandler *GRPCServer.SubscriptionHandler) *GRPCServer.Server {
	return GRPCServer.NewGRPCServer(cfg.Port, log, grpcSubscriptionHandler)
}

func ProvideLogrusLogger(log *logger.Logger) *logrus.Logger {
	return log.Logger
}

func ProvideBaseHandler(log *logrus.Logger, cfg *config.Config) *HTTPServer.BaseHandler {
	return HTTPServer.NewBaseHandler(log, cfg.Mode)
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

func ProvideFirebaseAuthMiddleware(client *auth.Client) *HTTPServer.FirebaseAuth {
	return HTTPServer.NewFirebaseAuth(client)
}

func ProvideFirebaseAuthClient(cfg *config.Config) *auth.Client {
	app, err := gcp.NewFirebaseApp(cfg.GcpProjectId, cfg.GoogleApplicationCredentials)
	if err != nil {
		panic(err)
	}
	return app.Client
}

func ProvideCustomErrorHandler(cfg *config.Config) *echo.HTTPErrorHandler {
	errorHandler := CustomErrors.NewEchoErrorHandler(cfg.Mode)
	return &errorHandler
}
