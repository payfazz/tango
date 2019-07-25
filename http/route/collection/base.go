package collection

import (
	"github.com/payfazz/go-apt/pkg/fazzrouter"
	"github.com/payfazz/tango/http/container"
	"github.com/payfazz/tango/http/controller/base"
)

// RouteBase is a base router
func RouteBase(r *fazzrouter.Route, services *container.ServiceContainer) {
	r.Get("ping", base.Ping())
}
