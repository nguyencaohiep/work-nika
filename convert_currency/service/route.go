package service

import (
	"convert-service/pkg/log"
	"convert-service/pkg/router"
	"convert-service/service/convert"
	"convert-service/service/convert/crawler"

	"convert-service/service/index"
	"time"

	"github.com/go-chi/chi/middleware"
)

// LoadRoutes to Load Routes to Router
func LoadRoutes() {
	//* Set middleware for router
	router.Router.Use(middleware.RealIP)
	router.Router.Use(middleware.Timeout(60 * time.Second))
	router.Router.Use(middleware.Compress(5, "application/json"))

	//* Set endpoint for admin
	router.Router.Get(router.RouterBasePath+"/", index.GetIndex)
	router.Router.Get(router.RouterBasePath+"/health", index.GetHealth)
	router.Router.Mount(router.RouterBasePath+"/convert", convert.ConvertServiceSubRoute)

	go func() {
		for {
			crawler.GetTopPriceInfoBinance()
			time.Sleep(2 * time.Second)
		}
	}()

	go func() {
		for {
			crawler.CrawlPriceCoingecko()
			time.Sleep(20 * time.Minute)
		}
	}()

	go func() {
		for {
			// time.Sleep(12 * time.Hour)
			err := crawler.RepoCryptos.UpdateInfo()
			if err != nil {
				log.Println(log.LogLevelError, "crawler.RepoCryptos.UpdateInfo()", err.Error())
			}
		}
	}()

}
