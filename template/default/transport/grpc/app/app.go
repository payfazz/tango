package app

import (
	"context"

	"github.com/go-kit/kit/endpoint"
	"github.com/payfazz/go-apt/pkg/fazzdb"
	"github.com/payfazz/go-apt/pkg/fazzkv/redis"
)

func DB(queryDb *fazzdb.Query) endpoint.Middleware {
	return func(f endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, in interface{}) (out interface{}, err error) {
			ctx = fazzdb.NewQueryContext(ctx, queryDb)
			return f(ctx, in)
		}
	}
}

func Redis(rds redis.RedisInterface) endpoint.Middleware {
	return func(f endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, in interface{}) (out interface{}, err error) {
			ctx = redis.NewRedisContext(ctx, rds)
			return f(ctx, in)
		}
	}
}
