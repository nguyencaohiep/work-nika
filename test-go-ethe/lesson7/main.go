package main

import (
	"context"
	"encoding/hex"
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {
	client, err := ethclient.Dial("https://goerli.infura.io/v3/faaf285d543d4068bc975e212a497409")
	if err != nil {
		log.Fatal(err)
	}

	contractAddress := common.HexToAddress("0x0E3AA779d5d30246F632938ae5EaA70a8730D135")
	bytecode, err := client.CodeAt(context.Background(), contractAddress, nil) // nil is latest block
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(hex.EncodeToString(bytecode)) // 60806...10033
}
