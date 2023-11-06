package controller

import (
	"context"
	"encoding/hex"
	"explore_address/pkg/log"
	"explore_address/pkg/server"
	"explore_address/service/address/constant"
	"explore_address/service/address/model/dao"
	"fmt"
	"math"
	"math/big"
	"strconv"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

var (
	ethClient         *ethclient.Client
	listNodeRPCETH    []string
	PreLatestBlockETH int64
	MapDecimalETH     map[string]int64
	MapAmountETH      map[string]float64
)

func init() {
	listNodeRPCETH = strings.Fields(server.Config.GetString("LIST_NODES_ETH"))
	MapDecimalETH = map[string]int64{}
	MapAmountETH = map[string]float64{}
}

func ListenTransfersETH() {
	var latestBlock uint64
	ethClient, latestBlock = ChooseClient(listNodeRPCETH, ethClient)
	latestBlockInt64 := int64(latestBlock)

	if PreLatestBlockETH == 0 {
		PreLatestBlockETH = latestBlockInt64
	}

	if ethClient != nil {
		for i := PreLatestBlockETH; i <= latestBlockInt64; i++ {
			// mapAmount := map[string]float64{}
			query := ethereum.FilterQuery{
				FromBlock: big.NewInt(i),
				ToBlock:   big.NewInt(i),
				Topics:    [][]common.Hash{{topicTransferHash}},
			}

			logs, err := ethClient.FilterLogs(context.Background(), query)
			fmt.Println("len log", len(logs))
			if err != nil {
				log.Println(log.LogLevelError, "ListenTransfersETH: FilterLogs()", err.Error())
				continue
			}

			for _, logData := range logs {
				if !logData.Removed && len(logData.Topics) > 0 {
					if hex.EncodeToString(logData.Topics[0].Bytes()) == eventSignTransfers {
						tokenAddress := strings.ToLower(logData.Address.Hex())
						fromAddress := "0x" + strings.ToLower(hex.EncodeToString(logData.Topics[1].Bytes())[24:])
						toAddress := "0x" + strings.ToLower(hex.EncodeToString(logData.Topics[2].Bytes())[24:])

						var decimal int64
						decimal, exist := MapDecimalETH[tokenAddress]

						if !exist { // get decimal
							crypto := &dao.Address{
								ChainName: constant.CHAIN_NAME_ETH,
								Address:   toAddress,
							}
							err = crypto.GetDecimal()
							if err != nil {
								log.Println(log.LogLevelError, "ListenTransfersETH: crypto.GetDecimal()", err.Error())
								continue
							} else {
								decimal = crypto.Decimal
							}
						}

						amount, err := convertToAmountFloat(logData.Data, decimal)
						if err != nil {
							log.Println(log.LogLevelError, "ListenTransfersETH: convertToAmountFloat()", err.Error())
							continue
						}

						if fromAddress != constant.MINT_ADD {
							currentAmount, exist := MapAmountETH[fromAddress+"-"+tokenAddress]
							if !exist {
								address := &dao.Address{
									Address:   fromAddress,
									ChainName: constant.CHAIN_NAME_ETH,
								}

							}
						}
					}
				}
			}
		}

		PreLatestBlockETH = latestBlockInt64 + 1

	} else {
		time.Sleep(10 * time.Second)
	}
}

func convertToAmountFloat(logData []byte, decimal int64) (float64, error) {
	amountFloat := float64(0)
	data := make([]interface{}, lenNonIndexed)
	err := contractABI.UnpackIntoInterface(&data, eventNameTransfer, logData)
	if err != nil {
		log.Println(log.LogLevelError, "convertToAmountFloat: UnpackIntoInterface ", err.Error())
		return amountFloat, err
	}

	amount := fmt.Sprint(data[0])

	amountFloat, err = strconv.ParseFloat(amount, 64)
	if err != nil {
		log.Println(log.LogLevelError, "convertToAmountFloat: strconv.ParseFloat(amount, 64) ", err.Error())
		return amountFloat, err
	}
	return (amountFloat / math.Pow(10, float64(decimal))), nil
}
