package lesson4

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

var (
	url = "https://goerli.infura.io/v3/faaf285d543d4068bc975e212a497409"
)

func main() {
	// fmt.Print(1)
	// ks := keystore.NewKeyStore("../wallets", keystore.StandardScryptN, keystore.StandardScryptP)

	// password := "password"
	// a, err := ks.NewAccount(password)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Print(a.Address)
	// // _, err = ks.NewAccount(password)
	// // if err != nil {
	// // 	log.Fatal(err)
	// // }

	// // client, err := ethclient.Dial(url)
	// // if err != nil {
	// // 	log.Fatal(err)
	// // }
	// // defer client.Close()

	client, err := ethclient.Dial(url)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()
	add1 := common.HexToAddress("dad66c8b39e2e95da774b6a95b7520dcaf33fdd2")
	add2 := common.HexToAddress("E2AB4A7b2dd3d1a2E9DBB2d7Ef2CDf3a3C9D8496")

	b1, err := client.BalanceAt(context.Background(), add1, nil)
	if err != nil {
		log.Fatal(err)
	}

	b2, err := client.BalanceAt(context.Background(), add2, nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(b1, b2)

	nonce, err := client.PendingNonceAt(context.Background(), add1)
	if err != nil {
		log.Fatal(err)
	}

	amount := big.NewInt(int64(math.Pow10(17)))
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	tx := types.NewTransaction(nonce, add2, amount, 21000, gasPrice, nil)
	chainId, err := client.NetworkID(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	b, err := ioutil.ReadFile("../wallets/UTC--2022-09-22T02-31-05.652139062Z--dad66c8b39e2e95da774b6a95b7520dcaf33fdd2")
	if err != nil {
		log.Fatal(err)
	}

	key, err := keystore.DecryptKey(b, "password")
	if err != nil {
		log.Fatal(err)
	}

	tx, err = types.SignTx(tx, types.NewEIP155Signer(chainId), key.PrivateKey)
	if err != nil {
		log.Fatal(err)
	}

	err = client.SendTransaction(context.Background(), tx)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("tx sent")

	b1, err = client.BalanceAt(context.Background(), add1, nil)
	if err != nil {
		log.Fatal(err)
	}

	b2, err = client.BalanceAt(context.Background(), add2, nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(b1, b2)
}
