package HttpServer

import (
	"github.com/aerosystems/subscription-service/internal/models"
)

func (s *Server) setupRoutes() {
	s.echo.GET("/v1/prices", s.subscriptionHandler.GetPrices)

	s.echo.GET("/v1/subscriptions", s.subscriptionHandler.GetSubscriptions, s.firebaseAuthMiddleware.RoleBased(models.CustomerRole))
	s.echo.POST("/v1/subscriptions", s.subscriptionHandler.CreateSubscription, s.firebaseAuthMiddleware.RoleBased(models.CustomerRole))
	s.echo.PATCH("/v1/subscriptions/:id", s.subscriptionHandler.UpdateSubscription, s.firebaseAuthMiddleware.RoleBased(models.StaffRole))
	s.echo.DELETE("/v1/subscriptions/:id", s.subscriptionHandler.DeleteSubscription, s.firebaseAuthMiddleware.RoleBased(models.StaffRole))

	s.echo.POST("/v1/invoices/:payment_method", s.paymentHandler.CreateInvoice, s.firebaseAuthMiddleware.RoleBased(models.CustomerRole))

	s.echo.POST("/v1/webhook/:payment_method", s.paymentHandler.WebhookPayment)
}
