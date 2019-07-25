package test

import (
	"github.com/payfazz/go-apt/pkg/fazzdb"
	"github.com/payfazz/tango/config"
	"github.com/payfazz/tango/database/migration"
)

func main() {
	config.Set(config.ENV, config.ENV_TESTING)

	// Running application migration table without seeder
	fazzdb.Migrate(
		config.GetMigrateDb(),
		"tango-test",
		true,
		false,
		migration.Sequence...,
	)

	// Running test seeder
	runTestSeeder()
}

func runTestSeeder() {
	// Test Migration Sequence
	sequence := []fazzdb.MigrationVersion{}

	// Running test migration sequence
	fazzdb.Migrate(
		config.GetMigrateDb(),
		"tango-test",
		false,
		true,
		sequence...,
	)
}
