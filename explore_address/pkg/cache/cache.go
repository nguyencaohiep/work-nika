package cache

import (
	"explore_address/pkg/server"
	"strings"
)

// Initialize Function in Cache Package
func init() {
	// Remote Cache Configuration Value
	switch strings.ToLower(server.Config.GetString("REMOTE_CACHE_DRIVER")) {
	case "redis":
		localCacheCfg.NumCounters = server.Config.GetInt64("LOCAL_CACHE_NUM_COUNTERS")
		localCacheCfg.MaxCost = server.Config.GetInt64("LOCAL_CACHE_MAX_COST")
		localCacheCfg.BufferItems = server.Config.GetInt64("LOCAL_CACHE_BUFFER_ITEMS")
		localCacheCfg.Metrics = server.Config.GetBool("LOCAL_CACHE_METRICS")

		redisCfg.Host = server.Config.GetString("REMOTE_CACHE_HOST")
		redisCfg.Port = server.Config.GetString("REMOTE_CACHE_PORT")
		redisCfg.Password = server.Config.GetString("REMOTE_CACHE_PASSWORD")
		redisCfg.Name = server.Config.GetInt("REMOTE_CACHE_NAME")

		redisCfg.Enable = server.Config.GetBool("CACHE_ENABLE")
		if redisCfg.Enable {
			RedisCache = redisConnect()
		}
	}
}
