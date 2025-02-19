//go:build wireinject
// +build wireinject

package main

import (
	"context"

	"cloud.google.com/go/firestore"
	"firebase.google.com/go/v4/auth"
	"github.com/google/wire"
	"github.com/sirupsen/logrus"

	"github.com/aerosystems/subscription-service/internal/adapters"
	GRPCServer "github.com/aerosystems/subscription-service/internal/ports/grpc"
	HTTPServer "github.com/aerosystems/subscription-service/internal/ports/http"
	"github.com/aerosystems/subscription-service/internal/usecases"
	"github.com/aerosystems/subscription-service/pkg/monobank"

	"github.com/aerosystems/common-service/clients/gcpclient"
	"github.com/aerosystems/common-service/logger"
	"github.com/aerosystems/common-service/presenters/grpcserver"
	"github.com/aerosystems/common-service/presenters/httpserver"
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
		wire.Bind(new(usecases.AcquiringOperations), new(*usecases.MonobankAcquiring)),
		ProvideApp,
		ProvideLogger,
		ProvideConfig,
		ProvideHTTPServer,
		ProvideGRPCServer,
		ProvideLogrusLogger,
		ProvideHandler,
		ProvidePaymentUsecase,
		ProvideMonobankStrategy,
		ProvideMonobankAcquiring,
		ProvideSubscriptionUsecase,
		ProvidePriceRepo,
		ProvideFirestoreClient,
		ProvideSubscriptionRepo,
		ProvideFirebaseAuthMiddleware,
		ProvideFirebaseAuthClient,
		ProvideInvoiceRepo,
		ProvideSubscriptionService,
	),
	)
}

func ProvideApp(log *logrus.Logger, cfg *Config, httpServer *HTTPServer.Server, gpcServer *GRPCServer.Server) *App {
	panic(wire.Build(NewApp))
}

func ProvideLogger() *logger.Logger {
	panic(wire.Build(logger.NewLogger))
}

func ProvideConfig() *Config {
	panic(wire.Build(NewConfig))
}

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

func ProvideHandler(paymentUsecase HTTPServer.PaymentUsecase, subscriptionUsecase HTTPServer.SubscriptionUsecase) *HTTPServer.Handler {
	panic(wire.Build(HTTPServer.NewHandler))
}

func ProvideSubscriptionService(subscriptionUsecase GRPCServer.SubscriptionUsecase) *GRPCServer.SubscriptionService {
	panic(wire.Build(GRPCServer.NewSubscriptionService))
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

func ProvideSubscriptionUsecase(subscriptionRepo usecases.SubscriptionRepository) *usecases.SubscriptionUsecase {
	panic(wire.Build(usecases.NewSubscriptionUsecase))
}

func ProvidePriceRepo() *adapters.PriceRepo {
	panic(wire.Build(adapters.NewPriceRepo))
}

func ProvideFirestoreClient(cfg *Config) *firestore.Client {
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
