package main

import (
	"github.com/payfazz/tango/config"
	"github.com/payfazz/tango/transport/container"
	grpcServer "github.com/payfazz/tango/transport/grpc/server"
	httpServer "github.com/payfazz/tango/transport/http/server"
	monitorServer "github.com/payfazz/tango/transport/monitor/server"
)

func main() {
	config.SetVerboseQuery()

	app := container.CreateAppContainer()

	grpc := grpcServer.CreateGrpcServer(app)
	grpc.Serve()

	monitor := monitorServer.CreateMonitorServer()
	monitor.Serve()

	http := httpServer.CreateHttpServer(app)
	http.Serve()
}
