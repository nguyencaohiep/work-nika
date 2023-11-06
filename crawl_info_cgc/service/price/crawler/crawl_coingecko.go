package crawler

import (
	"crawl_price_3rd/pkg/log"
	"crawl_price_3rd/pkg/utils"
	"crawl_price_3rd/service/price/dao"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

type CoinGeckoMarketInfo struct {
	ID                           string  `json:"id"`
	CurrentPrice                 float64 `json:"current_price"`
	MarketCap                    float64 `json:"market_cap"`
	TotalVolume                  float64 `json:"total_volume"`
	High24h                      float64 `json:"high_24h"`
	Low24h                       float64 `json:"low_24h"`
	PriceChange24h               float64 `json:"price_change_24h"`
	PriceChangePercentage24h     float64 `json:"price_change_percentage_24h"`
	MarketcapChange24h           float64 `json:"market_cap_change_24h"`
	MarketcapChangePercentage24h float64 `json:"market_cap_change_percentage_24h"`
	TotalSupply                  float64 `json:"total_supply"`
	ATH                          float64 `json:"ath"`
	ATHChangePercent             float64 `json:"ath_change_percentage"`
	ATHDate                      string  `json:"ath_date"`
	ATL                          float64 `json:"atl"`
	ATLChangePercentage          float64 `json:"atl_change_percentage"`
	ATLDate                      string  `json:"atl_date"`
}

var clientCoingecko http.Client

func init() {
	clientCoingecko = http.Client{}
}

func CrawlPriceCoingecko() {
	page := 1

	for lenListPerPage := -1; lenListPerPage != 0 && page <= 50; page++ {
		repo := dao.ListCrypto{
			CryptoSrc: SrcCGC,
		}
		// fmt.Println("page cgc", page)
		api := fmt.Sprintf(`https://api.coingecko.com/api/v3/coins/markets?vs_currency=usd&per_page=%v&page=%v`, 250, page)

		resp, err := clientCoingecko.Get(api)
		if err != nil {
			log.Println(log.LogLevelWarn, "CrawlPriceCoingecko clientCoingecko.Get(coingeckoAPI)", err.Error())
			time.Sleep(3 * time.Minute)
			return
		}

		if resp != nil {
			body, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				log.Println(log.LogLevelWarn, "Coingecko/CrawlPrices", err.Error())
				time.Sleep(3 * time.Minute)
				return
			}
			defer resp.Body.Close()

			rawCoingeckoDTOArr := make([]any, 0)
			err = json.Unmarshal(body, &rawCoingeckoDTOArr)
			if err != nil {
				log.Println(log.LogLevelWarn, "CrawlPriceCoingecko Unmarshal(body, &rawCoingeckoDTOArr)"+string(body), err.Error())
				time.Sleep(3 * time.Minute)
				return
			}
			// fmt.Println(rawCoingeckoDTOArr...)
			lenListPerPage = len(rawCoingeckoDTOArr)

			// Traverse each json object from response array data got above.
			for _, rawCoingeckoDTO := range rawCoingeckoDTOArr {
				coinGeckoMarketInfo := &CoinGeckoMarketInfo{}
				err = utils.Mapping(rawCoingeckoDTO, coinGeckoMarketInfo)
				if err != nil {
					log.Println(log.LogLevelWarn, "CrawlPriceCoingecko utils.Mapping(rawCoingeckoDTO, coingeckoDTO)", err.Error())
					continue
				}

				mutex.Lock()
				crypto, exist := MapCryptocodeCGC[coinGeckoMarketInfo.ID]

				if exist {
					crypto := &dao.Crypto{
						CryptoId:                     crypto.CryptoId,
						Name:                         crypto.Name,
						Symbol:                       crypto.Symbol,
						MarketcapUSD:                 coinGeckoMarketInfo.MarketCap,
						TotalSupply:                  fmt.Sprintf("%v", coinGeckoMarketInfo.TotalSupply),
						PriceUSD:                     coinGeckoMarketInfo.CurrentPrice,
						TotalVolume:                  coinGeckoMarketInfo.TotalVolume,
						High24h:                      coinGeckoMarketInfo.High24h,
						Low24h:                       coinGeckoMarketInfo.Low24h,
						PriceChange24h:               coinGeckoMarketInfo.PriceChange24h,
						PriceChangePercentage24h:     coinGeckoMarketInfo.PriceChangePercentage24h,
						MarketcapChange24h:           coinGeckoMarketInfo.MarketcapChange24h,
						MarketcapChangePercentage24h: coinGeckoMarketInfo.MarketcapChangePercentage24h,
						ATH:                          coinGeckoMarketInfo.ATH,
						ATHChangePercent:             coinGeckoMarketInfo.ATHChangePercent,
						ATHDate:                      coinGeckoMarketInfo.ATHDate,
						ATL:                          coinGeckoMarketInfo.ATL,
						ATLChangePercentage:          coinGeckoMarketInfo.ATHChangePercent,
						ATLDate:                      coinGeckoMarketInfo.ATLDate,
					}
					MapCryptocodeCGC[coinGeckoMarketInfo.ID] = *crypto
					repo.Cryptos = append(repo.Cryptos, *crypto)
				}

				// fmt.Println(crypto)

				mutex.Unlock()
			}
		}

		err = UpdatePrice(repo)
		if err != nil {
			log.Println(log.LogLevelError, "CrawlPriceCoingecko: UpdatePrice", err.Error())
		}
		// fmt.Println("update done")

		if page%2 == 0 {
			time.Sleep(1 * time.Minute)
		}
	}
}
