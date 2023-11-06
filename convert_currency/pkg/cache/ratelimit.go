package cache

import (
	"convert-service/pkg/log"
	"convert-service/pkg/router"
	"fmt"
	"net/http"
	"time"
)

func RateLimitByIP(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		//Get user ip address
		IPAddress := r.Header.Get("X-Real-Ip")
		if IPAddress == "" {
			IPAddress = r.Header.Get("X-Forwarded-For")
		}
		if IPAddress == "" {
			IPAddress = r.RemoteAddr
		}

		_, found, err := RedisCache.Get(IPAddress)
		if err != nil {
			log.Println(log.LogLevelError, "RateLimitByIP  RedisCache.Get(IPAddress)"+IPAddress, err.Error())
			next.ServeHTTP(w, r)
			return
		}
		if found {
			router.ResponseTooManyRequests(w)
			return
		}

		//Check ip rate limit
		count, dur, isUnderLimit := RedisCache.Limiter.Allow(IPAddress, 1, 3*time.Second)
		// fmt.Println(count, dur, isUnderLimit)

		if !isUnderLimit {
			log.Println(log.LogLevelInfo, fmt.Sprintf("RateLimitByIP: count: %d, dur: %v, isUnderLimit: %v, ip: %s", count, dur, isUnderLimit, IPAddress), "")
			router.ResponseTooManyRequests(w)
			//todo: cache this ip into blacklist (banned for 1 day)
			err := RedisCache.SetByKey(IPAddress, IPAddress, 24*time.Hour)
			if err != nil {
				log.Println(log.LogLevelError, "RateLimitByIP RedisCache.SetByKey"+IPAddress, err.Error())
			}

			return
		} else {
			next.ServeHTTP(w, r)
		}
	})
}
