package main

import (
	"github.com/payfazz/tango/template/default/config"
	"github.com/payfazz/tango/template/default/transport/container"
	grpcServer "github.com/payfazz/tango/template/default/transport/grpc/server"
	httpServer "github.com/payfazz/tango/template/default/transport/http/server"
	monitorServer "github.com/payfazz/tango/template/default/transport/monitor/server"
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
