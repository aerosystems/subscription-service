//go:build wireinject
// +build wireinject

package main

import (
	"github.com/aerosystems/subs-service/internal/config"
	HTTPServer "github.com/aerosystems/subs-service/internal/http"
	"github.com/aerosystems/subs-service/internal/infrastructure/rest"
	RPCServer "github.com/aerosystems/subs-service/internal/infrastructure/rpc"
	"github.com/aerosystems/subs-service/internal/models"
	"github.com/aerosystems/subs-service/internal/repository"
	"github.com/aerosystems/subs-service/internal/repository/pg"
	"github.com/aerosystems/subs-service/internal/usecases"
	"github.com/aerosystems/subs-service/pkg/gorm_postgres"
	"github.com/aerosystems/subs-service/pkg/logger"
	"github.com/aerosystems/subs-service/pkg/monobank"
	"github.com/aerosystems/subs-service/pkg/oauth"
	"github.com/google/wire"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

//go:generate wire
func InitApp() *App {
	panic(wire.Build(
		wire.Bind(new(HTTPServer.TokenService), new(*OAuthService.AccessTokenService)),
		wire.Bind(new(RPCServer.SubscriptionUsecase), new(*usecases.SubscriptionUsecase)),
		wire.Bind(new(rest.PaymentUsecase), new(*usecases.PaymentUsecase)),
		wire.Bind(new(rest.SubscriptionUsecase), new(*usecases.SubscriptionUsecase)),
		wire.Bind(new(usecases.SubscriptionRepository), new(*pg.SubscriptionRepo)),
		wire.Bind(new(usecases.InvoiceRepository), new(*pg.InvoiceRepo)),
		wire.Bind(new(usecases.PriceRepository), new(*repository.PriceRepo)),
		ProvideApp,
		ProvideLogger,
		ProvideConfig,
		ProvideHTTPServer,
		ProvideRPCServer,
		ProvideAccessTokenService,
		ProvideLogrusLogger,
		ProvideLogrusEntry,
		ProvideGormPostgres,
		ProvideBaseHandler,
		ProvidePaymentHandler,
		ProvideSubscriptionHandler,
		ProvidePaymentUsecase,
		ProvidePaymentMap,
		ProvideMonobankStrategy,
		ProvideMonobankAcquiring,
		ProvideSubscriptionUsecase,
		ProvideSubscriptionRepo,
		ProvideInvoiceRepo,
		ProvidePriceRepo,
	),
	)
}

func ProvideApp(log *logrus.Logger, cfg *config.Config, httpServer *HTTPServer.Server, rpcServer *RPCServer.Server) *App {
	panic(wire.Build(NewApp))
}

func ProvideLogger() *logger.Logger {
	panic(wire.Build(logger.NewLogger))
}

func ProvideConfig() *config.Config {
	panic(wire.Build(config.NewConfig))
}
func ProvideAccessTokenService(cfg *config.Config) *OAuthService.AccessTokenService {
	return OAuthService.NewAccessTokenService(cfg.AccessSecret)
}

func ProvideHTTPServer(log *logrus.Logger, subscriptionHandler *rest.SubscriptionHandler, paymentHandler *rest.PaymentHandler, tokenService HTTPServer.TokenService) *HTTPServer.Server {
	panic(wire.Build(HTTPServer.NewServer))
}

func ProvideRPCServer(log *logrus.Logger, subscriptionUsecase RPCServer.SubscriptionUsecase) *RPCServer.Server {
	panic(wire.Build(RPCServer.NewServer))
}

func ProvideLogrusEntry(log *logger.Logger) *logrus.Entry {
	return logrus.NewEntry(log.Logger)
}

func ProvideLogrusLogger(log *logger.Logger) *logrus.Logger {
	return log.Logger
}

func ProvideGormPostgres(e *logrus.Entry, cfg *config.Config) *gorm.DB {
	db := GormPostgres.NewClient(e, cfg.PostgresDSN)
	if err := db.AutoMigrate(models.Subscription{}, models.Invoice{}); err != nil { // TODO: Move to migration
		panic(err)
	}
	return db
}

func ProvideBaseHandler(log *logrus.Logger, cfg *config.Config) *rest.BaseHandler {
	return rest.NewBaseHandler(log, cfg.Mode)
}

func ProvidePaymentHandler(baseHandler *rest.BaseHandler, paymentUsecase rest.PaymentUsecase) *rest.PaymentHandler {
	panic(wire.Build(rest.NewPaymentHandler))
}

func ProvideSubscriptionHandler(baseHandler *rest.BaseHandler, subscriptionUsecase rest.SubscriptionUsecase) *rest.SubscriptionHandler {
	panic(wire.Build(rest.NewSubscriptionHandler))
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

func ProvideSubscriptionUsecase(subscriptionRepo usecases.SubscriptionRepository, priceRepo usecases.PriceRepository) *usecases.SubscriptionUsecase {
	panic(wire.Build(usecases.NewSubscriptionUsecase))
}

func ProvideSubscriptionRepo(client *gorm.DB) *pg.SubscriptionRepo {
	panic(wire.Build(pg.NewSubscriptionRepo))
}

func ProvideInvoiceRepo(client *gorm.DB) *pg.InvoiceRepo {
	panic(wire.Build(pg.NewInvoiceRepo))
}

func ProvidePriceRepo() *repository.PriceRepo {
	panic(wire.Build(repository.NewPriceRepo))
}
