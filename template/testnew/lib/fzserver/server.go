package fzserver

type ServerInterface interface {
	Serve(config *Config) error
	Shutdown() error
}
