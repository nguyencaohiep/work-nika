package price

import (
	"holder_contract/service/price/controller"
	"holder_contract/service/price/crawler"
	"net/http"
	"time"

	"github.com/go-chi/chi"
)

var CrawlHolderServiceSunRouter = chi.NewRouter()

func init() {
	go func() {
		for {
			crawler.ListenQueueETH()
		}
	}()

	go func() {
		for {
			crawler.ListenQueueBSC()
		}
	}()

	go func() {
		for {
			time.Sleep(24 * time.Hour)
			crawler.MapInfoBSC = map[string]crawler.InfoEle{}
		}
	}()

	go func() {
		for {
			time.Sleep(24 * time.Hour)
			crawler.MapInfoETH = map[string]crawler.InfoEle{}
		}
	}()

	// amount, err := crawler.CrawlHoldersETH("0x0000000000085d4780b73119b644ae5ecd22b376")
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Println(amount)

	// amount, err = crawler.CrawlHoldersBSC("")
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Println(amount)

	CrawlHolderServiceSunRouter.Group(func(r chi.Router) {
		CrawlHolderServiceSunRouter.Handle("/info-eth", http.HandlerFunc(controller.GetInfoETH))
		CrawlHolderServiceSunRouter.Handle("/info-bsc", http.HandlerFunc(controller.GetInfoBSC))
		CrawlHolderServiceSunRouter.Handle("/info-eth/list", http.HandlerFunc(controller.GetListInfoETH))
		CrawlHolderServiceSunRouter.Handle("/info-bsc/list", http.HandlerFunc(controller.GetListInfoBSC))
		CrawlHolderServiceSunRouter.Handle("/info-bsc/list-handle", http.HandlerFunc(controller.GetListInfoBSCHanddle))
	})
}
