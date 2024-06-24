package HttpServer

import (
	"github.com/aerosystems/subscription-service/internal/models"
)

func (s *Server) setupRoutes() {
	s.echo.GET("/v1/prices", s.paymentHandler.GetPrices)

	s.echo.GET("/v1/subscriptions", s.subscriptionHandler.GetSubscriptions, s.firebaseAuthMiddleware.RoleBasedAuth(models.CustomerRole))
	s.echo.POST("/v1/subscriptions", s.subscriptionHandler.CreateSubscription, s.firebaseAuthMiddleware.RoleBasedAuth(models.CustomerRole))
	s.echo.PATCH("/v1/subscriptions/:id", s.subscriptionHandler.UpdateSubscription, s.firebaseAuthMiddleware.RoleBasedAuth(models.StaffRole))
	s.echo.DELETE("/v1/subscriptions/:id", s.subscriptionHandler.DeleteSubscription, s.firebaseAuthMiddleware.RoleBasedAuth(models.StaffRole))

	s.echo.POST("/v1/subscriptions/create-free-trial", s.subscriptionHandler.CreateFreeTrial)

	s.echo.POST("/v1/invoices/:payment_method", s.paymentHandler.CreateInvoice, s.firebaseAuthMiddleware.RoleBasedAuth(models.CustomerRole))

	s.echo.POST("/v1/webhook/:payment_method", s.paymentHandler.WebhookPayment)
}
