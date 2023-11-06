package service

import (
	"explore_address/pkg/router"
	"explore_address/service/address"
	"explore_address/service/index"
	"time"

	"github.com/go-chi/chi/middleware"
)

// LoadRoutes to Load Routes to Router
func LoadRoutes() {
	//* Set middleware for router
	router.Router.Use(middleware.RealIP)
	router.Router.Use(middleware.Timeout(60 * time.Second))
	router.Router.Use(middleware.Compress(5, "application/json"))

	//* Set endpoint for admin
	router.Router.Get(router.RouterBasePath+"/", index.GetIndex)
	router.Router.Get(router.RouterBasePath+"/health", index.GetHealth)
	router.Router.Mount(router.RouterBasePath+"/explore", address.ExploreAddressServiceSubRoute)

	//* Set endpoint for client
}
