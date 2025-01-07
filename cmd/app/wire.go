//go:build wireinject
// +build wireinject

package main

import (
	"cloud.google.com/go/firestore"
	"context"
	"firebase.google.com/go/auth"
	"github.com/aerosystems/subscription-service/internal/adapters"
	"github.com/aerosystems/subscription-service/internal/common/config"
	CustomErrors "github.com/aerosystems/subscription-service/internal/common/custom_errors"
	"github.com/aerosystems/subscription-service/internal/models"
	GRPCServer "github.com/aerosystems/subscription-service/internal/presenters/grpc"
	HttpServer "github.com/aerosystems/subscription-service/internal/presenters/http"
	"github.com/aerosystems/subscription-service/internal/presenters/http/handlers"
	"github.com/aerosystems/subscription-service/internal/presenters/http/handlers/payment"
	"github.com/aerosystems/subscription-service/internal/presenters/http/handlers/subscription"
	"github.com/aerosystems/subscription-service/internal/presenters/http/middleware"
	"github.com/aerosystems/subscription-service/internal/usecases"
	"github.com/aerosystems/subscription-service/pkg/gcp"
	"github.com/aerosystems/subscription-service/pkg/logger"
	"github.com/aerosystems/subscription-service/pkg/monobank"
	"github.com/google/wire"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

//go:generate wire
func InitApp() *App {
	panic(wire.Build(
		wire.Bind(new(GRPCServer.SubscriptionUsecase), new(*usecases.SubscriptionUsecase)),
		wire.Bind(new(handlers.PaymentUsecase), new(*usecases.PaymentUsecase)),
		wire.Bind(new(handlers.SubscriptionUsecase), new(*usecases.SubscriptionUsecase)),
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
		ProvideXAPiKeyMiddleware,
		ProvideCustomErrorHandler,
	),
	)
}

func ProvideApp(log *logrus.Logger, cfg *config.Config, httpServer *HttpServer.Server, gpcServer *GRPCServer.Server) *App {
	panic(wire.Build(NewApp))
}

func ProvideLogger() *logger.Logger {
	panic(wire.Build(logger.NewLogger))
}

func ProvideConfig() *config.Config {
	panic(wire.Build(config.NewConfig))
}

func ProvideHttpServer(cfg *config.Config, log *logrus.Logger, errorHandler *echo.HTTPErrorHandler, firebaseAuthMiddleware *middleware.FirebaseAuth, xApiKeyAuthMiddleware *middleware.ServiceApiKeyAuth, subscriptionHandler *subscription.Handler, paymentHandler *payment.Handler) *HttpServer.Server {
	return HttpServer.NewServer(cfg.Port, log, errorHandler, firebaseAuthMiddleware, xApiKeyAuthMiddleware, subscriptionHandler, paymentHandler)
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

func ProvideBaseHandler(log *logrus.Logger, cfg *config.Config) *handlers.BaseHandler {
	return handlers.NewBaseHandler(log, cfg.Mode)
}

func ProvidePaymentHandler(baseHandler *handlers.BaseHandler, paymentUsecase handlers.PaymentUsecase) *payment.Handler {
	panic(wire.Build(payment.NewPaymentHandler))
}

func ProvideSubscriptionHandler(baseHandler *handlers.BaseHandler, subscriptionUsecase handlers.SubscriptionUsecase) *subscription.Handler {
	panic(wire.Build(subscription.NewSubscriptionHandler))
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

func ProvideFirebaseAuthMiddleware(client *auth.Client) *middleware.FirebaseAuth {
	return middleware.NewFirebaseAuth(client)
}

func ProvideXAPiKeyMiddleware(cfg *config.Config) *middleware.ServiceApiKeyAuth {
	xApiKeyAuthMiddleware, err := middleware.NewServiceApiKeyAuth(cfg.ApiKey)
	if err != nil {
		panic(err)
	}
	return xApiKeyAuthMiddleware
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
