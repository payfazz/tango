package controller

import (
	"context"
	"encoding/json"
	"log"

	"github.com/payfazz/venlog/pkg/venlog"
)

func Empty(ctx context.Context, event *venlog.Event) error {
	by, err := json.Marshal(event)
	if nil != err {
		return err
	}

	log.Println("received empty event:", string(by))
	return nil
}
