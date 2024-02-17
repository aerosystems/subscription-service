//go:build wireinject
// +build wireinject

package main

import (
	"github.com/aerosystems/subs-service/internal/config"
	"github.com/aerosystems/subs-service/internal/http"
	RPCServer "github.com/aerosystems/subs-service/internal/infrastructure/rpc"
	"github.com/aerosystems/subs-service/internal/models"
	"github.com/aerosystems/subs-service/internal/repository"
	"github.com/aerosystems/subs-service/internal/repository/pg"
	"github.com/aerosystems/subs-service/internal/rest"
	"github.com/aerosystems/subs-service/internal/usecases"
	"github.com/aerosystems/subs-service/pkg/gorm_postgres"
	"github.com/aerosystems/subs-service/pkg/logger"
	"github.com/google/wire"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

//go:generate wire
func InitApp() {
	panic(wire.Build(
		wire.Bind(new(http.TokenService), new(*rest.TokenService)),
		wire.Bind(new(RPCServer.SubscriptionUsecase), new(*usecases.SubscriptionUsecase)),
		wire.Bind(new(usecases.SubscriptionRepository), new(*pg.SubscriptionRepo)),
		wire.Bind(new(usecases.InvoiceRepository), new(*pg.InvoiceRepo)),
		wire.Bind(new(usecases.PaymentRepository), new(*pg.PriceRepo)),
		ProvideApp,
		ProvideLogger,
		ProvideConfig,
		ProvideHTTPServer,
		ProvideRPCServer,
		ProvideLogrusLogger,
		ProvideLogrusEntry,
		ProvideGormPostgres,
		ProvideBaseHandler,
		ProvidePaymentHandler,
		ProvideSubscriptionHandler,
		ProvidePaymentUsecase,
		ProvideSubscriptionUsecase,
		ProvideSubscriptionRepo,
		ProvideInvoiceRepo,
		ProvidePriceRepo,
	),
	)
}

func ProvideApp(log *logrus.Logger, cfg *config.Config, httpServer *HTTPServer.Server, rpcServer RPCServer.Server) *App {
	panic(wire.Build(NewApp))
}

func ProvideLogger() *logger.Logger {
	panic(wire.Build(logger.NewLogger))
}

func ProvideConfig() *config.Config {
	panic(wire.Build(config.NewConfig))
}

func ProvideHTTPServer(log *logrus.Logger, userHandler *rest.UserHandler, tokenHandler *rest.TokenHandler, tokenService HTTPServer.TokenService) *HTTPServer.Server {
	panic(wire.Build(HTTPServer.NewServer))
}

func ProvideRPCServer(log *logrus.Logger, subscriptionUsecase RPCServer.SubscriptionUsecase) *RPCServer.Server {
	panic(wire.Build(RPCServer.NewSubsServer))
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

func ProvidePaymentUsecase(invoiceRepo usecases.InvoiceRepository, priceRepo usecases.PaymentRepository, subscriptionRepo usecases.SubscriptionRepository) *usecases.PaymentUsecase {
	panic(wire.Build(usecases.NewPaymentUsecase))
}

func ProvideSubscriptionUsecase(subscriptionRepo usecases.SubscriptionRepository, priceRepo usecases.PaymentRepository) *usecases.SubscriptionUsecase {
	panic(wire.Build(usecases.NewSubsServiceImpl))
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
