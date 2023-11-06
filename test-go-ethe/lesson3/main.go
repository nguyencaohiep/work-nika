package main

import (
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/accounts/keystore"
)

func main() {
	ks := keystore.NewKeyStore("../wallets", keystore.StandardScryptN, keystore.StandardScryptP)
	password := "password"
	account, err := ks.NewAccount(password)
	if err != nil {
		log.Fatal(err)	
	}
	fmt.Println(account.Address.Hex())

	// b, err := ioutil.ReadFile("wallets/UTC--2022-09-21T07-15-58.572294500Z--ec1c9a269e6e23902da12027a26b332c45304eb7")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// key, err := keystore.DecryptKey(b, password)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// privateData := crypto.FromECDSA(key.PrivateKey)

	// publicData := crypto.FromECDSAPub(&key.PrivateKey.PublicKey)
	// address := crypto.PubkeyToAddress(key.PrivateKey.PublicKey)
	// fmt.Println(hexutil.Encode(publicData))
	// fmt.Println(hexutil.Encode(privateData))
	// fmt.Println(address.Hex())
}
