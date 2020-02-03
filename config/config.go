package config

import (
	"database/sql"
	"time"

	_ "github.com/lib/pq"
	"github.com/payfazz/go-apt/pkg/fazzdb"
	"github.com/payfazz/tango/lib/fazzthrottle/value"
)

var base = map[string]string{
	// APP
	// ENV ENV_PRODUCTION | ENV_STAGING | ENV_DEVELOPMENT | ENV_TESTING
	ENV: ENV_DEVELOPMENT,
	// DB_HOST host for DB
	DB_HOST: "localhost",
	// DB_PORT port for DB
	DB_PORT: "5432",
	// DB_NAME name for DB
	DB_NAME: "tango",
	// DB_USER user for DB
	DB_USER: "postgres",
	// DB_PASS pass for DB
	DB_PASS: "postgres",
	// DB_SLAVE_HOST host for DB_SLAVE
	DB_SLAVE_HOST: "localhost",
	// DB_SLAVE_PORT port for DB_SLAVE
	DB_SLAVE_PORT: "5432",
	// DB_SLAVE_NAME name for DB_SLAVE
	DB_SLAVE_NAME: "tango",
	// DB_SLAVE_USER user for DB_SLAVE
	DB_SLAVE_USER: "postgres",
	// DB_SLAVE_PASS pass for DB_SLAVE
	DB_SLAVE_PASS: "postgres",
	// DB_MIGRATE_USER user for DB migrate
	DB_MIGRATE_USER: "postgres",
	// DB_MIGRATE_PASS pass for DB migrate
	DB_MIGRATE_PASS: "postgres",
	// FORCE_MIGRATE default: on
	FORCE_MIGRATE: ON,
	// MAX_OPEN_CONNS: MAX_OPEN_CONNS
	MAX_OPEN_CONNS: "1024",
	// MAX_IDLE_CONNS MAX_IDLE_CONNS
	MAX_IDLE_CONNS: "512",
	// REDIS_HOST host for redis, format: {host}:{port}
	REDIS_HOST: "localhost:6379",
	// REDIS_PASS pass for redis
	REDIS_PASS: "redis",
	// PORT default port for server
	PORT: ":8080",
	// BASE_URL:
	BASE_URL: "http://localhost:8080",
	// DEBUG_LOG default: on
	DEBUG_LOG: ON,
	// THROTTLE_PREFIX default: throttle
	THROTTLE_PREFIX: "throttle",
	// THROTTLE_TRIGGER default: on
	THROTTLE_TRIGGER: ON,
	// GRPC_PORT :1301
	GRPC_PORT: ":1301",

	// TEST
	// TEST_DB_HOST localhost
	TEST_DB_HOST: "localhost",
	// TEST_DB_PORT 5432
	TEST_DB_PORT: "5432",
	// TEST_DB_NAME go-test
	TEST_DB_NAME: "tango",
	// TEST_DB_USER postgres
	TEST_DB_USER: "postgres",
	// TEST_DB_PASS postgres
	TEST_DB_PASS: "postgres",
	// TEST_DB_MIGRATE_USER postgres
	TEST_DB_MIGRATE_USER: "postgres",
	// TEST_DB_MIGRATE_PASS postgres
	TEST_DB_MIGRATE_PASS: "postgres",
	// TEST_REDIS_HOST localhost:6379
	TEST_REDIS_HOST: "localhost:6379",
	// TEST_REDIS_PASS redis
	TEST_REDIS_PASS: "redis",
}

var baseInterface = map[string]interface{}{
	// I_QUERY_CONFIG conf for DB limit, offset and lock
	I_QUERY_CONFIG: fazzdb.Config{
		Limit:           20,
		Offset:          0,
		Lock:            fazzdb.LO_NONE,
		DevelopmentMode: Get(ENV) != ENV_PRODUCTION,
	},
	// I_SLAVE_QUERY_CONFIG conf for slave DB limit, offset and lock
	I_SLAVE_QUERY_CONFIG: fazzdb.Config{
		Limit:           20,
		Offset:          0,
		Lock:            fazzdb.LO_NONE,
		DevelopmentMode: Get(ENV) != ENV_PRODUCTION,
		Opts: &sql.TxOptions{
			Isolation: sql.LevelDefault,
			ReadOnly:  true,
		},
	},
	// I_READ_TIMEOUT read timeout
	I_READ_TIMEOUT: 5 * time.Minute,
	// I_WRITE_TIMEOUT write timeout
	I_WRITE_TIMEOUT: 5 * time.Minute,
	// I_IDLE_TIMEOUT idle timeout
	I_IDLE_TIMEOUT: 30 * time.Second,
	// I_WAIT_SHUTDOWN graceful shutdown delay
	I_WAIT_SHUTDOWN: 5 * time.Second,
	// I_CREDENTIAL_EXPIRE credential expired time
	I_CREDENTIAL_EXPIRE: 24 * 25 * time.Hour,
	// I_REDIS_SESSION_EXPIRE const for redis expired session
	I_REDIS_SESSION_EXPIRE: 24 * time.Hour,
	// I_THROTTLE_LIMIT limit per hit for throttling
	I_THROTTLE_LIMIT: 40,
	// I_THROTTLE_DURATION duration for throttle TTL
	I_THROTTLE_DURATION: 10 * time.Second,
	// I_THROTTLE_TYPE type checker for throttle
	I_THROTTLE_TYPE: value.IP_ENDPOINT,
}
