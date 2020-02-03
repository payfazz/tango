package main

import (
	"github.com/payfazz/tango/config"
	"github.com/payfazz/tango/transport/container"
	grpcServer "github.com/payfazz/tango/transport/grpc/server"
	httpServer "github.com/payfazz/tango/transport/http/server"
	prometheusServer "github.com/payfazz/tango/transport/prometheus/server"
)

func main() {
	config.SetVerboseQuery()

	app := container.CreateAppContainer()

	api := httpServer.CreateApiServer(app)
	api.Serve()

	grpc := grpcServer.CreateGrpcServer(app)
	grpc.Serve()

	promet := prometheusServer.CreatePrometheusServer()
	promet.Serve()
}
