package test

import (
	"context"

	"github.com/payfazz/go-apt/pkg/fazzdb"
	"github.com/payfazz/go-apt/pkg/fazzkv/redis"
	"github.com/payfazz/tango/config"
	"github.com/payfazz/tango/http/app"
)

// PrepareTest function that connecting the test environment to the testing db and testing redis
func PrepareTest() context.Context {
	config.Set(config.ENV, config.ENV_TESTING)
	queryDb := fazzdb.QueryDb(config.GetDb(), config.GetQueryConfig())

	ctx := context.Background()
	ctx = app.NewAppContext(ctx)
	ctx = redis.NewRedisContext(ctx, config.GetRedis())
	ctx = fazzdb.NewQueryContext(ctx, queryDb)

	return ctx
}

// Truncate is a function that used to truncate all testing table
func Truncate(ctx context.Context, tables ...string) error {
	q, _ := fazzdb.GetQueryContext(ctx)
	_, err := q.Truncate(tables...)
	if err != nil {
		return err
	}
	return nil
}
