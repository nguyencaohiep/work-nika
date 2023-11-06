package main

import (
	"log"

	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/ethclient"
)

var (
	url = "https://goerli.infura.io/v3/faaf285d543d4068bc975e212a497409"
)

func main() {
	ks := keystore.NewKeyStore("../wallets", keystore.StandardScryptN, keystore.StandardScryptP)

	password := "password"
	_, err := ks.NewAccount(password)
	if err != nil {
		log.Fatal(err)
	}
	_, err = ks.NewAccount(password)
	if err != nil {
		log.Fatal(err)
	}

	client, err := ethclient.Dial(url)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()
}
