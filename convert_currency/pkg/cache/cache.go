package cache

import (
	"convert-service/pkg/server"
	"strings"
)

// Initialize Function in Cache Package
func init() {
	// Remote Cache Configuration Value
	switch strings.ToLower(server.Config.GetString("REDIS_DRIVER")) {
	case "redis":
		redisCfg.Host = server.Config.GetString("REDIS_HOST")
		redisCfg.Port = server.Config.GetString("REDIS_PORT")
		redisCfg.Password = server.Config.GetString("REDIS_PASSWORD")
		redisCfg.Name = server.Config.GetInt("REDIS_NAME")

		if len(redisCfg.Host) != 0 && len(redisCfg.Port) != 0 {
			RedisCache = redisConnect()
		}
	}
}
