package container

import (
	"net/http"

	"github.com/payfazz/go-apt/pkg/fazzdb"
	"github.com/payfazz/go-apt/pkg/fazzmonitor/prometheusclient"
	"github.com/payfazz/tango/config"
	"github.com/payfazz/tango/lib/fazzthrottle"
	"github.com/payfazz/tango/transport/http/app"
	"github.com/payfazz/tango/transport/http/middleware"
)

// MiddlewareContainer is a struct to handle all middleware used in project
type MiddlewareContainer struct {
	App        func(next http.HandlerFunc) http.HandlerFunc
	DB         func(next http.HandlerFunc) http.HandlerFunc
	Cors       func(next http.HandlerFunc) http.HandlerFunc
	Redis      func(next http.HandlerFunc) http.HandlerFunc
	Throttle   func(next http.HandlerFunc) http.HandlerFunc
	Prometheus *prometheusMiddleware
}

type prometheusMiddleware struct {
	RequestCounter  func(next http.HandlerFunc) http.HandlerFunc
	RequestDuration func(next http.HandlerFunc) http.HandlerFunc
	StatusCounter   func(next http.HandlerFunc) http.HandlerFunc
}

// CreateMiddlewareContainer is a constructor for creating all middlewares used in the app
func CreateMiddlewareContainer() *MiddlewareContainer {
	return &MiddlewareContainer{
		App:        createApp(),
		DB:         createDB(),
		Cors:       middleware.Cors(),
		Redis:      createRedis(),
		Throttle:   createThrottle(),
		Prometheus: createPrometheus(),
	}
}

func createPrometheus() *prometheusMiddleware {
	return &prometheusMiddleware{
		RequestCounter:  prometheusclient.RequestCounter(),
		RequestDuration: prometheusclient.RequestDuration(),
		StatusCounter:   prometheusclient.StatusCounter(),
	}
}

func createThrottle() func(next http.HandlerFunc) http.HandlerFunc {
	return fazzthrottle.MiddlewarePrefix(
		config.Get(config.THROTTLE_PREFIX),
		config.GetIfInteger(config.I_THROTTLE_LIMIT),
		config.GetIfDuration(config.I_THROTTLE_DURATION),
		config.GetIfLimitType(config.I_THROTTLE_TYPE),
		config.UseThrottle(),
	)
}

func createApp() func(next http.HandlerFunc) http.HandlerFunc {
	return app.New()
}

func createRedis() func(next http.HandlerFunc) http.HandlerFunc {
	rds := config.GetRedis()
	return app.Redis(rds)
}

func createDB() func(next http.HandlerFunc) http.HandlerFunc {
	queryDb := fazzdb.QueryDb(config.GetDb(), config.GetQueryConfig())
	return app.DB(queryDb)
}
