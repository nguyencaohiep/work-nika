package price

import (
	"native_service/service/native/crawler"

	"github.com/go-chi/chi"
)

var PriceServiceSunRouter = chi.NewRouter()

func init() {
	crawler.CrawlPriceNative()
}
