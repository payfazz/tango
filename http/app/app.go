package app

import (
	"context"
	"net/http"

	"github.com/payfazz/go-apt/pkg/fazzdb"
	"github.com/payfazz/go-apt/pkg/fazzkv/redis"
	"github.com/payfazz/tango/config"
)

type appKeyType struct{}

var appKey appKeyType

// App is a struct that will be send in the context
type App struct {
	AuthId    string
	UserAgent string
	Token     string
}

// New is a function that as a http handler
func New(rds redis.RedisInterface) func(http.HandlerFunc) http.HandlerFunc {
	return func(next http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			queryDb := fazzdb.QueryDb(config.GetDb(), config.GetQueryConfig())

			ctx := r.Context()
			ctx = NewAppContext(ctx)
			ctx = redis.NewRedisContext(ctx, rds)
			ctx = fazzdb.NewQueryContext(ctx, queryDb)

			next(w, r.WithContext(ctx))
		}
	}
}

// GetApp is a function that used to get the app context
func GetApp(ctx context.Context) *App {
	return ctx.Value(appKey).(*App)
}

// NewAppContext is a function that used to create new app context
func NewAppContext(ctx context.Context) context.Context {
	return context.WithValue(ctx, appKey, &App{})
}
