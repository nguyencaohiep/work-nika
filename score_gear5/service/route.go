package service

import (
	"score_gear5/pkg/router"
	"score_gear5/service/index"
	review "score_gear5/service/score"
)

// LoadRoutes to Load Routes to Router
func LoadRoutes() {
	// Set Endpoint for admin
	router.Router.Get(router.RouterBasePath+"/", index.GetIndex)
	router.Router.Get(router.RouterBasePath+"/health", index.GetHealth)
	router.Router.Mount(router.RouterBasePath+"/scores", review.ScoreServiceSunRouter)
}
