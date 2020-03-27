package main

import (
	"github.com/payfazz/go-apt/pkg/fazzdb"
	"github.com/payfazz/tango/config"
	"github.com/payfazz/tango/database/migration"
)

var testMigrations = []fazzdb.MigrationVersion{}

func main() {
	runningMigrations := append(migration.Sequence, testMigrations...)
	fazzdb.Migrate(
		config.GetMigrateDb(),
		"test-tango-backend",
		config.ForceMigrate(),
		config.RunSeeder(),
		runningMigrations...,
	)
}
