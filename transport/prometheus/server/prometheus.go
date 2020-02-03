package server

import (
	"log"
	"net/http"

	"github.com/payfazz/tango/config"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

type PrometheusServer struct{}

func (ps *PrometheusServer) Serve() {
	if config.Get(config.PROMET_FLAG) == config.ON {
		go func() {
			http.Handle("/metrics", promhttp.Handler())
			log.Println(http.ListenAndServe(config.Get(config.PROMET_PORT), nil))
		}()
	}
}

func CreatePrometheusServer() *PrometheusServer {
	return &PrometheusServer{}
}
