package service

import (
	"native_service/pkg/router"
	"native_service/service/index"
	price_current "native_service/service/native"
)

// LoadRoutes to Load Routes to Router
func LoadRoutes() {
	// Set Endpoint for admin
	router.Router.Get(router.RouterBasePath+"/", index.GetIndex)
	router.Router.Get(router.RouterBasePath+"/health", index.GetHealth)

	// Mount services sub routes to the main router
	router.Router.Mount(router.RouterBasePath+"/price_current", price_current.PriceServiceSunRouter)
}
