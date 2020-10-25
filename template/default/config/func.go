package config

import (
	"fmt"
	"os"
	"reflect"
	"sort"
	"time"

	"github.com/payfazz/go-apt/pkg/fazzdb"
	"github.com/payfazz/go-apt/pkg/fazzthrottle"
)

// Set add / change value to base config
func Set(key string, value string) {
	base[key] = value
}

// Get get config by key from env then base config
func Get(key string) string {
	r := os.Getenv(key)
	if r != "" {
		return r
	}

	if configValue, ok := base[key]; ok {
		return configValue
	}

	return ""
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

// GetIfInt64 get config as int64
func GetIfInt64(key string) int64 {
	var t int64
	return getIf(key, t).(int64)
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
func GetIfLimitType(key string) fazzthrottle.LimitType {
	var t fazzthrottle.LimitType
	return getIf(key, t).(fazzthrottle.LimitType)
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

// PrintEnv prints environment variables sorted by key in ascending order.
func PrintEnv(w io.Writer) error {
	keys := make([]string, 0)
	for key, _ := range base {
		keys = append(keys, key)
	}
	sort.Strings(keys)

	for _, key := range keys {
		val := base[key]
		if _, err := w.Write([]byte(fmt.Sprintf("%s: \"%s\"\n", key, val))); err != nil {
			return err
		}
	}
	return nil
}
