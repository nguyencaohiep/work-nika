package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"math/big"
	"test3/work/todo-go"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

var url = "https://goerli.infura.io/v3/faaf285d543d4068bc975e212a497409"

func main() {
	b, err := ioutil.ReadFile("wallets/UTC--2022-09-22T02-31-05.652139062Z--dad66c8b39e2e95da774b6a95b7520dcaf33fdd2")
	if err != nil {
		log.Fatal(err)
	}

	key, err := keystore.DecryptKey(b, "password")
	if err != nil {
		log.Fatal(err)
	}

	client, err := ethclient.Dial(url)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	add := crypto.PubkeyToAddress(key.PrivateKey.PublicKey)
	nonce, err := client.PendingNonceAt(context.Background(), add)
	if err != nil {
		log.Fatal(err)
	}

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	chainId, err := client.NetworkID(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	auth, err := bind.NewKeyedTransactorWithChainID(key.PrivateKey, chainId)
	if err != nil {
		log.Fatal(err)
	}

	auth.GasPrice = gasPrice
	auth.GasLimit = uint64(3000000)
	auth.Nonce = big.NewInt(int64(nonce))
	a, tx, _, err := todo.DeployTodo(auth, client)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(a.Hex())
	fmt.Println(tx.Hash().Hex())

}
