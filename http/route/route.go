package route

import (
	"net/http"

	"github.com/payfazz/go-apt/pkg/fazzmonitor/prometheusclient"

	"github.com/payfazz/go-apt/pkg/fazzrouter"
	"github.com/payfazz/tango/config"
	"github.com/payfazz/tango/http/app"
	"github.com/payfazz/tango/http/container"
	"github.com/payfazz/tango/http/middleware"
	"github.com/payfazz/tango/http/route/collection"
)

// Compile is a function to compile the data
func Compile(services *container.ServiceContainer) http.HandlerFunc {
	rds := config.GetRedis()

	r := fazzrouter.BaseRoute()
	r.Use(
		app.New(rds),
		middleware.Cors(),
		prometheusclient.RequestCounter(),
		prometheusclient.RequestDuration(),
		prometheusclient.StatusCounter(),
	)
	collection.RouteBase(r, services)
	collection.RouteVersion1(r, services)

	return r.Compile()
}
