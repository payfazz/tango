package server

import (
	"fmt"
	"log"
	"net"

	"github.com/payfazz/tango/template/default/transport/grpc/middleware"

	"github.com/payfazz/tango/template/default/transport"

	"github.com/go-kit/kit/endpoint"
	"github.com/payfazz/go-apt/pkg/fazzdb"
	"github.com/payfazz/tango/template/default/config"
	"github.com/payfazz/tango/template/default/transport/container"
	"github.com/payfazz/tango/template/default/transport/grpc/app"
	"google.golang.org/grpc"
)

// grpcServer used for serving GRPC endpoint
type grpcServer struct {
	app *container.AppContainer
}

// Serve handle actual serving of GRPC endpoint
func (gs *grpcServer) Serve() {
	if config.Get(config.GRPC_FLAG) == config.OFF {
		return
	}

	go func() {
		listener, err := net.Listen("tcp", config.Get(config.GRPC_PORT))
		if nil != err {
			panic(err)
		}

		serverInstance := grpc.NewServer()

		// Call RegisterEndpoint from grpc-client repository
		// include baseMiddleware

		log.Println(fmt.Sprintf("GRPC Server running at port %s", config.Get(config.GRPC_PORT)))
		if err := serverInstance.Serve(listener); nil != err {
			panic(err)
		}
	}()
}

// CreateGrpcServer construct GRPC server
func CreateGrpcServer(app *container.AppContainer) transport.ServerInterface {
	return &grpcServer{
		app: app,
	}
}

func baseMiddleware() endpoint.Middleware {
	queryDb := fazzdb.QueryDb(config.GetDb(), config.GetQueryConfig()) // remove if db is not used
	rds := config.GetRedis()                                           // remove if redis is not used

	return endpoint.Chain(
		app.DB(queryDb), // remove if db is not used
		app.Redis(rds),  // remove if redis is not used
		middleware.TranslateToGrpcError(),
	)
}
