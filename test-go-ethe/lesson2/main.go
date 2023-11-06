package main

import (
	"crypto/ecdsa"
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
)

func main() {
	fmt.Println("Generate Wallet!")

	privatekey, err := crypto.GenerateKey() // generate random private key
	if err != nil {
		log.Fatal(err)
	}
	privateKeyBytes := crypto.FromECDSA(privatekey)
	hexPrivateKey := hexutil.Encode(privateKeyBytes) //  convert to hex string
	realPrivatekey := hexPrivateKey[2:]              // cut "0x" in head string to have real privateKey
	fmt.Println("Real private key: ", realPrivatekey)

	publicKey := privatekey.Public()                   // generate public key from private key
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey) // convert to ecdsaKey
	if !ok {
		log.Fatal("error casting public key to ECDSA")
	}
	publicKeyBytes := crypto.FromECDSAPub(publicKeyECDSA) // convert to byte array
	hexPublicKey := hexutil.Encode(publicKeyBytes)
	realPublicKey := hexPublicKey[4:]
	fmt.Println("Real publicKey ", realPublicKey)

	address := crypto.PubkeyToAddress(*publicKeyECDSA).Hex()
	fmt.Println("Address : ", address)
}
