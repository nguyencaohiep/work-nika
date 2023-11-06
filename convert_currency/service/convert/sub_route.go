package convert

import (
	"convert-service/service/convert/controller"

	"github.com/go-chi/chi"
)

var ConvertServiceSubRoute = chi.NewRouter()

func init() {

	ConvertServiceSubRoute.Group(func(r chi.Router) {
		ConvertServiceSubRoute.Get("/detail", controller.GetInfoPrice)
		ConvertServiceSubRoute.Get("/info-crypto", controller.GetInfoCrypto)
		ConvertServiceSubRoute.Get("/info-currency", controller.GetInfoCurrency)
	})
}
