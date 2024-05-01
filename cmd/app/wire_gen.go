// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"cloud.google.com/go/firestore"
	"context"
	"firebase.google.com/go/auth"
	"github.com/aerosystems/subs-service/internal/config"
	"github.com/aerosystems/subs-service/internal/infrastructure/http"
	"github.com/aerosystems/subs-service/internal/infrastructure/http/handlers"
	"github.com/aerosystems/subs-service/internal/infrastructure/http/middleware"
	"github.com/aerosystems/subs-service/internal/infrastructure/rpc"
	"github.com/aerosystems/subs-service/internal/models"
	"github.com/aerosystems/subs-service/internal/repository"
	"github.com/aerosystems/subs-service/internal/repository/fire"
	"github.com/aerosystems/subs-service/internal/repository/pg"
	"github.com/aerosystems/subs-service/internal/usecases"
	"github.com/aerosystems/subs-service/pkg/firebase"
	"github.com/aerosystems/subs-service/pkg/gorm_postgres"
	"github.com/aerosystems/subs-service/pkg/logger"
	"github.com/aerosystems/subs-service/pkg/monobank"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

// Injectors from wire.go:

//go:generate wire
func InitApp() *App {
	logger := ProvideLogger()
	logrusLogger := ProvideLogrusLogger(logger)
	config := ProvideConfig()
	client := ProvideFirebaseAuthClient(config)
	firebaseAuth := ProvideFirebaseAuthMiddleware(client)
	baseHandler := ProvideBaseHandler(logrusLogger, config)
	firestoreClient := ProvideFirestoreClient(config)
	subscriptionRepo := ProvideFireSubscriptionRepo(firestoreClient)
	priceRepo := ProvidePriceRepo()
	subscriptionUsecase := ProvideSubscriptionUsecase(subscriptionRepo, priceRepo)
	subscriptionHandler := ProvideSubscriptionHandler(baseHandler, subscriptionUsecase)
	entry := ProvideLogrusEntry(logger)
	db := ProvideGormPostgres(entry, config)
	invoiceRepo := ProvideInvoiceRepo(db)
	acquiring := ProvideMonobankAcquiring(config)
	monobankStrategy := ProvideMonobankStrategy(acquiring, config)
	v := ProvidePaymentMap(monobankStrategy)
	paymentUsecase := ProvidePaymentUsecase(invoiceRepo, priceRepo, v)
	paymentHandler := ProvidePaymentHandler(baseHandler, paymentUsecase)
	server := ProvideHttpServer(logrusLogger, firebaseAuth, subscriptionHandler, paymentHandler)
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

func ProvidePaymentHandler(baseHandler *handlers.BaseHandler, paymentUsecase handlers.PaymentUsecase) *handlers.PaymentHandler {
	paymentHandler := handlers.NewPaymentHandler(baseHandler, paymentUsecase)
	return paymentHandler
}

func ProvideSubscriptionHandler(baseHandler *handlers.BaseHandler, subscriptionUsecase handlers.SubscriptionUsecase) *handlers.SubscriptionHandler {
	subscriptionHandler := handlers.NewSubscriptionHandler(baseHandler, subscriptionUsecase)
	return subscriptionHandler
}

func ProvidePaymentUsecase(invoiceRepo usecases.InvoiceRepository, priceRepo usecases.PriceRepository, strategies map[models.PaymentMethod]usecases.AcquiringOperations) *usecases.PaymentUsecase {
	paymentUsecase := usecases.NewPaymentUsecase(invoiceRepo, priceRepo, strategies)
	return paymentUsecase
}

func ProvideSubscriptionUsecase(subscriptionRepo usecases.SubscriptionRepository, priceRepo usecases.PriceRepository) *usecases.SubscriptionUsecase {
	subscriptionUsecase := usecases.NewSubscriptionUsecase(subscriptionRepo, priceRepo)
	return subscriptionUsecase
}

func ProvideSubscriptionRepo(client *gorm.DB) *pg.SubscriptionRepo {
	subscriptionRepo := pg.NewSubscriptionRepo(client)
	return subscriptionRepo
}

func ProvideInvoiceRepo(client *gorm.DB) *pg.InvoiceRepo {
	invoiceRepo := pg.NewInvoiceRepo(client)
	return invoiceRepo
}

func ProvidePriceRepo() *repository.PriceRepo {
	priceRepo := repository.NewPriceRepo()
	return priceRepo
}

func ProvideFireSubscriptionRepo(client *firestore.Client) *fire.SubscriptionRepo {
	subscriptionRepo := fire.NewSubscriptionRepo(client)
	return subscriptionRepo
}

// wire.go:

func ProvideHttpServer(log *logrus.Logger, firebaseAuthMiddleware *middleware.FirebaseAuth, subscriptionHandler *handlers.SubscriptionHandler, paymentHandler *handlers.PaymentHandler) *HttpServer.Server {
	return HttpServer.NewServer(log, firebaseAuthMiddleware, subscriptionHandler, paymentHandler)
}

func ProvideLogrusEntry(log *logger.Logger) *logrus.Entry {
	return logrus.NewEntry(log.Logger)
}

func ProvideLogrusLogger(log *logger.Logger) *logrus.Logger {
	return log.Logger
}

func ProvideGormPostgres(e *logrus.Entry, cfg *config.Config) *gorm.DB {
	db := GormPostgres.NewClient(e, cfg.PostgresDSN)
	if err := db.AutoMigrate(pg.Subscription{}, pg.Invoice{}); err != nil {
		panic(err)
	}
	return db
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

func ProvideFirebaseAuthClient(cfg *config.Config) *auth.Client {
	app, err := firebaseApp.NewApp(cfg.GcpProjectId, cfg.GcpServiceAccountFilePath)
	if err != nil {
		panic(err)
	}
	return app.Client
}
