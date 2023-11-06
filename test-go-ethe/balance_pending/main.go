package main

import (
	"context"
	"fmt"
	"log"
	"math"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

var infuraURl = "https://mainnet.infura.io/v3/faaf285d543d4068bc975e212a497409"

func main() {
	client, err := ethclient.DialContext(context.Background(), infuraURl)
	if err != nil {
		log.Fatal("Error creating client :v", err)
	}
	defer client.Close()

	blockNumber := big.NewInt(5671744)
	block, err := client.BlockByNumber(context.Background(), blockNumber)
	if err != nil {
		log.Fatal("Error to get a block :", err)
	}
	fmt.Println(block.Number())

	fmt.Println("Get BalanceAt and PendingBalance")
	add := "0xF977814e90dA44bFA03b6295A0616a897441aceC"
	address := common.HexToAddress(add)
	fmt.Println(address)                                                 // convert address to common.Address to inteact on go_ethereum
	balance, err := client.BalanceAt(context.Background(), address, nil) //  balance in wei
	if err != nil {
		log.Fatal("Error to get a balance :", err)
	}

	// conver to big float
	fBalance := new(big.Float)
	fBalance.SetString(balance.String())
	value := new(big.Float).Quo(fBalance, big.NewFloat(math.Pow10(18))) // division for 10^18 to amount ether
	fmt.Println(value)

	// pending balance
	pendingBalance, err := client.PendingBalanceAt(context.Background(), address)
	if err != nil {
		log.Fatal("Error to get a pending balance :", err)
	}
	fPendingBalance := new(big.Float)
	fPendingBalance.SetString(pendingBalance.String())
	valuePending := new(big.Float).Quo(fPendingBalance, big.NewFloat(math.Pow10(18)))
	fmt.Println(valuePending)
}
