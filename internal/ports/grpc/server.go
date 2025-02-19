package GRPCServer

import (
	"github.com/aerosystems/common-service/gen/protobuf/subscription"
	"github.com/aerosystems/common-service/presenters/grpcserver"
	"github.com/sirupsen/logrus"
)

type Server struct {
	grpcServer *grpcserver.Server
}

func NewGRPCServer(
	cfg *grpcserver.Config,
	log *logrus.Logger,
	subscriptionService *SubscriptionService,
) *Server {
	server := grpcserver.NewGRPCServer(cfg, log)

	server.RegisterService(subscription.SubscriptionService_ServiceDesc, subscriptionService)

	return &Server{
		grpcServer: server,
	}
}

func (s *Server) Run() error {
	return s.grpcServer.Run()
}
