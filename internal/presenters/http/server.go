package HttpServer

import (
	"fmt"
	"github.com/aerosystems/subs-service/internal/presenters/http/handlers"
	"github.com/aerosystems/subs-service/internal/presenters/http/middleware"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

const webPort = 80

type Server struct {
	log                    *logrus.Logger
	echo                   *echo.Echo
	firebaseAuthMiddleware *middleware.FirebaseAuth
	subscriptionHandler    *handlers.SubscriptionHandler
	paymentHandler         *handlers.PaymentHandler
}

func NewServer(
	log *logrus.Logger,
	firebaseAuthMiddleware *middleware.FirebaseAuth,
	subscriptionHandler *handlers.SubscriptionHandler,
	paymentHandler *handlers.PaymentHandler,

) *Server {
	return &Server{
		log:                    log,
		echo:                   echo.New(),
		firebaseAuthMiddleware: firebaseAuthMiddleware,
		subscriptionHandler:    subscriptionHandler,
		paymentHandler:         paymentHandler,
	}
}

func (s *Server) Run() error {
	s.setupMiddleware()
	s.setupRoutes()
	s.setupValidator()
	s.log.Infof("starting HTTP server subs-service on port %d\n", webPort)
	return s.echo.Start(fmt.Sprintf(":%d", webPort))
}
