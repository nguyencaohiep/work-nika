package crawler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"native_service/pkg/log"
	"native_service/pkg/utils"
	"native_service/service/native/constant"
	"native_service/service/native/model"
	"net/http"
	"time"
)

type CoingeckoList struct {
	ListPrice []PriceCoingeckoStr
}

type PriceCoingeckoStr struct {
	Symbol       string  `json:"symbol"`
	Name         string  `json:"name"`
	CurrentPrice float64 `json:"current_price"`
}

func CrawlPriceCoingecko() {
	page := 1
	foundedNumberCoin := int64(0)
	for lenListPerPage := -1; lenListPerPage != 0 && page <= 250 && foundedNumberCoin < LenMapCoins; page++ {
		// fmt.Println(page)
		coingeckoAPI := fmt.Sprintf(constant.API_PRICE_COINGECKO, constant.MAX_RECORDS_PER_PAGE, page)
		clientCoingecko := http.Client{
			Timeout: 10 * time.Second,
		}

		resp, err := clientCoingecko.Get(coingeckoAPI)
		if err != nil {
			log.Println(log.LogLevelWarn, "CrawlPriceCoingecko clientCoingecko.Get(coingeckoAPI)", "")
		}

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Println(log.LogLevelWarn, "Coingecko/CrawlPrices", "Convert data to byte : Error")
		}

		rawCoingeckoDTOArr := make([]any, 0)
		err = json.Unmarshal(body, &rawCoingeckoDTOArr)
		if err != nil {
			log.Println(log.LogLevelWarn, "CrawlPriceCoingecko Unmarshal(body, &rawCoingeckoDTOArr)", "")
		}
		lenListPerPage = len(rawCoingeckoDTOArr)

		// Traverse each json object from response array data got above.
		for _, rawCoingeckoDTO := range rawCoingeckoDTOArr {
			coingeckoDTO := &PriceCoingeckoStr{}
			err = utils.Mapping(rawCoingeckoDTO, coingeckoDTO)
			if err != nil {
				log.Println(log.LogLevelWarn, "CrawlPriceCoingecko utils.Mapping(rawCoingeckoDTO, coingeckoDTO)", "")
				continue
			}

			coinStruct, exist := MapCoinInfo[coingeckoDTO.Name+"-"+coingeckoDTO.Symbol]
			if exist {
				coinStruct.Price = fmt.Sprint(coingeckoDTO.CurrentPrice)
				coinStruct.Source = "coingecko"
				MapCoinInfo[coingeckoDTO.Name+"-"+coingeckoDTO.Symbol] = coinStruct
				foundedNumberCoin++
			}
		}
	}
}

func CrawlPriceNative() {
	for {
		CrawlPriceCoingecko()
		for _, coinPrice := range MapCoinInfo {
			coin := &model.Coin{
				Name:    coinPrice.Name,
				Symbol:  coinPrice.Symbol,
				Source:  coinPrice.Source,
				Address: coinPrice.Address,
			}
			fmt.Println(coinPrice.Symbol)
			err := coin.InsertPrice()
			if err != nil {
				log.Println(log.LogLevelError, "CrawlPriceNative : insert price", err)
			}
		}
		time.Sleep(5 * time.Minute)
	}
}
