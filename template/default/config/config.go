package config

import (
	_ "github.com/lib/pq"
)

const (
	MAX_PER_PAGE        = 20
	DEFAULT_FORMAT_DEEP = 8
)

var basicConfig = map[string]string{
	SERVICE_NAME: "tango",
	ENV:          ENV_DEVELOPMENT,

	THROTTLE_TRIGGER: OFF,
	PROMET_FLAG:      ON,
	GRPC_FLAG:        ON,
	HTTP_FLAG:        ON,

	DEBUG_LOG: ON,
}

var base = mergeConfig(
	basicConfig,
	appConfig,
	postgresConfig,
	redisConfig,
	prometheusConfig,
	httpConfig,
	grpcConfig,
)

var baseInterface = mergeConfigInterface(
	appInterfaceConfig,
	postgresInterfaceConfig,
	httpInterfaceConfig,
)
