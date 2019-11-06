package main

import (
	"github.com/payfazz/tango/config"
	"github.com/payfazz/tango/http/server"
)

func main() {
	config.SetVerboseQuery()
	api := server.CreateApiServer()
	api.Serve()
}
