package app

import (
	"context"
	"net/http"

	"github.com/payfazz/go-apt/pkg/fazzdb"
	"github.com/payfazz/go-apt/pkg/fazzkv/redis"
)

type appKeyType struct{}

var appKey appKeyType

// App is a struct that will be send in the context
type App struct {
	AuthId    string
	UserAgent string
	Token     string
}

// Redis is a function that construct handler to append redis into context
func Redis(rds redis.RedisInterface) func(next http.HandlerFunc) http.HandlerFunc {
	return func(next http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			ctx := redis.NewRedisContext(r.Context(), rds)
			next(w, r.WithContext(ctx))
		}
	}
}

// DB is a function that construct handler to append queryDb into context
func DB(queryDb *fazzdb.Query) func(next http.HandlerFunc) http.HandlerFunc {
	return func(next http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			ctx := fazzdb.NewQueryContext(r.Context(), queryDb)
			next(w, r.WithContext(ctx))
		}
	}
}

// New is a function that construct handler for AppContext
func New() func(http.HandlerFunc) http.HandlerFunc {
	return func(next http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			ctx := NewAppContext(r.Context())
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
