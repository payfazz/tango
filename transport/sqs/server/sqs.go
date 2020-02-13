package server

import (
	"context"

	"github.com/aws/aws-sdk-go/service/sqs"
	"github.com/payfazz/go-apt/pkg/fazzdb"
	"github.com/payfazz/go-apt/pkg/fazzkv/redis"
	"github.com/payfazz/tango/config"
	"github.com/payfazz/tango/transport"
	"github.com/payfazz/tango/transport/sqs/route"
	"github.com/payfazz/venlog/pkg/venlogaws"
	"github.com/payfazz/venlog/pkg/venlogsubs"
)

// MonitorServer used for serving monitor server
type sqsServer struct{}

// Serve handle actual serving of sqs server
func (ss *sqsServer) Serve() {
	if config.Get(config.SQS_FLAG) == config.OFF {
		return
	}

	sqsClient := sqs.New(config.GetSqsAwsSession())

	subs := &venlogsubs.Subscription{
		Source:  venlogaws.NewSQSSource(sqsClient, config.GetReceiveMessageInput()),
		Handler: route.Route(),
	}
	subs.Watch(sqsContext())
}

// CreateSqsServer construct SqsServer
func CreateSqsServer() transport.ServerInterface {
	return &sqsServer{}
}

func sqsContext() context.Context {
	queryDb := fazzdb.QueryDb(config.GetDb(), config.GetQueryConfig())
	rds := config.GetRedis()

	ctx := context.Background()

	ctx = fazzdb.NewQueryContext(ctx, queryDb)
	ctx = redis.NewRedisContext(ctx, rds)

	return ctx
}
