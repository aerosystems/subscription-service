package HttpServer

import (
	"fmt"
	"github.com/aerosystems/subscription-service/internal/presenters/http/handlers/payment"
	"github.com/aerosystems/subscription-service/internal/presenters/http/handlers/subscription"
	"github.com/aerosystems/subscription-service/internal/presenters/http/middleware"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

type Server struct {
	port                        int
	log                         *logrus.Logger
	echo                        *echo.Echo
	firebaseAuthMiddleware      *middleware.FirebaseAuth
	serviceApiKeyAuthMiddleware *middleware.ServiceApiKeyAuth
	subscriptionHandler         *subscription.Handler
	paymentHandler              *payment.Handler
}

func NewServer(
	port int,
	log *logrus.Logger,
	errorHandler *echo.HTTPErrorHandler,
	firebaseAuthMiddleware *middleware.FirebaseAuth,
	serviceApiKeyAuthMiddleware *middleware.ServiceApiKeyAuth,
	subscriptionHandler *subscription.Handler,
	paymentHandler *payment.Handler,
) *Server {
	server := &Server{
		port:                        port,
		log:                         log,
		echo:                        echo.New(),
		firebaseAuthMiddleware:      firebaseAuthMiddleware,
		serviceApiKeyAuthMiddleware: serviceApiKeyAuthMiddleware,
		subscriptionHandler:         subscriptionHandler,
		paymentHandler:              paymentHandler,
	}
	if errorHandler != nil {
		server.echo.HTTPErrorHandler = *errorHandler
	}
	return server
}

func (s *Server) Run() error {
	s.setupMiddleware()
	s.setupRoutes()
	s.setupValidator()
	return s.echo.Start(fmt.Sprintf(":%d", s.port))
}
