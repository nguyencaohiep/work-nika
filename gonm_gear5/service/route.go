package service

import (
	"gonm_service/pkg/router"
	"gonm_service/service/index"
	"gonm_service/service/stocks"
)

// LoadRoutes to Load Routes to Router
func LoadRoutes() {
	// Set Endpoint for admin
	router.Router.Get(router.RouterBasePath+"/", index.GetIndex)
	router.Router.Get(router.RouterBasePath+"/health", index.GetHealth)
	router.Router.Mount(router.RouterBasePath+"/gonm", stocks.GonmService)
}
