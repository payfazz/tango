package base

import (
	"net/http"

	"github.com/payfazz/go-apt/pkg/fazzcommon/response"
)

// Ping check if service status is up or down
func Ping() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		response.Text(w, "ready to serve", http.StatusOK)
		return
	}
}
