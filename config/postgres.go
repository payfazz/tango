package config

import (
	"database/sql"

	"github.com/payfazz/go-apt/pkg/fazzdb"
)

var postgresConfig = map[string]string{
	DB_HOST: "localhost",
	DB_PORT: "5432",
	DB_NAME: "wuilly-inbox",
	DB_USER: "postgres",
	DB_PASS: "postgres",

	DB_SLAVE_HOST: "localhost",
	DB_SLAVE_PORT: "5432",
	DB_SLAVE_NAME: "wuilly-inbox",
	DB_SLAVE_USER: "postgres",
	DB_SLAVE_PASS: "postgres",

	DB_MIGRATE_USER: "postgres",
	DB_MIGRATE_PASS: "postgres",
	FORCE_MIGRATE:   ON,

	MAX_OPEN_CONNS: "1024",
	MAX_IDLE_CONNS: "512",

	TEST_DB_HOST:         "localhost",
	TEST_DB_PORT:         "5432",
	TEST_DB_NAME:         "wuilly-inbox-test",
	TEST_DB_USER:         "postgres",
	TEST_DB_PASS:         "postgres",
	TEST_DB_MIGRATE_USER: "postgres",
	TEST_DB_MIGRATE_PASS: "postgres",
}

var postgresInterface = map[string]interface{}{
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

	DB_SLAVE_HOST   = "DB_SLAVE_HOST"
	DB_SLAVE_PORT   = "DB_SLAVE_PORT"
	DB_SLAVE_NAME   = "DB_SLAVE_NAME"
	DB_SLAVE_USER   = "DB_SLAVE_USER"
	DB_SLAVE_PASS   = "DB_SLAVE_PASS"
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
