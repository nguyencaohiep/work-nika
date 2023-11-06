package main

import (
	"context"
	"math/big"

	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/ethclient"
)

const (
	URLINFURA = "wss://mainnet.infura.io/ws/v3/faaf285d543d4068bc975e212a497409"

	BLOCK_NUMBER = 15643489
)

func main() {

	blockNumber := big.NewInt(BLOCK_NUMBER)
	client, err := ethclient.Dial(URLINFURA)
	if err != nil {
		log.Fatal(err)
	}
	header, err := client.HeaderByNumber(context.Background(), blockNumber)
	if err != nil {
		log.Fatal(err)
	}

	block, err := client.BlockByNumber(context.Background(), blockNumber)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("header number : ", header.Number.String())
	fmt.Println("Header time : ", block.Header().Time)
	fmt.Println("Hash : ", block.Hash())
	fmt.Println("SanityCheck: ", block.SanityCheck())
	fmt.Println("stateroot : ", block.Header().)
	fmt.Println("stateroot : ", block.Transactions())
}
