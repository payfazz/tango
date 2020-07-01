package fazzthrottle

import (
	"fmt"
	"net"
	"net/http"
	"strings"
	"time"

	"github.com/payfazz/go-apt/pkg/fazzcommon/formatter"
	"github.com/payfazz/go-apt/pkg/fazzcommon/httpError"
	"github.com/payfazz/go-apt/pkg/fazzcommon/response"
	"github.com/payfazz/go-apt/pkg/fazzkv/redis"
	"github.com/payfazz/tango/template/default/config"
	"github.com/payfazz/tango/template/default/lib/fazzthrottle/value"
)

// REDIS_NIL_ERROR redis: nil
//noinspection GoSnakeCaseUsage
const REDIS_NIL_ERROR = "redis: nil"

// LimiterInterface interface for throttling middleware
type LimiterInterface interface {
	Compile(limit int, duration time.Duration, limitType value.LimitType) func(next http.HandlerFunc) http.HandlerFunc
}

// Limiter struct for the limiter
type Limiter struct {
	Prefix string
}

func (l *Limiter) hit(r redis.RedisInterface, key string, limit int, duration time.Duration) (bool, error) {
	result, err := r.Get(key)
	if "" == result {
		errRedis := r.SetWithExpire(key, 1, duration)
		if nil != errRedis {
			return false, errRedis
		}

		return true, nil
	}
	if nil != err && REDIS_NIL_ERROR != err.Error() {
		return false, err
	}

	count := formatter.StringToInteger(result)
	if count > limit {
		return false, nil
	}

	err = r.Increment(key)
	if nil != err {
		return false, err
	}

	return true, nil
}

// Compile compile middleware and process the limitter
func (l *Limiter) Compile(
	limit int,
	duration time.Duration,
	limitType value.LimitType,
	useThrottle bool,
) func(next http.HandlerFunc) http.HandlerFunc {
	return func(next http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			var err error
			var allowed bool

			if useThrottle {
				endpoint := r.URL.Path
				ip := l.getIP(r)
				rds := config.GetRedis()

				if value.IP == limitType {
					allowed, err = l.hit(rds, fmt.Sprintf("%s:%s", l.Prefix, ip), limit, duration)
				} else if value.ENDPOINT == limitType {
					allowed, err = l.hit(rds, fmt.Sprintf("%s:%s", l.Prefix, endpoint), limit, duration)
				} else if value.IP_ENDPOINT == limitType {
					allowed, err = l.hit(rds, fmt.Sprintf("%s:%s:%s", l.Prefix, ip, endpoint), limit, duration)
				}

				if nil != err {
					response.Error(w, err)
					return
				}

				if !allowed {
					response.Error(w, httpError.TooManyRequest("rate limit reached"))
					return
				}
			}

			next(w, r)
		}
	}
}

func (l *Limiter) getIP(r *http.Request) string {
	ip := r.Header.Get("X-Forwarded-For")
	ip = strings.TrimSpace(strings.Split(ip, ",")[0])

	if ip == "" {
		ip = strings.TrimSpace(r.Header.Get("X-Real-Ip"))
	}

	if ip != "" {
		return ip
	}

	if ip, _, err := net.SplitHostPort(strings.TrimSpace(r.RemoteAddr)); err == nil {
		return ip
	}

	return ""
}

// MiddlewarePrefix construct throttle middleware with prefix
func MiddlewarePrefix(
	prefix string,
	limit int,
	duration time.Duration,
	limitType value.LimitType,
	useThrottle bool,
) func(next http.HandlerFunc) http.HandlerFunc {
	l := &Limiter{
		Prefix: prefix,
	}
	return l.Compile(limit, duration, limitType, useThrottle)
}

// Middleware construct throttle middleware without any prefix
func Middleware(
	limit int,
	duration time.Duration,
	limitType value.LimitType,
	useThrottle bool,
) func(next http.HandlerFunc) http.HandlerFunc {
	l := &Limiter{
		Prefix: "fazz",
	}
	return l.Compile(limit, duration, limitType, useThrottle)
}
