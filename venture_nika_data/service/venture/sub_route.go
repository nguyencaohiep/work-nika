package venture

import (
	"github.com/go-chi/chi"
)

var VentureService = chi.NewRouter()

func init() {

	VentureService.Group(func(r chi.Router) {

	})
}
