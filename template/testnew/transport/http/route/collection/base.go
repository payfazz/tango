package collection

import (
	"github.com/payfazz/go-apt/pkg/fazzrouter"
	"github.com/payfazz/tango/template/default/transport/container"
	"github.com/payfazz/tango/template/default/transport/http/controller/base"
)

// RouteBase default route for app
func RouteBase(r *fazzrouter.Route, app *container.AppContainer) {
	r.Get("ping", base.Ping())
}
