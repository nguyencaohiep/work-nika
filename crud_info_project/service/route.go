package service

import (
	"info_project_service/pkg/router"
	"info_project_service/service/index"
	"info_project_service/service/info"
)

// LoadRoutes to Load Routes to Router
func LoadRoutes() {
	// Set Endpoint for admin
	router.Router.Get(router.RouterBasePath+"/", index.GetIndex)
	router.Router.Get(router.RouterBasePath+"/health", index.GetHealth)

	// Mount services sub routes to the main router
	router.Router.Mount(router.RouterBasePath+"/projectInfo", info.ProjectInfoServiceSunRouter)
}
