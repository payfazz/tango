package config

import (
	"database/sql"
	"github.com/jmoiron/sqlx"
	"github.com/payfazz/go-apt/pkg/fazzdb"
)

var postgresConfig = map[string]string{
	DB_HOST: "localhost",
	DB_PORT: "5432",
	DB_NAME: "tango",
	DB_USER: "postgres",
	DB_PASS: "postgres",

	DB_MIGRATE_USER: "postgres",
	DB_MIGRATE_PASS: "postgres",
	FORCE_MIGRATE:   ON,

	MAX_OPEN_CONNS: "1024",
	MAX_IDLE_CONNS: "512",

	TEST_DB_HOST:         "localhost",
	TEST_DB_PORT:         "5432",
	TEST_DB_NAME:         "tango-test",
	TEST_DB_USER:         "postgres",
	TEST_DB_PASS:         "postgres",
	TEST_DB_MIGRATE_USER: "postgres",
	TEST_DB_MIGRATE_PASS: "postgres",
}

// GetMigrateDb create DB Migration instance
func GetMigrateDb() *sqlx.DB {
	return Config.GetMigrateDb()
}

// GetDB create DB instance
func GetDb() *sqlx.DB {
	return Config.GetDb()
}

// GetQueryConfig get default query config
func GetQueryConfig() fazzdb.Config {
	return GetIfQueryConfig(I_QUERY_CONFIG)
}

// SetVerboseQuery set verbose mode on fazzdb
func SetVerboseQuery() {
	if OFF == Get(DEBUG_LOG) || ENV_PRODUCTION == Get(ENV) {
		return
	}

	fazzdb.Verbose()
}

// ForceMigrate get force migrate status depending on FORCE_MIGRATE and current ENV
func ForceMigrate() bool {
	return Get(ENV) != ENV_PRODUCTION && Get(FORCE_MIGRATE) == ON
}

var postgresInterfaceConfig = map[string]interface{}{
	I_QUERY_CONFIG: fazzdb.Config{
		Limit:           MAX_PER_PAGE,
		Offset:          0,
		Lock:            fazzdb.LO_NONE,
		DevelopmentMode: Get(ENV) != ENV_PRODUCTION,
	},
	I_SLAVE_QUERY_CONFIG: fazzdb.Config{
		Limit:           MAX_PER_PAGE,
		Offset:          0,
		Lock:            fazzdb.LO_NONE,
		DevelopmentMode: Get(ENV) != ENV_PRODUCTION,
		Opts: &sql.TxOptions{
			Isolation: sql.LevelDefault,
			ReadOnly:  true,
		},
	},
}

const (
	DB_HOST = "DB_HOST"
	DB_PORT = "DB_PORT"
	DB_NAME = "DB_NAME"
	DB_USER = "DB_USER"
	DB_PASS = "DB_PASS"

	DB_MIGRATE_USER = "DB_MIGRATE_USER"
	DB_MIGRATE_PASS = "DB_MIGRATE_PASS"

	FORCE_MIGRATE  = "FORCE_MIGRATE"
	MAX_OPEN_CONNS = "MAX_OPEN_CONNS"
	MAX_IDLE_CONNS = "MAX_IDLE_CONNS"

	TEST_DB_HOST = "TEST_DB_HOST"
	TEST_DB_PORT = "TEST_DB_PORT"
	TEST_DB_NAME = "TEST_DB_NAME"
	TEST_DB_USER = "TEST_DB_USER"
	TEST_DB_PASS = "TEST_DB_PASS"

	TEST_DB_MIGRATE_USER = "TEST_DB_MIGRATE_USER"
	TEST_DB_MIGRATE_PASS = "TEST_DB_MIGRATE_PASS"

	I_QUERY_CONFIG       = "I_QUERY_CONFIG"
	I_SLAVE_QUERY_CONFIG = "I_SLAVE_QUERY_CONFIG"
)
