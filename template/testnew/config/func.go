package config

import (
	"time"

	"github.com/payfazz/go-apt/pkg/fazzdb"
	"github.com/payfazz/go-apt/pkg/fazzthrottle"
)

// Set add / change value to base config
func Set(key string, value string) {
	Config.Set(key, value)
}

// Get get config by key from env then base config
func Get(key string) string {
	return Config.Get(key)
}

// GetIfString get config as string
func GetIfString(key string) string {
	var t string
	return Config.GetInterface(key, t).(string)
}

// GetIfInteger get config as int
func GetIfInteger(key string) int {
	var t int
	return Config.GetInterface(key, t).(int)
}

// GetIfInt64 get config as int64
func GetIfInt64(key string) int64 {
	var t int64
	return Config.GetInterface(key, t).(int64)
}

// GetIfDuration get config as duration
func GetIfDuration(key string) time.Duration {
	var t time.Duration
	return Config.GetInterface(key, t).(time.Duration)
}

// GetIfQueryConfig get config as fazzdb.Config
func GetIfQueryConfig(key string) fazzdb.Config {
	var t fazzdb.Config
	return Config.GetInterface(key, t).(fazzdb.Config)
}

// GetIfLimitType get config as fazzthrottle.LimitType
func GetIfLimitType(key string) fazzthrottle.LimitType {
	var t fazzthrottle.LimitType
	return Config.GetInterface(key, t).(fazzthrottle.LimitType)
}

// RunSeeder only run seeder on development environment
func RunSeeder() bool {
	return Get(ENV) == ENV_DEVELOPMENT
}

// UseThrottle get throttle trigger status
func UseThrottle() bool {
	return Get(THROTTLE_TRIGGER) == ON
}

func PrintEnv() {
	Config.PrintEnv()
}
