package HttpServer

import (
	"github.com/aerosystems/subs-service/internal/models"
)

func (s *Server) setupRoutes() {
	s.echo.GET("/v1/prices", s.subscriptionHandler.GetPrices)

	s.echo.GET("/v1/subscriptions", s.subscriptionHandler.GetSubscriptions, s.AuthTokenMiddleware(models.CustomerRole))
	s.echo.POST("/v1/subscriptions", s.subscriptionHandler.CreateSubscription, s.AuthTokenMiddleware(models.CustomerRole))
	s.echo.PATCH("/v1/subscriptions/:id", s.subscriptionHandler.UpdateSubscription, s.AuthTokenMiddleware(models.StaffRole))
	s.echo.DELETE("/v1/subscriptions/:id", s.subscriptionHandler.DeleteSubscription, s.AuthTokenMiddleware(models.StaffRole))

	s.echo.POST("/v1/invoices/:payment_method", s.paymentHandler.CreateInvoice, s.AuthTokenMiddleware(models.CustomerRole))

	s.echo.POST("/v1/webhook/:payment_method", s.paymentHandler.WebhookPayment)
}
