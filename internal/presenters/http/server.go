package HTTPServer

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

type Server struct {
	port                   int
	log                    *logrus.Logger
	echo                   *echo.Echo
	firebaseAuthMiddleware *FirebaseAuth
	subscriptionHandler    *SubscriptionHandler
	paymentHandler         *PaymentHandler
}

func NewServer(
	port int,
	log *logrus.Logger,
	errorHandler *echo.HTTPErrorHandler,
	firebaseAuthMiddleware *FirebaseAuth,
	subscriptionHandler *SubscriptionHandler,
	paymentHandler *PaymentHandler,
) *Server {
	server := &Server{
		port:                   port,
		log:                    log,
		echo:                   echo.New(),
		firebaseAuthMiddleware: firebaseAuthMiddleware,
		subscriptionHandler:    subscriptionHandler,
		paymentHandler:         paymentHandler,
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
