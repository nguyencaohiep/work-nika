package crawler

import (
	"convert-service/pkg/log"
	"convert-service/pkg/server"
	"convert-service/service/convert/dao"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

type CurrencyInfoBinance struct {
	Data []Data `json:"data"`
}

type Data struct {
	Pair     string  `json:"pair"`
	Rate     float64 `json:"rate"`
	Symbol   string  `json:"symbol"`
	FullName string  `json:"fullName"`
	ImageURL string  `json:"imageUrl"`
}

type CryptoInfo struct {
	Name         string  `json:"name"`
	Symbol       string  `json:"symbol"`
	SmallLogo    string  `json:"smallLogo"`
	TotalSupply  string  `json:"totalSupply"`
	MarketcapUSD float64 `json:"marketcapusd"`
	Price        float64 `json:"price"`
}

var (
	RepoCryptos           dao.RepoCryptos
	MapCryptoBinance      map[string]int // key is pair binance, value is index in RepoCryptos
	MapCryptoCGC          map[string]int
	MapIndexCrypto        map[string]int // key is gear5Id, value is index in RepoCryptos
	MapCurrencies         map[string]dao.Currency
	ListCompactInfoCrypto []CryptoInfo
)

func init() {
	GetDes()
	// MapCryptoBinance = map[string]int{}
	// MapCryptoCGC = map[string]int{}
	// MapIndexCrypto = map[string]int{}
	// MapCurrencies = map[string]dao.Currency{}
	// ListCompactInfoCrypto = []CryptoInfo{}
	// prepareCryptoInfo()
	// prepareCurrencyInfo()
	// go func() {
	// 	for {
	// 		GetCurrenciesBinance()
	// 		time.Sleep(2 * time.Second)
	// 	}
	// }()
}

func prepareCryptoInfo() {
	repo := dao.RepoCryptos{}
	err := repo.GetCryptos()
	if err != nil {
		log.Println(log.LogLevelError, "prepareCryptoInfo: repo.GetCryptos()", err.Error())
		return
	}
	RepoCryptos = repo

	for i, crypto := range repo.Cryptos {
		_, exist := MapCryptoBinance[crypto.Symbol+"USDT"]
		if !exist {
			MapCryptoBinance[crypto.Symbol+"USDT"] = i
		}

		_, exist = MapCryptoCGC[crypto.CryptoCode]
		if !exist {
			MapCryptoCGC[crypto.CryptoCode] = i
		}

		cryptoInfo := &CryptoInfo{
			Name:         crypto.Name,
			Symbol:       crypto.Symbol,
			SmallLogo:    crypto.SmallLogo,
			MarketcapUSD: crypto.MarketcapUSD,
			TotalSupply:  crypto.TotalSupply,
		}
		ListCompactInfoCrypto = append(ListCompactInfoCrypto, *cryptoInfo)
		MapIndexCrypto[crypto.Symbol] = i
	}
}

func prepareCurrencyInfo() {
	repo := dao.RepoCurrency{}
	err := repo.GetCurrencies()
	if err != nil {
		log.Println(log.LogLevelError, "prepareCurrencyInfo: repo.GetCurrencies()", err.Error())
		return
	}
	for _, currency := range repo.Currencies {
		MapCurrencies[currency.Symbol] = currency
	}

}

func GetCurrenciesBinance() {
	api := server.Config.GetString("API_CURRENCY_BINANCE")
	clientBinance := http.Client{}

	resp, err := clientBinance.Get(api)
	if err != nil {
		log.Println(log.LogLevelWarn, "GetCurrenciesBinance client.Get(api)", err.Error())
		return
	}

	if resp != nil {
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Println(log.LogLevelWarn, "GetCurrenciesBinance ioutil.ReadAll(resp.Body)", err.Error())
			return
		}
		defer resp.Body.Close()

		currencyInfoBinance := &CurrencyInfoBinance{}
		err = json.Unmarshal(body, &currencyInfoBinance)
		if err != nil {
			log.Println(log.LogLevelWarn, "GetCurrenciesBinance json.Unmarshal(body, &resSol)", err.Error())
			return
		}

		indexUSDT := 2
		crypto := RepoCryptos.Cryptos[indexUSDT]
		currentPriceUSDT := map[string]float64{}
		currentPriceUSDT["USD"] = 1

		compactCrypto := ListCompactInfoCrypto[indexUSDT]
		compactCrypto.Price = 1
		ListCompactInfoCrypto[indexUSDT] = compactCrypto

		for _, currencyInfo := range currencyInfoBinance.Data {
			pairCurrency := strings.Split(currencyInfo.Pair, "_")
			if pairCurrency[1] == "USD" {
				mutex.Lock()
				currency, exist := MapCurrencies[pairCurrency[0]]
				if exist {
					currency.Rate = currencyInfo.Rate
					currentPriceUSDT[pairCurrency[0]] = currencyInfo.Rate
					MapCurrencies[pairCurrency[0]] = currency
				}
				mutex.Unlock()
			}
		}
		crypto.CurrentPrice = currentPriceUSDT
		RepoCryptos.Cryptos[indexUSDT] = crypto
	}
}

func GetDes() {
	repo := &dao.RepoCryptos{}
	err := repo.GetDes()
	if err != nil {
		fmt.Println(err)
	}
}
