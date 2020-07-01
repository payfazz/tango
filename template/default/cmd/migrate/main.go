package main

import (
	"github.com/payfazz/go-apt/pkg/fazzdb"
	"github.com/payfazz/tango/template/default/config"
	"github.com/payfazz/tango/template/default/database/migration"
)

func main() {
	config.SetVerboseQuery()
	fazzdb.Migrate(
		config.GetMigrateDb(),
		"tango-backend",
		config.ForceMigrate(),
		config.RunSeeder(),
		migration.Sequence...,
	)
}
