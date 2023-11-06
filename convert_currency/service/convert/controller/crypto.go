package controller

import (
	"convert-service/pkg/router"
	"convert-service/service/convert/crawler"
	"net/http"
	"sync"
)

var (
	mutex sync.Mutex
)

func GetInfoPrice(w http.ResponseWriter, r *http.Request) {
	symbol := r.URL.Query().Get("symbol")
	if len(symbol) <= 0 {
		router.ResponseBadRequest(w, "B.400", "Missing symbol param!")
		return
	}

	index, exist := crawler.MapIndexCrypto[symbol]
	if !exist {
		router.ResponseBadRequest(w, "B.400", "Not found!")
		return
	}

	router.ResponseCreatedWithData(w, "B.200", "Get info successfully", crawler.RepoCryptos.Cryptos[index])
}

func GetInfoCurrency(w http.ResponseWriter, r *http.Request) {
	mutex.Lock()
	res := crawler.MapCurrencies
	mutex.Unlock()
	router.ResponseCreatedWithData(w, "B.200", "Get info successfully", res)
}

func GetInfoCrypto(w http.ResponseWriter, r *http.Request) {
	repo := []crawler.CryptoInfo{}

	for i, crypto := range crawler.RepoCryptos.Cryptos {
		if len(crypto.CurrentPrice) > 0 {
			repo = append(repo, crawler.ListCompactInfoCrypto[i])
		}
	}

	router.ResponseCreatedWithData(w, "B.200", "Get info successfully", repo)
}
