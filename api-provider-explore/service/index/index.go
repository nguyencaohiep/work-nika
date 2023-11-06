package index

import (
	"explore_address/pkg/cache"
	"explore_address/pkg/db"
	"explore_address/pkg/router"
	"explore_address/pkg/server"
	"net/http"
	"strings"
)

// GetIndex Function to Show API Information
func GetIndex(w http.ResponseWriter, r *http.Request) {
	router.ResponseSuccess(w, "", "Go Framework is running")
}

// GetHealth Function to Show Health Check Status
func GetHealth(w http.ResponseWriter, r *http.Request) {
	// Check Database Connections
	if len(server.Config.GetString("DB_DRIVER")) != 0 {
		switch strings.ToLower(server.Config.GetString("DB_DRIVER")) {
		case "mysql":
			err := db.MySQL.Ping()
			if err != nil {
				router.ResponseInternalError(w, "mysql-health-check", err)
				return
			}
		case "postgres":
			err := db.PSQL.Ping()
			if err != nil {
				router.ResponseInternalError(w, "postgres-health-check", err)
				return
			}
		case "mongo":
			err := db.MongoSession.Ping()
			if err != nil {
				router.ResponseInternalError(w, "mongo-health-check", err)
				return
			}
		}
	}

	if len(server.Config.GetString("REMOTE_CACHE_DRIVER")) != 0 {
		switch strings.ToLower(server.Config.GetString("REMOTE_CACHE_DRIVER")) {
		case "redis":
			_, err := cache.RedisCache.Ping()
			if err != nil {
				router.ResponseInternalError(w, "redis-health-check", err)
				return
			}
		}
	}

	// Return Success response
	router.ResponseSuccess(w, "", "Health is ok")
}
