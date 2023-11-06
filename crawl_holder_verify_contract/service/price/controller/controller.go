package controller

import (
	"holder_contract/pkg/ms"
	"holder_contract/pkg/router"
	"holder_contract/pkg/server"
	"holder_contract/service/price/crawler"
	"net/http"
)

var (
	keyMSETH = server.Config.GetString("KEY_MS_ETH")
	keyMSBSC = server.Config.GetString("KEY_MS_BSC")
)

type InfoUpdate struct {
	LastUpdateTime string `json:"LastUpdateTime"`
	Update         int    `json:"Update"`
	Insert         int    `json:"Insert"`
}
type Info struct {
	ChainId  string                     `json:"chainId"`
	LenQueue int64                      `json:"lenQueue"`
	MapInfo  map[string]crawler.InfoEle `json:"mapInfo"`
}

func GetInfoETH(w http.ResponseWriter, r *http.Request) {
	infoUpdate := &InfoUpdate{
		LastUpdateTime: crawler.TimeUpdateETH,
		Insert:         0,
		Update:         len(crawler.MapInfoETH),
	}

	router.ResponseSuccessWithData(w, "B.200", "Get info update ETH successfully!", infoUpdate)
}

func GetInfoBSC(w http.ResponseWriter, r *http.Request) {
	infoUpdate := &InfoUpdate{
		LastUpdateTime: crawler.TimeUpdateBSC,
		Insert:         0,
		Update:         len(crawler.MapInfoBSC),
	}

	router.ResponseSuccessWithData(w, "B.200", "Get info update BSC successfully!", infoUpdate)
}

func GetListInfoETH(w http.ResponseWriter, r *http.Request) {

	info := Info{
		ChainId:  "1",
		LenQueue: ms.Redis.Store.LLen(keyMSETH).Val(),
		MapInfo:  crawler.MapInfoETH,
	}
	router.ResponseSuccessWithData(w, "B.200", "Get info update ETH successfully!", info)
}

func GetListInfoBSC(w http.ResponseWriter, r *http.Request) {
	info := Info{
		ChainId:  "56",
		LenQueue: ms.Redis.Store.LLen(keyMSBSC).Val(),
		MapInfo:  crawler.MapInfoBSC,
	}

	router.ResponseSuccessWithData(w, "B.200", "Get info update BSC successfully!", info)
}

func GetListInfoBSCHanddle(w http.ResponseWriter, r *http.Request) {

	router.ResponseSuccessWithData(w, "B.200", "Get info update BSC successfully!", len(crawler.MapHandleBSC))
}
