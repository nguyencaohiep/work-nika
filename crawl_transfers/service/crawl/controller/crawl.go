package controller

import (
	"admin_service/pkg/log"
	"fmt"
	"io/ioutil"
	"net/http"
	"score_gear5/pkg/db"

	"github.com/google/uuid"
)

type CryptoScore struct {
	Id        uuid.UUID
	Address   *string
	Transfers int64
}

type CryptoScoreRepo struct {
	CryptoScores []CryptoScore `json:"cryptoScores"`
}

func (repo *CryptoScoreRepo) GetInfo() error {
	query := `select id, address 
			from crypto where type = 'token' and chainid = '1'  and transfers is null order by score  desc nulls last;`

	rows, err := db.PSQL.Query(query)
	if err != nil {
		return err
	}
	defer rows.Close()

	for rows.Next() {
		var cryptoScore = CryptoScore{}
		err := rows.Scan(&cryptoScore.Id, &cryptoScore.Address)
		if err != nil {
			return err
		}
		repo.CryptoScores = append(repo.CryptoScores, cryptoScore)

	}
	return nil
}

func Score() {
	repo := &CryptoScoreRepo{}
	err := repo.GetInfo()

	if err != nil {
		fmt.Println(err)
		return

	}

	CrawlTransfers(*repo.CryptoScores[0].Address)
}

func CrawlTransfers(address string) int64 {
	api := "https://etherscan.io/token/" + address
	clientCoingecko := http.Client{}

	resp, err := clientCoingecko.Get(api)
	if err != nil {
		log.Println(log.LogLevelWarn, "CrawlPriceCoingecko clientCoingecko.Get(coingeckoAPI)", "")
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(log.LogLevelWarn, "Coingecko/CrawlPrices", "Convert data to byte : Error")
	}
	fmt.Println(string(body))
	return 0
}
