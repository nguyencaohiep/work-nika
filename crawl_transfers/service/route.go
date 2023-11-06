package service

import (
	"admin_service/pkg/router"
	review "admin_service/service/crawl"
	"admin_service/service/index"
)

// LoadRoutes to Load Routes to Router
func LoadRoutes() {
	// Set Endpoint for admin
	router.Router.Get(router.RouterBasePath+"/", index.GetIndex)
	router.Router.Get(router.RouterBasePath+"/health", index.GetHealth)
	router.Router.Mount(router.RouterBasePath+"/reviews", review.AdminReviewServiceSunRouter)
}
