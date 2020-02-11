package server

import (
	"log"
	"net/http"

	"github.com/payfazz/tango/transport"

	"github.com/payfazz/go-apt/pkg/fazzmonitor/ping/messagebroker"

	"github.com/payfazz/go-apt/pkg/fazzmonitor/ping/database"

	"github.com/payfazz/go-apt/pkg/fazzmonitor/ping"

	"github.com/payfazz/tango/config"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

// MonitorServer used for serving monitor server
type MonitorServer struct{}

// Serve handle actual serving of monitor server
func (ps *MonitorServer) Serve() {
	if config.Get(config.PROMET_FLAG) == config.ON {
		go func() {
			http.Handle("/metrics", promhttp.Handler())
			http.Handle("/ping", ping.Ping(config.Get(config.SERVICE_NAME), reportChecks()))
			log.Println(http.ListenAndServe(config.Get(config.PROMET_PORT), nil))
		}()
	}
}

// CreateMonitorServer construct MonitorServer
func CreateMonitorServer() transport.ServerInterface {
	return &MonitorServer{}
}

func reportChecks() []ping.ReportInterface {
	return []ping.ReportInterface{
		database.NewPgSQLReport(
			config.Get(config.DB_HOST),
			config.Get(config.DB_PORT),
			config.Get(config.DB_USER),
			config.Get(config.DB_PASS),
			config.Get(config.DB_NAME),
		),
		messagebroker.NewRedisReportWithAddress(
			config.Get(config.REDIS_HOST),
			config.Get(config.REDIS_PASS),
		),
	}
}
