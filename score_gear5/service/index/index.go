package index

import (
	"net/http"
	"score_gear5/pkg/db"
	"score_gear5/pkg/router"
	"score_gear5/pkg/server"
	"strings"
)

// GetIndex Function to Show API Information
func GetIndex(w http.ResponseWriter, r *http.Request) {
	router.ResponseSuccess(w, "200", "Score is running")
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
}
