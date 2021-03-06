package container

import (
	"github.com/payfazz/go-apt/pkg/fazzthrottle"
	"net/http"

	"github.com/payfazz/go-apt/pkg/fazzdb"
	prometheus "github.com/payfazz/go-apt/pkg/fazzrouter/middleware"
	"github.com/payfazz/tango/template/default/config"
	"github.com/payfazz/tango/template/default/transport/http/app"
	"github.com/payfazz/tango/template/default/transport/http/middleware"
)

// MiddlewareContainer handle all middleware used in project
type MiddlewareContainer struct {
	Auth       func(next http.HandlerFunc) http.HandlerFunc
	DB         func(next http.HandlerFunc) http.HandlerFunc
	Cors       func(next http.HandlerFunc) http.HandlerFunc
	Redis      func(next http.HandlerFunc) http.HandlerFunc
	Throttle   func(next http.HandlerFunc) http.HandlerFunc
	Prometheus *prometheusMiddleware
}

type prometheusMiddleware struct {
	RequestCounter  func(next http.HandlerFunc) http.HandlerFunc
	RequestDuration func(next http.HandlerFunc) http.HandlerFunc
}

// CreateMiddlewareContainer construct all middlewares used in the app
func CreateMiddlewareContainer() *MiddlewareContainer {
	return &MiddlewareContainer{
		Auth:       createAuth(),
		DB:         createDB(),
		Cors:       middleware.Cors(),
		Redis:      createRedis(),
		Throttle:   createThrottle(),
		Prometheus: createPrometheus(),
	}
}

func createPrometheus() *prometheusMiddleware {
	return &prometheusMiddleware{
		RequestCounter:  prometheus.HTTPRequestCounterMiddleware(),
		RequestDuration: prometheus.HTTPRequestDurationMiddleware(),
	}
}

func createThrottle() func(next http.HandlerFunc) http.HandlerFunc {
	return fazzthrottle.MiddlewarePrefix(
		config.Get(config.THROTTLE_PREFIX),
		config.GetIfInteger(config.I_THROTTLE_LIMIT),
		config.GetIfDuration(config.I_THROTTLE_DURATION),
		config.GetIfLimitType(config.I_THROTTLE_TYPE),
		config.UseThrottle(),
		config.GetRedis,
	)
}

func createAuth() func(next http.HandlerFunc) http.HandlerFunc {
	return app.Auth()
}

func createRedis() func(next http.HandlerFunc) http.HandlerFunc {
	rds := config.GetRedis()
	return app.Redis(rds)
}

func createDB() func(next http.HandlerFunc) http.HandlerFunc {
	queryDb := fazzdb.QueryDb(config.GetDb(), config.GetQueryConfig())
	return app.DB(queryDb)
}
