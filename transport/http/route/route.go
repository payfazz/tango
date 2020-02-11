package route

import (
	"net/http"

	"github.com/payfazz/go-apt/pkg/fazzrouter"
	"github.com/payfazz/tango/transport/container"
	"github.com/payfazz/tango/transport/http/route/collection"
)

// Compile is a function to compile the data
func Compile(app *container.AppContainer) http.HandlerFunc {
	r := fazzrouter.BaseRoute()
	r.Use(
		app.Middlewares.Auth,  // remove this line if app doesn't use authentication
		app.Middlewares.DB,    // remove this line if app doesn't use DB
		app.Middlewares.Redis, // remove this line if app doesn't use Redis
		app.Middlewares.Cors,
		app.Middlewares.Prometheus.RequestDuration,
		app.Middlewares.Prometheus.RequestCounter,
		app.Middlewares.Prometheus.StatusCounter,
	)
	collection.RouteBase(r, app)
	collection.RouteVersion1(r, app)

	return r.Compile()
}
