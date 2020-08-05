package config

import (
	_ "github.com/lib/pq"
	"github.com/payfazz/tango/template/default/lib/fzserver"
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

var Config = fzserver.BuildConfig(fzserver.ConfigPayload{
	BaseMaps: []map[string]string{
		basicConfig,
		appConfig,
		postgresConfig,
		redisConfig,
		prometheusConfig,
		httpConfig,
		grpcConfig,
	},
	InterfaceMaps: []map[string]interface{}{
		appInterfaceConfig,
		postgresInterfaceConfig,
		httpInterfaceConfig,
	},
})
