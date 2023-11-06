package cache

import (
	"explore_address/pkg/router"
	"net/http"
	"time"
)

func RateLimitByIP(maxRate int64, durationSecond int) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			_, _, isUnderLimit := RedisCache.Limiter.Allow(r.RemoteAddr, maxRate, time.Duration(durationSecond)*time.Second)
			if !isUnderLimit {
				router.ResponseTooManyRequests(w, "")
				return
			} else {
				next.ServeHTTP(w, r)
			}
		})
	}
}
