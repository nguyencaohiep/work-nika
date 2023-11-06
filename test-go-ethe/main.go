package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"test3/work/todo-go"

	"github.com/ethereum/go-ethereum/common"

	"github.com/ethereum/go-ethereum/ethclient"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/keystore"
)

func main() {
	b, err := ioutil.ReadFile("wallets/UTC--2022-09-22T02-31-05.652139062Z--dad66c8b39e2e95da774b6a95b7520dcaf33fdd2")
	if err != nil {
		log.Fatal(err)
	}

	key, err := keystore.DecryptKey(b, "password")
	if err != nil {
		log.Fatal(err)
	}

	client, err := ethclient.Dial("https://goerli.infura.io/v3/faaf285d543d4068bc975e212a497409")
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	chainId, err := client.NetworkID(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	cAdd := common.HexToAddress("0x0E3AA779d5d30246F632938ae5EaA70a8730D135") // address smart contract
	smartContract, err := todo.NewTodo(cAdd, client)                          // get address of smart contract and return a smart contract instance
	if err != nil {
		log.Fatal(err)
	}

	tx, err := bind.NewKeyedTransactorWithChainID(key.PrivateKey, chainId)
	if err != nil {
		log.Fatal(err)
	}

	tx.GasPrice = gasPrice
	tx.GasLimit = 3000000

	tra, err := smartContract.Add(tx, "First task")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(tra.Hash())
}
