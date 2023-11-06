package index

import (
	"holder_contract/pkg/db"
	"holder_contract/pkg/router"
	"holder_contract/pkg/server"
	"net/http"
	"strings"
)

// GetIndex Function to Show API Information
func GetIndex(w http.ResponseWriter, r *http.Request) {
	router.ResponseSuccess(w, "200", "Go Framework is running")
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

	// Return Success response
	router.ResponseSuccess(w, "200", "Health is ok")
}
