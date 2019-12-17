package server

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/payfazz/tango/config"
	"github.com/payfazz/tango/http/container"
	"github.com/payfazz/tango/http/route"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

// ApiServer empty struct
type ApiServer struct{}

func (as *ApiServer) runPrometheus() {
	if config.Get(config.PROMET_FLAG) == config.ON {
		go func() {
			http.Handle("/metrics", promhttp.Handler())
			log.Println(http.ListenAndServe(config.Get(config.PROMET_PORT), nil))
		}()
	}
}

// Serve is a function to serve the server
func (as *ApiServer) Serve() {
	services := container.CreateServiceContainer()

	s := &http.Server{
		Addr:         config.Get(config.PORT),
		ReadTimeout:  config.GetIfDuration(config.I_READ_TIMEOUT),
		WriteTimeout: config.GetIfDuration(config.I_WRITE_TIMEOUT),
		IdleTimeout:  config.GetIfDuration(config.I_IDLE_TIMEOUT),
		Handler:      route.Compile(services),
	}

	serverErrCh := make(chan error)
	go func() {
		defer close(serverErrCh)
		log.Println(fmt.Sprintf("Server running at port %s", config.Get(config.PORT)))
		serverErrCh <- s.ListenAndServe()
	}()

	as.runPrometheus()

	signalChan := make(chan os.Signal, 1)
	signals := []os.Signal{syscall.SIGTERM, syscall.SIGINT}
	signal.Notify(signalChan, signals...)

	select {
	case err := <-serverErrCh:
		log.Println("Server returning error: ", err)
	case sig := <-signalChan:
		signal.Reset(signals...)
		waitFor := config.GetIfDuration(config.I_WAIT_SHUTDOWN)

		log.Printf("Got '%s' signal, Stopping (Waiting for graceful shutdown: %s)\n", sig.String(), waitFor.String())

		ctx, cancel := context.WithTimeout(context.Background(), waitFor)
		defer cancel()

		if nil != s.Shutdown(ctx) {
			log.Println("Shutting down server returning error", s.Shutdown(ctx))
		} else {
			log.Println("Shutting down server")
		}
	}
}

// CreateApiServer is a function as a constructor to create an API server
func CreateApiServer() *ApiServer {
	return &ApiServer{}
}
