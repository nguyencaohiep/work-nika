package controller

import (
	"context"
	"explore_address/pkg/log"
	"explore_address/pkg/server"
	"math/big"

	"strings"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

var (
	topicTransferHash  common.Hash
	abiTransfers       string
	eventSignTransfers string
	eventNameTransfer  string
	contractABI        *abi.ABI
	lenNonIndexed      int
)

func init() {
	abiTransfers = server.Config.GetString("CONTRACT_ABI")
	contractABI = GetContractABI()
	eventSignTransfers = server.Config.GetString("EVENT_SIGN_TRANSFER")
	topicTransferHash = common.HexToHash("0x" + eventSignTransfers)
	eventNameTransfer = server.Config.GetString("EVENT_NAME_TRANSFER")
	lenNonIndexed = len(contractABI.Events[eventNameTransfer].Inputs.NonIndexed())
}

func GetContractABI() *abi.ABI {
	contractABI, err := abi.JSON(strings.NewReader(abiTransfers))
	if err != nil {
		log.Println(log.LogLevelError, "GetContractABI: abi.JSON(strings.NewReader(abiTransfers))", err.Error())
		return nil
	}
	return &contractABI
}

func ChooseClient(listNodeRPC []string, currentClient *ethclient.Client) (*ethclient.Client, uint64) {
	var latestBlock uint64
	var client *ethclient.Client
	if currentClient != nil {
		var err error
		latestBlock, err = currentClient.BlockNumber(context.Background())
		if err != nil {
			log.Println(log.LogLevelError, "currentClient", err.Error())
		} else {
			query := ethereum.FilterQuery{
				FromBlock: big.NewInt(int64(latestBlock)),
				ToBlock:   big.NewInt(int64(latestBlock)),
				Topics:    [][]common.Hash{{topicTransferHash}},
			}

			_, err = currentClient.FilterLogs(context.Background(), query)
			if err != nil {
				log.Println(log.LogLevelError, "ChooseClient: FilterLogs", err.Error())
			} else {
				client = currentClient
				return client, latestBlock
			}
		}
	}

	for _, link := range listNodeRPC {
		clientSub, err := ethclient.Dial(link)
		if err != nil {
			log.Println(log.LogLevelInfo, "ChooseClient: ethclient.Dial(link) "+link, err.Error())
			continue
		} else {
			latestBlock, err = clientSub.BlockNumber(context.Background())
			if err != nil {
				log.Println(log.LogLevelInfo, "ChooseClient: clientSub.BlockNumber(context.Background()) "+link, err.Error())
				continue
			}

			query := ethereum.FilterQuery{
				FromBlock: big.NewInt(int64(latestBlock)),
				ToBlock:   big.NewInt(int64(latestBlock)),
				Topics:    [][]common.Hash{{topicTransferHash}},
			}

			_, err = clientSub.FilterLogs(context.Background(), query)
			if err != nil {
				log.Println(log.LogLevelError, "ChooseClient: FilterLogs ", err.Error())
				continue
			} else {
				client = clientSub
				break
			}
		}
	}
	return client, latestBlock
}
