package crawler

import (
	"encoding/json"
	"holder_contract/pkg/log"
	"holder_contract/pkg/ms"
	"holder_contract/pkg/server"
	"holder_contract/pkg/utils"
	"holder_contract/service/price/dao"
	"net/http"
	"time"
)

var (
	bscClient         http.Client
	PreLatestBlockBSC int64
	keyMSBSC          string
	NumberUpdateETH   int
	TimeUpdateETH     string
	LenQueueBSC       int64
	MapInfoBSC        map[string]InfoEle
	MapHandleBSC      map[string]Token
)

type Token struct {
	ChainId  string `json:"chainId"`
	CryptoId string `json:"cryptoId"`
	Address  string `json:"address"`
}

func init() {
	PreLatestBlockBSC = 0
	bscClient = http.Client{}
	keyMSBSC = server.Config.GetString("KEY_MS_BSC")
	LenQueueBSC = 0
	MapInfoBSC = map[string]InfoEle{}
}

func ListenQueueBSC() {
	repo := dao.CryptoRepo{}
	LenQueueBSC = ms.Redis.Store.LLen(keyMSBSC).Val()
	// fmt.Println("len queue bsc", LenQueueBSC)
	MapHandleBSC := map[string]Token{}

	for i := 0; i < int(LenQueueBSC); i++ {
		result, err := ms.Redis.Store.BLPop(1*time.Second, keyMSBSC).Result() // get and remove first element with timeout 1 second
		if err != nil {
			log.Println(log.LogLevelError, "ListenQueueBSC: ms.Redis.Store.BLPop", err)
			continue
		}
		tmp := []byte(result[1])

		token := &Token{}
		err = json.Unmarshal(tmp, &token)
		if err != nil {
			log.Println(log.LogLevelError, "ListenQueueBSC: json.Unmarshal(tmp, &token)", err)
			continue
		}
		MapHandleBSC[token.Address] = *token
	}

	// fmt.Println("len map", len(MapHandleBSC))
	for _, token := range MapHandleBSC {
		holders, err := CrawlHoldersBSC(token.Address)
		// log.Println(log.LogLevelInfo, "bsc "+token.Address, holders)
		if err != nil {
			log.Println(log.LogLevelError, "ListenQueueBSC: CrawlHolders", err)
		}

		updateCrypto := &dao.Crypto{
			Holders:  holders,
			CryptoId: token.CryptoId,
		}

		verify, err := checkContractVerifiedDB(token.CryptoId)
		if err != nil {
			log.Println(log.LogLevelError, "ListenQueueBSC: checkContractVerifiedDB", err)
		}

		if !verify {
			verify, err = checkContractBSC(token.Address)
			if err != nil {
				log.Println(log.LogLevelError, "ListenQueueBSC: checkContractBSC", err)
			}
		}
		updateCrypto.Contractverified = verify
		repo.Cryptos = append(repo.Cryptos, *updateCrypto)
		// }()
	}
	err := updateHolderContract(repo)
	if err != nil {
		log.Println(log.LogLevelError, "ListenQueueBSC: updateHolderContract(repo)", err)
	}

	for _, ele := range repo.Cryptos {
		infoEle := InfoEle{
			Time: utils.TimeNowString(),
		}

		oleEle, exist := MapInfoBSC[ele.CryptoId]
		if exist && ele.Holders == 0 {
			infoEle.Holder = oleEle.Holder
		} else {
			infoEle.Holder = ele.Holders
		}
		MapInfoBSC[ele.CryptoId] = infoEle
	}

	NumberUpdateBSC = len(repo.Cryptos)
	TimeUpdateBSC = utils.TimeNowString()
}
