package HTTPServer

import (
	"fmt"
	"github.com/aerosystems/subs-service/internal/infrastructure/rest"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

const webPort = 80

type Server struct {
	log                 *logrus.Logger
	echo                *echo.Echo
	subscriptionHandler *rest.SubscriptionHandler
	paymentHandler      *rest.PaymentHandler
	tokenService        TokenService
}

func NewServer(
	log *logrus.Logger,
	subscriptionHandler *rest.SubscriptionHandler,
	paymentHandler *rest.PaymentHandler,
	tokenService TokenService,

) *Server {
	return &Server{
		log:                 log,
		echo:                echo.New(),
		subscriptionHandler: subscriptionHandler,
		paymentHandler:      paymentHandler,
		tokenService:        tokenService,
	}
}

func (s *Server) Run() error {
	s.setupMiddleware()
	s.setupRoutes()
	s.setupValidator()
	s.log.Infof("starting HTTP server subs-service on port %d\n", webPort)
	return s.echo.Start(fmt.Sprintf(":%d", webPort))
}
