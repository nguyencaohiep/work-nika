package service

import (
	"venture-data-service/pkg/router"
	"venture-data-service/service/index"
	"venture-data-service/service/venture"
	"venture-data-service/service/venture/crawler"

	"github.com/go-chi/chi/middleware"
)

// LoadRoutes to Load Routes to Router
func LoadRoutes() {

	router.Router.Use(middleware.RealIP)

	//* Set Endpoint for admin
	router.Router.Get(router.RouterBasePath+"/", index.GetIndex)
	router.Router.Get(router.RouterBasePath+"/health", index.GetHealth)
	router.Router.Mount(router.RouterBasePath+"/venture_nika", venture.VentureService)

	crawler.GenContact()
}
