package config

import (
	"time"
)

var httpConfig = map[string]string{
	HTTP_PORT: ":8080",
}

var httpInterfaceConfig = map[string]interface{}{
	I_READ_TIMEOUT:  5 * time.Minute,
	I_WRITE_TIMEOUT: 5 * time.Minute,
	I_IDLE_TIMEOUT:  30 * time.Second,
	I_WAIT_SHUTDOWN: 5 * time.Second,
}

const (
	HTTP_FLAG = "HTTP_FLAG"
	HTTP_PORT = "HTTP_PORT"

	I_READ_TIMEOUT  = "I_READ_TIMEOUT"
	I_WRITE_TIMEOUT = "I_WRITE_TIMEOUT"
	I_IDLE_TIMEOUT  = "I_IDLE_TIMEOUT"
	I_WAIT_SHUTDOWN = "I_WAIT_SHUTDOWN"
)
