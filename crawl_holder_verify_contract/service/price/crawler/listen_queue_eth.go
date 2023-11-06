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

	"golang.org/x/exp/slices"
)

var (
	keyMSETH        string
	ethClient       http.Client
	NumberUpdateBSC int
	TimeUpdateBSC   string
	MapInfoETH      map[string]InfoEle
	LenQueueETH     int64
)

type InfoEle struct {
	Holder int64  `json:"holder"`
	Time   string `json:"time"`
}

type ResCheckAPI struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	Result  []struct {
		SourceCode string `json:"SourceCode"`
		Abi        string `json:"ABI"`
	} `json:"result"`
}

func init() {
	ethClient = http.Client{}
	keyMSETH = server.Config.GetString("KEY_MS_ETH")
	MapInfoETH = map[string]InfoEle{}
}

func ListenQueueETH() {
	arrayAddressETH := []string{}
	repo := dao.CryptoRepo{}
	LenQueueETH = ms.Redis.Store.LLen(keyMSETH).Val()
	// fmt.Println("len queue eth: ", LenQueueETH)
	for i := 0; i < int(LenQueueETH); i++ {
		result, err := ms.Redis.Store.BLPop(1*time.Second, keyMSETH).Result() // get and remove first element with timeout 1 second
		if err != nil {
			log.Println(log.LogLevelError, "ListenQueueETH: ms.Redis.Store.BLPop", err)
			continue
		}
		tmp := []byte(result[1])

		token := &Token{}
		err = json.Unmarshal(tmp, &token)
		if err != nil {
			log.Println(log.LogLevelError, "ListenQueueETH: json.Unmarshal(tmp, &token)", err)
			continue
		}

		if !slices.Contains(arrayAddressETH, token.Address) {
			arrayAddressETH = append(arrayAddressETH, token.Address)
			holders, err := CrawlHoldersETH(token.Address)
			// log.Println(log.LogLevelInfo, "eth "+token.Address, holders)
			if err != nil {
				log.Println(log.LogLevelError, "ListenQueueETH: CrawlHolders", err)
				continue
			}

			updateCrypto := &dao.Crypto{
				Holders:  holders,
				CryptoId: token.CryptoId,
			}

			verify, err := checkContractVerifiedDB(token.CryptoId) // check contractVerified in db
			if err != nil {
				log.Println(log.LogLevelError, "ListenQueueETH: checkContractVerifiedDB", err)
			}

			if !verify {
				verify, err = checkContractETH(token.Address)
				if err != nil {
					log.Println(log.LogLevelError, "ListenQueueETH: checkContractETH", err)
				}
			}
			updateCrypto.Contractverified = verify
			repo.Cryptos = append(repo.Cryptos, *updateCrypto)
		}
	}

	err := updateHolderContract(repo)
	if err != nil {
		log.Println(log.LogLevelError, "ListenQueueETH: updateHolderContract(repo)", err)
		return
	}
	for _, ele := range repo.Cryptos {
		infoEle := InfoEle{
			Time:   utils.TimeNowString(),
			Holder: ele.Holders,
		}
		MapInfoETH[ele.CryptoId] = infoEle
	}
	NumberUpdateETH = len(repo.Cryptos)
	TimeUpdateETH = utils.TimeNowString()
}
