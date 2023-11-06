package review

import (
	"score_gear5/service/score/controller"

	"github.com/go-chi/chi"
)

var ScoreServiceSunRouter = chi.NewRouter()

// Init package with sub-route for price service
func init() {
	controller.Score()
}
