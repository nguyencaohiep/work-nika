package stocks

import (
	"gonm_service/service/stocks/controller"

	"github.com/go-chi/chi"
)

var GonmService = chi.NewRouter()

// Init package with sub-route for price service
func init() {
	controller.GetExchange()
}
