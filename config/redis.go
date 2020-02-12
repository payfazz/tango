package config

var redisConfig = map[string]string{
	REDIS_HOST: "localhost:6379",
	REDIS_PASS: "redis",

	TEST_REDIS_HOST: "localhost:6379",
	TEST_REDIS_PASS: "redis",
}

const (
	REDIS_HOST = "REDIS_HOST"
	REDIS_PASS = "REDIS_PASS"

	TEST_REDIS_HOST = "TEST_REDIS_HOST"
	TEST_REDIS_PASS = "TEST_REDIS_PASS"
)
