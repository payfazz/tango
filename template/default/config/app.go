package config

import (
	"time"

	"github.com/payfazz/tango/template/default/lib/fazzthrottle/value"
)

var appConfig = map[string]string{
	THROTTLE_PREFIX: "tango-throttle",
}

var appInterfaceConfig = map[string]interface{}{
	I_THROTTLE_LIMIT:    40,
	I_THROTTLE_DURATION: 60 * time.Second,
	I_THROTTLE_TYPE:     value.IP_ENDPOINT,
}

const (
	SERVICE_NAME     = "SERVICE_NAME"
	ENV              = "ENV"
	DEBUG_LOG        = "DEBUG_LOG"
	THROTTLE_PREFIX  = "THROTTLE_PREFIX"
	THROTTLE_TRIGGER = "THROTTLE_TRIGGER"

	I_THROTTLE_LIMIT    = "I_THROTTLE_LIMIT"
	I_THROTTLE_DURATION = "I_THROTTLE_DURATION"
	I_THROTTLE_TYPE     = "I_THROTTLE_TYPE"
)
