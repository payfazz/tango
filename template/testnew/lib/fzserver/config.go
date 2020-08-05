package fzserver

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/payfazz/go-apt/pkg/fazzcommon/formatter"
	"github.com/payfazz/go-apt/pkg/fazzconfig"
	"github.com/payfazz/go-apt/pkg/fazzconfig/configsource"
	"github.com/payfazz/go-apt/pkg/fazzkv/redis"
	"github.com/payfazz/tango/template/default/lib/fzserver/value"
	"reflect"
	"sort"
	"sync"
)

type Config struct {
	baseMap                  map[string]string
	interfaceMap             map[string]interface{}
	reader                   fazzconfig.ConfigReader
	mdb, db                  *sqlx.DB
	rds                      redis.RedisInterface
	mdbOnce, dbOnce, rdsOnce sync.Once
}

type ConfigPayload struct {
	BaseMaps      []map[string]string
	InterfaceMaps []map[string]interface{}
}

func (c *Config) Set(key string, value string) {
	c.baseMap[key] = value
}

func (c *Config) Get(key string) string {
	return c.reader.Get(key)
}

func (c *Config) GetNotEmpty(key string) string {
	val := c.reader.Get(key)
	if val == "" {
		panic(fmt.Sprintf(`config with key "%s" not found`, key))
	}
	return c.reader.Get(key)
}

func (c *Config) GetInterface(key string, sample interface{}) interface{} {
	t := reflect.TypeOf(sample)

	if _, ok := c.interfaceMap[key]; !ok {
		panic(fmt.Sprintf(`config with key "%s" not found`, key))
	}

	if keyType := reflect.TypeOf(c.interfaceMap[key]); t != keyType {
		panic(fmt.Sprintf(`different type of config with key "%s" got %s expected %s`, key, keyType, t))
	}

	return c.interfaceMap[key]
}

// GetMigrateDb create DB Migration instance
func (c *Config) GetMigrateDb() *sqlx.DB {
	c.mdbOnce.Do(func() {
		var err error
		var conn string
		switch c.Get(value.ENV) {
		case value.ENV_TESTING:
			conn = fmt.Sprintf(
				"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
				c.GetNotEmpty(value.TEST_DB_HOST),
				c.GetNotEmpty(value.TEST_DB_PORT),
				c.GetNotEmpty(value.TEST_DB_MIGRATE_USER),
				c.GetNotEmpty(value.TEST_DB_MIGRATE_PASS),
				c.GetNotEmpty(value.TEST_DB_NAME),
			)
		default:
			conn = fmt.Sprintf(
				"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
				c.GetNotEmpty(value.DB_HOST),
				c.GetNotEmpty(value.DB_PORT),
				c.GetNotEmpty(value.DB_MIGRATE_USER),
				c.GetNotEmpty(value.DB_MIGRATE_PASS),
				c.GetNotEmpty(value.DB_NAME),
			)
		}

		c.mdb, err = sqlx.Connect("postgres", conn)
		if nil != err {
			panic(err)
		}
	})
	return c.mdb
}

// GetDB create DB instance
func (c *Config) GetDb() *sqlx.DB {
	c.dbOnce.Do(func() {
		var err error
		var conn string
		switch c.Get(value.ENV) {
		case value.ENV_TESTING:
			conn = fmt.Sprintf(
				"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
				c.GetNotEmpty(value.TEST_DB_HOST),
				c.GetNotEmpty(value.TEST_DB_PORT),
				c.GetNotEmpty(value.TEST_DB_USER),
				c.Get(value.TEST_DB_PASS),
				c.Get(value.TEST_DB_NAME),
			)
		default:
			conn = fmt.Sprintf(
				"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
				c.GetNotEmpty(value.DB_HOST),
				c.GetNotEmpty(value.DB_PORT),
				c.GetNotEmpty(value.DB_USER),
				c.GetNotEmpty(value.DB_PASS),
				c.GetNotEmpty(value.DB_NAME),
			)
		}
		c.db, err = sqlx.Connect("postgres", conn)
		if nil != err {
			panic(err)
		}

		c.db.SetMaxOpenConns(formatter.StringToInteger(c.GetNotEmpty(value.MAX_OPEN_CONNS)))
		c.db.SetMaxIdleConns(formatter.StringToInteger(c.GetNotEmpty(value.MAX_IDLE_CONNS)))
	})
	return c.db
}

// GetRedis create redis instance
func (c *Config) GetRedis() redis.RedisInterface {
	c.rdsOnce.Do(func() {
		var err error

		switch c.Get(value.ENV) {
		case value.ENV_TESTING:
			c.rds, err = redis.NewFazzRedis(
				c.Get(value.TEST_REDIS_HOST),
				c.Get(value.TEST_REDIS_PASS),
			)
		default:
			c.rds, err = redis.NewFazzRedis(
				c.Get(value.REDIS_HOST),
				c.Get(value.REDIS_PASS),
			)
		}
		if nil != err {
			panic(err)
		}
	})
	return c.rds
}

func (c *Config) PrintEnv() {
	keys := make([]string, 0)
	for key, _ := range c.baseMap {
		keys = append(keys, key)
	}
	sort.Strings(keys)

	for _, key := range keys {
		val := c.Get(key)
		fmt.Printf("%s: \"%s\"\n", key, val)
	}
}

func BuildConfig(payload ConfigPayload) *Config {
	baseMap := mergeConfig(payload.BaseMaps...)
	interfaceMap := mergeConfigInterface(payload.InterfaceMaps...)
	return &Config{
		baseMap:      baseMap,
		interfaceMap: interfaceMap,
		reader:       fazzconfig.NewReader(configsource.FromEnv(), configsource.FromMap(baseMap)),
	}
}

func mergeConfig(configs ...map[string]string) map[string]string {
	result := map[string]string{}

	for _, configMap := range configs {
		for key, configValue := range configMap {
			if _, ok := result[key]; ok {
				panic(fmt.Sprintf(`duplicate config key "%s" detected`, key))
			}

			result[key] = configValue
		}
	}

	return result
}

func mergeConfigInterface(configInterfaces ...map[string]interface{}) map[string]interface{} {
	result := map[string]interface{}{}

	for _, configInterfaceMap := range configInterfaces {
		for key, configValue := range configInterfaceMap {
			if _, ok := result[key]; ok {
				panic(fmt.Sprintf(`duplicate config interface key "%s" detected`, key))
			}

			result[key] = configValue
		}
	}

	return result
}
