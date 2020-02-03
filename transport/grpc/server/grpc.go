package server

import (
	"net"

	"github.com/payfazz/tango/config"

	"google.golang.org/grpc"
)

type GrpcServer struct{}

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

func CreateGrpcServer() *GrpcServer {
	return &GrpcServer{}
}
