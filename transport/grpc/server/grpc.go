package server

import (
	"net"

	"github.com/payfazz/tango/config"
	"github.com/payfazz/tango/transport/container"
	"google.golang.org/grpc"
)

type GrpcServer struct {
	app *container.AppContainer
}

func (gs *GrpcServer) Serve() {
	go func() {
		listener, err := net.Listen("tcp", config.Get(config.GRPC_PORT))
		if nil != err {
			panic(err)
		}

		serverInstance := grpc.NewServer()

		// Call RegisterEndpoint from grpc-client repository

		if err := serverInstance.Serve(listener); nil != err {
			panic(err)
		}
	}()
}

func CreateGrpcServer(app *container.AppContainer) *GrpcServer {
	return &GrpcServer{
		app: app,
	}
}
