package service

import (
	"price_service/pkg/router"
	"price_service/service/index"
	price_current "price_service/service/price"
)

// LoadRoutes to Load Routes to Router
func LoadRoutes() {
	// Set Endpoint for admin
	router.Router.Get(router.RouterBasePath+"/", index.GetIndex)
	router.Router.Get(router.RouterBasePath+"/health", index.GetHealth)

	// Mount services sub routes to the main router
	router.Router.Mount(router.RouterBasePath+"/prices", price_current.PriceServiceSunRouter)
}
