//go:build wireinject
// +build wireinject

package main

import (
	"context"

	"cloud.google.com/go/firestore"
	"firebase.google.com/go/v4/auth"
	"github.com/google/wire"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"

	"github.com/aerosystems/subscription-service/internal/adapters"
	"github.com/aerosystems/subscription-service/internal/common/config"
	CustomErrors "github.com/aerosystems/subscription-service/internal/common/custom_errors"
	"github.com/aerosystems/subscription-service/internal/models"
	GRPCServer "github.com/aerosystems/subscription-service/internal/presenters/grpc"
	HTTPServer "github.com/aerosystems/subscription-service/internal/presenters/http"
	"github.com/aerosystems/subscription-service/internal/usecases"
	"github.com/aerosystems/subscription-service/pkg/gcp"
	"github.com/aerosystems/subscription-service/pkg/logger"
	"github.com/aerosystems/subscription-service/pkg/monobank"
)

//go:generate wire
func InitApp() *App {
	panic(wire.Build(
		wire.Bind(new(GRPCServer.SubscriptionUsecase), new(*usecases.SubscriptionUsecase)),
		wire.Bind(new(HTTPServer.PaymentUsecase), new(*usecases.PaymentUsecase)),
		wire.Bind(new(HTTPServer.SubscriptionUsecase), new(*usecases.SubscriptionUsecase)),
		wire.Bind(new(usecases.SubscriptionRepository), new(*adapters.SubscriptionRepo)),
		wire.Bind(new(usecases.InvoiceRepository), new(*adapters.InvoiceRepo)),
		wire.Bind(new(usecases.PriceRepository), new(*adapters.PriceRepo)),
		ProvideApp,
		ProvideLogger,
		ProvideConfig,
		ProvideHttpServer,
		ProvideGRPCServer,
		ProvideGRPCSubscriptionHandler,
		ProvideLogrusLogger,
		ProvideBaseHandler,
		ProvidePaymentHandler,
		ProvideSubscriptionHandler,
		ProvidePaymentUsecase,
		ProvidePaymentMap,
		ProvideMonobankStrategy,
		ProvideMonobankAcquiring,
		ProvideSubscriptionUsecase,
		ProvidePriceRepo,
		ProvideFirestoreClient,
		ProvideSubscriptionRepo,
		ProvideFirebaseAuthMiddleware,
		ProvideFirebaseAuthClient,
		ProvideInvoiceRepo,
		ProvideCustomErrorHandler,
	),
	)
}

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

func ProvideLogger() *logger.Logger {
	panic(wire.Build(logger.NewLogger))
}

func ProvideConfig() *config.Config {
	panic(wire.Build(config.NewConfig))
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

func ProvideGRPCSubscriptionHandler(subscriptionUsecase GRPCServer.SubscriptionUsecase) *GRPCServer.SubscriptionHandler {
	panic(wire.Build(GRPCServer.NewSubscriptionHandler))
}

func ProvideLogrusLogger(log *logger.Logger) *logrus.Logger {
	return log.Logger
}

func ProvideBaseHandler(log *logrus.Logger, cfg *config.Config) *HTTPServer.BaseHandler {
	return HTTPServer.NewBaseHandler(log, cfg.Mode)
}

func ProvidePaymentHandler(baseHandler *HTTPServer.BaseHandler, paymentUsecase HTTPServer.PaymentUsecase) *HTTPServer.PaymentHandler {
	panic(wire.Build(HTTPServer.NewPaymentHandler))
}

func ProvideSubscriptionHandler(baseHandler *HTTPServer.BaseHandler, subscriptionUsecase HTTPServer.SubscriptionUsecase) *HTTPServer.SubscriptionHandler {
	panic(wire.Build(HTTPServer.NewSubscriptionHandler))
}

func ProvidePaymentUsecase(invoiceRepo usecases.InvoiceRepository, priceRepo usecases.PriceRepository, strategies map[models.PaymentMethod]usecases.AcquiringOperations) *usecases.PaymentUsecase {
	panic(wire.Build(usecases.NewPaymentUsecase))
}

func ProvidePaymentMap(monobankStrategy *usecases.MonobankStrategy) map[models.PaymentMethod]usecases.AcquiringOperations {
	return map[models.PaymentMethod]usecases.AcquiringOperations{
		models.MonobankPaymentMethod: monobankStrategy,
	}
}

func ProvideMonobankStrategy(acquiring *monobank.Acquiring, cfg *config.Config) *usecases.MonobankStrategy {
	return usecases.NewMonobankStrategy(acquiring, cfg.MonobankRedirectUrl, cfg.MonobankWebHookUrl, monobank.USD)
}

func ProvideMonobankAcquiring(cfg *config.Config) *monobank.Acquiring {
	return monobank.NewAcquiring(cfg.MonobankToken)
}

func ProvideSubscriptionUsecase(subscriptionRepo usecases.SubscriptionRepository) *usecases.SubscriptionUsecase {
	panic(wire.Build(usecases.NewSubscriptionUsecase))
}

func ProvidePriceRepo() *adapters.PriceRepo {
	panic(wire.Build(adapters.NewPriceRepo))
}

func ProvideFirestoreClient(cfg *config.Config) *firestore.Client {
	ctx := context.Background()
	client, err := firestore.NewClient(ctx, cfg.GcpProjectId)
	if err != nil {
		panic(err)
	}
	return client
}

func ProvideSubscriptionRepo(client *firestore.Client) *adapters.SubscriptionRepo {
	panic(wire.Build(adapters.NewSubscriptionRepo))
}

func ProvideInvoiceRepo(client *firestore.Client) *adapters.InvoiceRepo {
	panic(wire.Build(adapters.NewInvoiceRepo))
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
