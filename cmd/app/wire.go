//go:build wireinject
// +build wireinject

package main

import (
	"cloud.google.com/go/firestore"
	"context"
	"firebase.google.com/go/auth"
	"github.com/aerosystems/subscription-service/internal/config"
	"github.com/aerosystems/subscription-service/internal/infrastructure/adapters/broker"
	"github.com/aerosystems/subscription-service/internal/infrastructure/repository/fire"
	"github.com/aerosystems/subscription-service/internal/infrastructure/repository/memory"
	"github.com/aerosystems/subscription-service/internal/models"
	HttpServer "github.com/aerosystems/subscription-service/internal/presenters/http"
	"github.com/aerosystems/subscription-service/internal/presenters/http/handlers"
	"github.com/aerosystems/subscription-service/internal/presenters/http/handlers/payment"
	"github.com/aerosystems/subscription-service/internal/presenters/http/handlers/subscription"
	"github.com/aerosystems/subscription-service/internal/presenters/http/middleware"
	RpcServer "github.com/aerosystems/subscription-service/internal/presenters/rpc"
	"github.com/aerosystems/subscription-service/internal/usecases"
	"github.com/aerosystems/subscription-service/pkg/firebase"
	"github.com/aerosystems/subscription-service/pkg/logger"
	"github.com/aerosystems/subscription-service/pkg/monobank"
	"github.com/aerosystems/subscription-service/pkg/pubsub"
	"github.com/google/wire"
	"github.com/sirupsen/logrus"
)

//go:generate wire
func InitApp() *App {
	panic(wire.Build(
		wire.Bind(new(RpcServer.SubscriptionUsecase), new(*usecases.SubscriptionUsecase)),
		wire.Bind(new(handlers.PaymentUsecase), new(*usecases.PaymentUsecase)),
		wire.Bind(new(handlers.SubscriptionUsecase), new(*usecases.SubscriptionUsecase)),
		wire.Bind(new(usecases.SubscriptionRepository), new(*fire.SubscriptionRepo)),
		wire.Bind(new(usecases.InvoiceRepository), new(*fire.InvoiceRepo)),
		wire.Bind(new(usecases.PriceRepository), new(*memory.PriceRepo)),
		wire.Bind(new(usecases.ProjectAdapter), new(*broker.ProjectEventsAdapter)),
		ProvideApp,
		ProvideLogger,
		ProvideConfig,
		ProvideHttpServer,
		ProvideRpcServer,
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
		ProvidePubSubClient,
		ProvideProjectEventsAdapter,
	),
	)
}

func ProvideApp(log *logrus.Logger, cfg *config.Config, httpServer *HttpServer.Server, rpcServer *RpcServer.Server) *App {
	panic(wire.Build(NewApp))
}

func ProvideLogger() *logger.Logger {
	panic(wire.Build(logger.NewLogger))
}

func ProvideConfig() *config.Config {
	panic(wire.Build(config.NewConfig))
}

func ProvideHttpServer(log *logrus.Logger, firebaseAuthMiddleware *middleware.FirebaseAuth, xApiKeyAuthMiddleware *middleware.ServiceApiKeyAuth, subscriptionHandler *subscription.Handler, paymentHandler *payment.Handler) *HttpServer.Server {
	return HttpServer.NewServer(log, firebaseAuthMiddleware, xApiKeyAuthMiddleware, subscriptionHandler, paymentHandler)
}

func ProvideRpcServer(log *logrus.Logger, subscriptionUsecase RpcServer.SubscriptionUsecase) *RpcServer.Server {
	panic(wire.Build(RpcServer.NewServer))
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

func ProvideSubscriptionUsecase(subscriptionRepo usecases.SubscriptionRepository, projectAdapter usecases.ProjectAdapter) *usecases.SubscriptionUsecase {
	panic(wire.Build(usecases.NewSubscriptionUsecase))
}

func ProvidePriceRepo() *memory.PriceRepo {
	panic(wire.Build(memory.NewPriceRepo))
}

func ProvideFirestoreClient(cfg *config.Config) *firestore.Client {
	ctx := context.Background()
	client, err := firestore.NewClient(ctx, cfg.GcpProjectId)
	if err != nil {
		panic(err)
	}
	return client
}

func ProvideSubscriptionRepo(client *firestore.Client) *fire.SubscriptionRepo {
	panic(wire.Build(fire.NewSubscriptionRepo))
}

func ProvideInvoiceRepo(client *firestore.Client) *fire.InvoiceRepo {
	panic(wire.Build(fire.NewInvoiceRepo))
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
	app, err := firebaseApp.NewApp(cfg.GcpProjectId, cfg.GoogleApplicationCredentials)
	if err != nil {
		panic(err)
	}
	return app.Client
}

func ProvidePubSubClient(cfg *config.Config) *PubSub.Client {
	client, err := PubSub.NewClientWithAuth(cfg.GoogleApplicationCredentials)
	if err != nil {
		panic(err)
	}
	return client
}

func ProvideProjectEventsAdapter(pubSubClient *PubSub.Client, cfg *config.Config) *broker.ProjectEventsAdapter {
	return broker.NewProjectEventsAdapter(pubSubClient, cfg.ProjectTopicId, cfg.ProjectSubName, cfg.ProjectCreateEndpoint, cfg.ProjectServiceApiKey)
}
