package cache

import (
	"encoding/json"
	"explore_address/pkg/log"
	"sync"
	"time"

	"github.com/dgraph-io/ristretto"
	"github.com/go-redis/redis"
	"github.com/go-redis/redis_rate"
)

const (
	_remoteKeyNotExist  = redis.Nil
	_localCacheItemCost = 1
)

// Local Cache Configuration Struct
type localCacheConfig struct {
	MaxCost     int64
	NumCounters int64
	BufferItems int64
	Metrics     bool
}

// Local Cache Configuration Variable
var localCacheCfg localCacheConfig

// Redis Configuration Struct
type redisConfig struct {
	Host     string
	Port     string
	Password string
	Name     int
	Enable   bool
}

var redisCfg = &redisConfig{}

type RedisCacheStore struct {
	RWMutex sync.RWMutex
	Limiter *redis_rate.Limiter
	Remote  *redis.Client
	Local   *ristretto.Cache
}

// Redis Cache Variable
var RedisCache *RedisCacheStore

// Redis Connect Function
func redisConnect() *RedisCacheStore {
	// Initialize Connection
	remote := redis.NewClient(&redis.Options{
		Addr:     redisCfg.Host + ":" + redisCfg.Port,
		Password: redisCfg.Password,
		DB:       redisCfg.Name,
	})

	// Test remote connection
	_, err := remote.Ping().Result()
	if err != nil {
		log.Println(log.LogLevelFatal, "redis-connect", err.Error())
	}

	local, err := ristretto.NewCache(&ristretto.Config{
		MaxCost:     localCacheCfg.MaxCost,
		NumCounters: localCacheCfg.NumCounters,
		BufferItems: localCacheCfg.BufferItems,
		Metrics:     localCacheCfg.Metrics,
	})

	if err != nil {
		log.Println(log.LogLevelFatal, "redis-connect", err.Error())
	}

	// Return Connection
	return &RedisCacheStore{
		Limiter: redis_rate.NewLimiter(remote),
		Remote:  remote,
		Local:   local,
	}
}

// Get method to check redis server connection
func (redisCache *RedisCacheStore) Ping() (string, error) {
	return redisCache.Remote.Ping().Result()
}

// Set method to set cache by given key with time to live
func (redisCache *RedisCacheStore) Set(key string, value any, timeToLive time.Duration) error {
	if redisCfg.Enable {
		redisCache.RWMutex.Lock()
		defer redisCache.RWMutex.Unlock()

		byteValue, err := json.Marshal(value)
		if err != nil {
			return err
		}
		// Set to local cache
		redisCache.Local.SetWithTTL(key, byteValue, _localCacheItemCost, timeToLive/2)

		// Set to redis cache
		return redisCache.Remote.Set(key, byteValue, timeToLive).Err()
	}
	return nil
}

// Get method to retrieve the value of a key. If not present, returns false.
func (redisCache *RedisCacheStore) Get(key string, timeToLive time.Duration) ([]byte, bool, error) {
	if redisCfg.Enable {
		// Check local cache
		byteValue, found := redisCache.Local.Get(key)
		if found {
			return byteValue.([]byte), true, nil
		}

		// Get redis cache
		byteValue, err := redisCache.Remote.Get(key).Bytes()
		if err == _remoteKeyNotExist {
			return nil, false, nil
		}
		if err != nil {
			return nil, false, err
		}

		// Set to local cache
		redisCache.Local.SetWithTTL(key, byteValue, _localCacheItemCost, timeToLive/2)
		return byteValue.([]byte), true, nil
	}

	return nil, false, nil
}

// Invalidate  method to delete a key from cahce.
func (redisCache *RedisCacheStore) Invalidate(key string) error {
	if redisCfg.Enable {
		redisCache.RWMutex.Lock()
		defer redisCache.RWMutex.Unlock()

		// Invalidate local cache
		redisCache.Local.Del(key)

		return redisCache.Remote.Del(key).Err()
	}
	return nil
}

// Close method clear and then close the cache Store.
func (redisCache *RedisCacheStore) Close() {
	if redisCfg.Enable {
		redisCache.Local.Clear()
		redisCache.Local.Close()
		redisCache.Remote.Close()
	}
}
