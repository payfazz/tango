package config

import (
	"fmt"
	"os"
	"reflect"
	"sync"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/payfazz/go-apt/pkg/fazzcommon/formatter"
	"github.com/payfazz/go-apt/pkg/fazzdb"
	"github.com/payfazz/go-apt/pkg/fazzkv/redis"
	"github.com/payfazz/tango/lib/fazzthrottle/value"
)

var rds redis.RedisInterface
var mdb, db, sdb *sqlx.DB
var rdsOnce, mdbOnce, dbOnce, sdbOnce sync.Once

// Set is a function to append value to base config
func Set(key string, value string) {
	base[key] = value
}

// GetRedis is a function to get Redis instance
func GetRedis() redis.RedisInterface {
	rdsOnce.Do(func() {
		var err error

		switch Get(ENV) {
		case ENV_TESTING:
			rds, err = redis.NewFazzRedis(
				Get(TEST_REDIS_HOST),
				Get(TEST_REDIS_PASS),
			)
		default:
			rds, err = redis.NewFazzRedis(
				Get(REDIS_HOST),
				Get(REDIS_PASS),
			)
		}
		if nil != err {
			panic(err)
		}
	})
	return rds
}

// GetMigrateDb is a function to get DB Migration instance
func GetMigrateDb() *sqlx.DB {
	mdbOnce.Do(func() {
		var err error
		var conn string
		switch Get(ENV) {
		case ENV_TESTING:
			conn = fmt.Sprintf(
				"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
				Get(TEST_DB_HOST),
				Get(TEST_DB_PORT),
				Get(TEST_DB_MIGRATE_USER),
				Get(TEST_DB_MIGRATE_PASS),
				Get(TEST_DB_NAME),
			)
		default:
			conn = fmt.Sprintf(
				"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
				Get(DB_HOST),
				Get(DB_PORT),
				Get(DB_MIGRATE_USER),
				Get(DB_MIGRATE_PASS),
				Get(DB_NAME),
			)
		}

		mdb, err = sqlx.Connect("postgres", conn)
		if nil != err {
			panic(err)
		}
	})
	return mdb
}

// GetDB get DB instance
func GetDb() *sqlx.DB {
	dbOnce.Do(func() {
		var err error
		var conn string
		switch Get(ENV) {
		case ENV_TESTING:
			conn = fmt.Sprintf(
				"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
				Get(TEST_DB_HOST),
				Get(TEST_DB_PORT),
				Get(TEST_DB_USER),
				Get(TEST_DB_PASS),
				Get(TEST_DB_NAME),
			)
		default:
			conn = fmt.Sprintf(
				"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
				Get(DB_HOST),
				Get(DB_PORT),
				Get(DB_USER),
				Get(DB_PASS),
				Get(DB_NAME),
			)
		}
		db, err = sqlx.Connect("postgres", conn)
		if nil != err {
			panic(err)
		}

		db.SetMaxOpenConns(formatter.StringToInteger(Get(MAX_OPEN_CONNS)))
		db.SetMaxIdleConns(formatter.StringToInteger(Get(MAX_IDLE_CONNS)))
	})
	return db
}

// GetDB get slave DB instance
func GetSlaveDb() *sqlx.DB {
	sdbOnce.Do(func() {
		var err error
		var conn string
		switch Get(ENV) {
		case ENV_TESTING:
			conn = fmt.Sprintf(
				"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
				Get(TEST_DB_HOST),
				Get(TEST_DB_PORT),
				Get(TEST_DB_USER),
				Get(TEST_DB_PASS),
				Get(TEST_DB_NAME),
			)
		default:
			conn = fmt.Sprintf(
				"host=%s port=%s user=%s password=%s DB_SLAVEname=%s sslmode=disable",
				Get(DB_SLAVE_HOST),
				Get(DB_SLAVE_PORT),
				Get(DB_SLAVE_USER),
				Get(DB_SLAVE_PASS),
				Get(DB_SLAVE_NAME),
			)
		}
		sdb, err = sqlx.Connect("postgres", conn)
		if nil != err {
			panic(err)
		}

		sdb.SetMaxOpenConns(formatter.StringToInteger(Get(MAX_OPEN_CONNS)))
		sdb.SetMaxIdleConns(formatter.StringToInteger(Get(MAX_IDLE_CONNS)))
	})
	return sdb
}

// GetQueryConfig is a function to get default query config
func GetQueryConfig() fazzdb.Config {
	return GetIfQueryConfig(I_QUERY_CONFIG)
}

// GetSlaveQueryConfig is a function to get default slave query config
func GetSlaveQueryConfig() fazzdb.Config {
	return GetIfQueryConfig(I_SLAVE_QUERY_CONFIG)
}

// SetVerboseQuery is a function to set verbose mode on fazzdb
func SetVerboseQuery() {
	if OFF == Get(DEBUG_LOG) || ENV_PRODUCTION == Get(ENV) {
		return
	}

	fazzdb.Verbose()
}

// Get config by key
func Get(key string) string {
	r := os.Getenv(key)
	if r == "" {
		if _, ok := base[key]; !ok {
			return ""
		}
		r = base[key]
	}
	return r
}

// GetIfString get config as string
func GetIfString(key string) string {
	var t string
	return getIf(key, t).(string)
}

// GetIfInteger get config as int
func GetIfInteger(key string) int {
	var t int
	return getIf(key, t).(int)
}

// GetIfDuration get config as duration
func GetIfDuration(key string) time.Duration {
	var t time.Duration
	return getIf(key, t).(time.Duration)
}

// GetIfQueryConfig get config as fazzdb.Config
func GetIfQueryConfig(key string) fazzdb.Config {
	var t fazzdb.Config
	return getIf(key, t).(fazzdb.Config)
}

// GetIfLimitType get config as fazzthrottle.LimitType
func GetIfLimitType(key string) value.LimitType {
	var t value.LimitType
	return getIf(key, t).(value.LimitType)
}

// ForceMigrate get force migrate status depending on FORCE_MIGRATE and current ENV
func ForceMigrate() bool {
	return Get(ENV) != ENV_PRODUCTION && Get(FORCE_MIGRATE) == ON
}

// RunSeeder only run seeder on development environment
func RunSeeder() bool {
	return Get(ENV) == ENV_DEVELOPMENT
}

// UseThrottle get throttle trigger status
func UseThrottle() bool {
	return Get(THROTTLE_TRIGGER) == ON
}

func getIf(key string, p interface{}) interface{} {
	t := reflect.TypeOf(p)

	if _, ok := baseInterface[key]; !ok {
		panic(fmt.Sprintf(`config with key "%s" not found`, key))
	}

	if keyType := reflect.TypeOf(baseInterface[key]); t != keyType {
		panic(fmt.Sprintf(`different type of config with key "%s" got %s expected %s`, key, keyType, t))
	}

	return baseInterface[key]
}
