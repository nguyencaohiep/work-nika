package crawler

import (
	"bytes"
	"crawl_price_3rd/pkg/server"
	"crawl_price_3rd/service/price/dao"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"sync"
)

var (
	MapCryptocodeCGC map[string]dao.Crypto
	MapCryptocodeCMC map[string]string
	MapCryptocodeSol map[string]string
	MapPriceBinance  map[string]string
	MapCryptoCode    map[string]string // to use link from MapPriceBinance to MapCryptocodeCGC to get crypto info
	NumberUpdateCMC  int
	NumberUpdateSol  int
	mutex            sync.Mutex
	LenCryptosCMC    int
	LenCryptosSol    int
)

var (
	SrcCGC = "coingecko"
	SrcCMC = "coinmarketcap"
	SrcSOL = "solscan"
	SrcBNB = "binance"
)

type InfoUpdate struct {
	LastUpdateTime string `json:"LastUpdateTime"`
	Update         int    `json:"Update"`
	Insert         int    `json:"Insert"`
}

type DataInfo struct {
	Status  bool   `json:"status"`
	Code    string `json:"code"`
	Message string `json:"message"`
	Data    Data   `json:"data"`
}
type Data struct {
	Infos []Info `json:"infos"`
}
type Info struct {
	Cryptoid   string `json:"cryptoid"`
	Cryptosrc  string `json:"cryptosrc"`
	Cryptocode string `json:"cryptocode"`
	Symbol     string `json:"symbol"`
	Name       string `json:"name"`
}

func init() {
	MapCryptocodeCGC = map[string]dao.Crypto{}
	MapCryptocodeCMC = map[string]string{}
	MapCryptocodeSol = map[string]string{}
	MapPriceBinance = map[string]string{}
	prepareInfo()
	LenCryptosSol = len(MapCryptocodeSol)
	LenCryptosCMC = len(MapCryptocodeCMC)
}

func prepareInfo() error {
	api := server.Config.GetString("DOMAIN_LOCAL") + server.Config.GetString("API_GET_INFO")
	req, err := http.NewRequest(http.MethodGet, api, nil)
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	client := http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}

	infoData := &DataInfo{}
	err = json.Unmarshal(body, &infoData)
	if err != nil {
		return err
	}

	numberTop := 0

	for _, info := range infoData.Data.Infos {
		crypto := dao.Crypto{
			Cryptocode: info.Cryptocode,
			CryptoId:   info.Cryptoid,
			CryptoSrc:  info.Cryptosrc,
			Symbol:     info.Symbol,
			Name:       info.Name,
		}
		if info.Cryptosrc == SrcCGC {
			_, exist := MapCryptocodeCGC[info.Cryptocode]
			if !exist {
				MapCryptocodeCGC[info.Cryptocode] = crypto
				if numberTop < 100 {
					numberTop++
					_, exist := MapPriceBinance[crypto.Symbol+"USDT"]
					if !exist {
						MapPriceBinance[crypto.Symbol+"USDT"] = crypto.Cryptocode
					}
				}
			}
		} else if info.Cryptosrc == SrcCMC {
			_, exist := MapCryptocodeCMC[crypto.Symbol]
			if !exist {
				MapCryptocodeCMC[crypto.Symbol] = crypto.CryptoId
			}
		} else if info.Cryptosrc == SrcSOL {
			_, exist := MapCryptocodeSol[crypto.Symbol]
			if !exist {
				MapCryptocodeSol[crypto.Symbol] = crypto.CryptoId
			}
		}
	}
	return nil
}

func UpdatePrice(repo dao.ListCrypto) error {
	// return nil
	jsonBody, err := json.Marshal(repo)
	if err != nil {
		return err
	}
	bodyReader := bytes.NewReader(jsonBody)

	api := server.Config.GetString("DOMAIN_LOCAL") + server.Config.GetString("API_UPDATE_PRICE")
	req, err := http.NewRequest(http.MethodPatch, api, bodyReader)
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}

	resBody, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if res.StatusCode > 299 || res.StatusCode < 200 {
		return errors.New(res.Status + " " + string(resBody))
	}
	return nil
}
