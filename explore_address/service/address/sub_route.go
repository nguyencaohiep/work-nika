package address

import (
	controller "explore_address/service/address/controller/logs"

	"github.com/go-chi/chi"
)

var ExploreAddressServiceSubRoute = chi.NewRouter()

func init() {

	controller.ListenTransfersETH()

	ExploreAddressServiceSubRoute.Group(func(r chi.Router) {
	})
}
