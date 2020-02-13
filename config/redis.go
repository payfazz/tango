package config

import (
	"sync"

	"github.com/payfazz/go-apt/pkg/fazzkv/redis"
)

var rds redis.RedisInterface
var rdsOnce sync.Once

var redisConfig = map[string]string{
	REDIS_HOST: "localhost:6379",
	REDIS_PASS: "redis",

	TEST_REDIS_HOST: "localhost:6379",
	TEST_REDIS_PASS: "redis",
}

// GetRedis create redis instance
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

const (
	REDIS_HOST = "REDIS_HOST"
	REDIS_PASS = "REDIS_PASS"

	TEST_REDIS_HOST = "TEST_REDIS_HOST"
	TEST_REDIS_PASS = "TEST_REDIS_PASS"
)
