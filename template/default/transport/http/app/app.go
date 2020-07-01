package app

import (
	"context"
	"net/http"

	"github.com/payfazz/go-apt/pkg/fazzdb"
	"github.com/payfazz/go-apt/pkg/fazzkv/redis"
)

type authKeyType struct{}

var authKey authKeyType

// Authentication hold authentication credentials to be passed into context
type Authentication struct {
	AuthId    string
	UserAgent string
	Token     string
}

// Redis middleware to append redis into context
func Redis(rds redis.RedisInterface) func(next http.HandlerFunc) http.HandlerFunc {
	return func(next http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			ctx := redis.NewRedisContext(r.Context(), rds)
			next(w, r.WithContext(ctx))
		}
	}
}

// DB middleware to append queryDb into context
func DB(queryDb *fazzdb.Query) func(next http.HandlerFunc) http.HandlerFunc {
	return func(next http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			ctx := fazzdb.NewQueryContext(r.Context(), queryDb)
			next(w, r.WithContext(ctx))
		}
	}
}

// Auth middleware to append for Authentication
func Auth() func(http.HandlerFunc) http.HandlerFunc {
	return func(next http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			ctx := NewAuthContext(r.Context())
			next(w, r.WithContext(ctx))
		}
	}
}

// GetApp is a function that used to get the app context
func GetApp(ctx context.Context) *Authentication {
	return ctx.Value(authKey).(*Authentication)
}

// NewAppContext is a function that used to create new app context
func NewAuthContext(ctx context.Context) context.Context {
	return context.WithValue(ctx, authKey, &Authentication{})
}
