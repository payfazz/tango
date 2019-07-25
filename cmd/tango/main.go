package main

import (
	"github.com/payfazz/go-apt/pkg/fazzdb"
	"github.com/payfazz/tango/config"
	"github.com/payfazz/tango/database/migration"
	"github.com/payfazz/tango/http/server"
)

func main() {
	fazzdb.Migrate(config.GetMigrateDb(),
		"go-backend",
		config.ForceMigrate(),
		true,
		migration.Sequence...,
	)
	config.SetVerboseQuery()

	api := server.CreateApiServer()
	api.Serve()
}
