package migrate

import (
	"github.com/payfazz/go-apt/pkg/fazzdb"
	"github.com/payfazz/tango/config"
	"github.com/payfazz/tango/database/migration"
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
