package crawler

import (
	"convert-service/pkg/log"
	"convert-service/pkg/server"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
	"sync"
	"time"

	"golang.org/x/exp/slices"
)

type PriceBinance struct {
	Symbol string `json:"symbol"` // symbol binanace return, ex BTCUSDT
	Price  string `json:"price"`
}

var (
	clientBinance     http.Client
	RepoInfoBinance   []PriceBinance
	ArrayIndexBinance []int
	mutex             sync.Mutex
)

func init() {
	clientBinance = http.Client{}
	ArrayIndexBinance = []int{}
}

func init() {
	GetTopPriceInfoBinance()
}

func GetTopPriceInfoBinance() {
	api := server.Config.GetString("API_BINANCE")
	resp, err := clientBinance.Get(api)
	if err != nil {
		log.Println(log.LogLevelWarn, "GetTopPriceInfo client.Get(api)", err.Error())
		time.Sleep(3 * time.Minute)
		return
	}

	if resp != nil {
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Println(log.LogLevelWarn, "GetTopPriceInfo ioutil.ReadAll(resp.Body)", err.Error())
			time.Sleep(3 * time.Minute)
			return
		}
		defer resp.Body.Close()

		err = json.Unmarshal(body, &RepoInfoBinance)
		if err != nil {
			log.Println(log.LogLevelWarn, "GetTopPriceInfo json.Unmarshal(body, &resSol)", err.Error())
			time.Sleep(3 * time.Minute)
			return
		}

		// fmt.Println(RepoInfoBinance)
		for _, priceInfo := range RepoInfoBinance {
			index, exist := MapCryptoBinance[priceInfo.Symbol]
			if exist {
				if !slices.Contains(ArrayIndexBinance, index) {
					ArrayIndexBinance = append(ArrayIndexBinance, index)
				}

				crypto := RepoCryptos.Cryptos[index]
				if len(crypto.CurrentPrice) == 0 {
					crypto.CurrentPrice = map[string]float64{}
				}

				priceFloat, err := strconv.ParseFloat(priceInfo.Price, 64)
				if err == nil {
					crypto.CurrentPrice["USD"] = priceFloat
					mutex.Lock()
					for key, currency := range MapCurrencies {
						if currency.Rate > 0 {
							crypto.CurrentPrice[key] = currency.Rate * priceFloat
						}
					}
					mutex.Unlock()
				}

				RepoCryptos.Cryptos[index] = crypto

				compactCrypto := ListCompactInfoCrypto[index]
				compactCrypto.Price = crypto.CurrentPrice["USD"]
				ListCompactInfoCrypto[index] = compactCrypto
			}
		}
	}
}
