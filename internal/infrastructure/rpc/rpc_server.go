package RPCServer

import (
	"fmt"
	"github.com/aerosystems/subs-service/internal/services"
	"github.com/sirupsen/logrus"
	"net"
	"net/rpc"
)

type SubsServer struct {
	rpcPort     int
	log         *logrus.Logger
	subsService services.SubsService
}

func NewSubsServer(
	rpcPort int,
	log *logrus.Logger,
	subsService services.SubsService,
) *SubsServer {
	return &SubsServer{
		rpcPort:     rpcPort,
		log:         log,
		subsService: subsService,
	}
}

func (ss *SubsServer) Listen(rpcPort int) error {
	listen, err := net.Listen("tcp", fmt.Sprintf("0.0.0.0:%d", rpcPort))
	if err != nil {
		return err
	}
	defer listen.Close()

	for {
		rpcConn, err := listen.Accept()
		if err != nil {
			continue
		}
		go rpc.ServeConn(rpcConn)
	}
}
