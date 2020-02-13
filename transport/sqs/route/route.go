package route

import (
	"context"
	"fmt"

	"github.com/payfazz/tango/transport/sqs/controller"
	"github.com/payfazz/venlog/pkg/venlog"
)

var routes = map[string]func(ctx context.Context, event *venlog.Event) error{
	"empty": controller.Empty,
}

func Route() func(ctx context.Context, event *venlog.Event) error {
	return func(ctx context.Context, event *venlog.Event) error {
		key := fmt.Sprintf("%s-%s", event.Category, event.Type)

		if "" == event.Category {
			return routes["empty"](ctx, event)
		}

		return routes[key](ctx, event)
	}
}
