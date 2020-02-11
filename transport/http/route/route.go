package route

import (
	"net/http"

	"github.com/payfazz/go-apt/pkg/fazzrouter"
	"github.com/payfazz/tango/transport/container"
	"github.com/payfazz/tango/transport/http/route/collection"
)

// Compile compile the data http endpoint and middleware
func Compile(app *container.AppContainer) http.HandlerFunc {
	r := fazzrouter.BaseRoute()
	r.Use(
		app.HttpMiddleware.Auth,  // remove this line if app doesn't use authentication
		app.HttpMiddleware.DB,    // remove this line if app doesn't use DB
		app.HttpMiddleware.Redis, // remove this line if app doesn't use Redis
		app.HttpMiddleware.Cors,
		app.HttpMiddleware.Prometheus.RequestDuration,
		app.HttpMiddleware.Prometheus.RequestCounter,
		app.HttpMiddleware.Prometheus.StatusCounter,
	)
	collection.RouteBase(r, app)
	collection.RouteVersion1(r, app)

	return r.Compile()
}
