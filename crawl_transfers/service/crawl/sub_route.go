package review

import (
	"github.com/go-chi/chi"
)

var AdminReviewServiceSunRouter = chi.NewRouter()

// Init package with sub-route for price service
func init() {

	AdminReviewServiceSunRouter.Group(func(r chi.Router) {

	})
}
