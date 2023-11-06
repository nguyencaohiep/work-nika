package service

import (
	"holder_contract/pkg/router"
	"holder_contract/service/index"
	"holder_contract/service/price"

	"github.com/go-chi/chi/middleware"
)

// LoadRoutes to Load Routes to Router
func LoadRoutes() {

	router.Router.Use(middleware.RealIP)

	//* Set Endpoint for admin
	router.Router.Get(router.RouterBasePath+"/", index.GetIndex)
	router.Router.Get(router.RouterBasePath+"/health", index.GetHealth)
	router.Router.Mount(router.RouterBasePath+"/holder-contract", price.CrawlHolderServiceSunRouter)
}
