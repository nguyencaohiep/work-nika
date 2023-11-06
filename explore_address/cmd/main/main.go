package main

import (
	"log"
	"math/rand"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"explore_address/pkg/cache"
	"explore_address/pkg/db"
	"explore_address/pkg/router"
	"explore_address/pkg/server"
	"explore_address/pkg/utils"
	"explore_address/service"
)

// Server Variable
var svr *server.Server

// Init Function
func init() {
	// Set Go Log Flags
	log.SetFlags(log.Flags() &^ (log.Ldate | log.Ltime))

	// Load Routes
	service.LoadRoutes()

	// Initialize Server
	svr = server.NewServer(router.Router)

	// Set random seed
	rand.Seed(utils.TimeNowVietNam().UnixNano())
}

// Main Function
func main() {
	// Starting Server
	svr.Start()

	sig := make(chan os.Signal, 1)
	// Notify Any Signal to OS Signal Channel
	signal.Notify(sig, os.Interrupt)
	signal.Notify(sig, syscall.SIGTERM)

	// Return OS Signal Channel
	// As Exit Sign
	<-sig

	// Log Break Line
	log.Println("")

	// Stopping Server
	defer svr.Stop()

	// Close Any Database Connections
	if len(server.Config.GetString("DB_DRIVER")) != 0 {
		switch strings.ToLower(server.Config.GetString("DB_DRIVER")) {
		case "postgres":
			log.Println("Stoped postgres !")
			defer db.PSQL.Close()
		case "mysql":
			log.Println("Stoped mysql !")
			defer db.MySQL.Close()
		case "mongo":
			log.Println("Stoped mongo !")
			defer db.MongoSession.Close()
		}
	}

	if len(server.Config.GetString("REMOTE_CACHE_DRIVER")) != 0 {
		switch strings.ToLower(server.Config.GetString("REMOTE_CACHE_DRIVER")) {
		case "redis":
			log.Println("Stoped redis !")
			defer cache.RedisCache.Close()
		}
	}
}
