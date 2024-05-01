package HttpServer

import (
	"github.com/aerosystems/subs-service/internal/models"
)

func (s *Server) setupRoutes() {
	s.echo.GET("/v1/prices", s.subscriptionHandler.GetPrices)

	s.echo.GET("/v1/subscriptions", s.subscriptionHandler.GetSubscriptions, s.firebaseAuthMiddleware.EchoRoleBased(models.CustomerRole))
	s.echo.POST("/v1/subscriptions", s.subscriptionHandler.CreateSubscription, s.firebaseAuthMiddleware.EchoRoleBased(models.CustomerRole))
	s.echo.PATCH("/v1/subscriptions/:id", s.subscriptionHandler.UpdateSubscription, s.firebaseAuthMiddleware.EchoRoleBased(models.StaffRole))
	s.echo.DELETE("/v1/subscriptions/:id", s.subscriptionHandler.DeleteSubscription, s.firebaseAuthMiddleware.EchoRoleBased(models.StaffRole))

	s.echo.POST("/v1/invoices/:payment_method", s.paymentHandler.CreateInvoice, s.firebaseAuthMiddleware.EchoRoleBased(models.CustomerRole))

	s.echo.POST("/v1/webhook/:payment_method", s.paymentHandler.WebhookPayment)
}
