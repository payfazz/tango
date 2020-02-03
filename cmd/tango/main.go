package main

import (
	"github.com/payfazz/tango/config"
	grpcServer "github.com/payfazz/tango/transport/grpc/server"
	httpServer "github.com/payfazz/tango/transport/http/server"
	prometheusServer "github.com/payfazz/tango/transport/prometheus/server"
)

func main() {
	config.SetVerboseQuery()

	api := httpServer.CreateApiServer()
	api.Serve()

	grpc := grpcServer.CreateGrpcServer()
	grpc.Serve()

	promet := prometheusServer.CreatePrometheusServer()
	promet.Serve()
}
