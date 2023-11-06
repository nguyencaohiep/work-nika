package crawler

import (
	"native_service/pkg/log"
	"native_service/service/native/model"
	"strings"
)

var (
	MapCoinInfo map[string]CoinStruct
	priceInit   string
	LenMapCoins int64
)

type CoinStruct struct {
	Address string
	Name    string
	Symbol  string
	Price   string
	Source  string
}

const (
	DEFAULT_ADD_COIN = "null"
)

func init() {
	MapCoinInfo = PreparePairAddress()
	LenMapCoins = int64(len(MapCoinInfo))
}

func PreparePairAddress() map[string]CoinStruct {
	mapCoinInfo := map[string]CoinStruct{}

	coinRepo := &model.CoinRepo{}
	err := coinRepo.GetAllCoins(DEFAULT_ADD_COIN)
	if err != nil {
		log.Println(log.LogLevelWarn, "PreparePairAddress clientCoingecko.Get(coingeckoAPI)", "")
	}

	for _, coin := range coinRepo.Coins {
		coinStruct := CoinStruct{
			Name:    coin.Name,
			Symbol:  coin.Symbol,
			Price:   priceInit,
			Address: coin.Address,
		}
		mapCoinInfo[strings.ToLower(coin.Symbol)+"-"+coin.Name] = coinStruct
	}

	return mapCoinInfo
}
