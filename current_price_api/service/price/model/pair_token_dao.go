package model

import (
	"price_service/pkg/db"
)

type PairRepo struct {
	PairAddresses []PairToken `json:"pairAddresses"`
}

type PairToken struct {
	Id           string `json:"id"`
	Source       string `json:"source"`
	Defi         string `json:"defi"`
	ChainId      int64  `json:"chainId"`
	PaireAddress string `json:"paireAddress"`
	Token0       string `json:"token0"`
	Token1       string `json:"token1"`
}

func (repo *PairRepo) GetAllPairTokens(chainId string, defi string) error {
	query := `SELECT pairaddress, token0, token1 FROM pair_token where chainid = $1 and defi = $2;`
	rows, err := db.PSQL.Query(query, chainId, defi)
	if err != nil {
		return err
	}
	defer rows.Close()
	for rows.Next() {
		var pairToken = &PairToken{}
		err := rows.Scan(&pairToken.PaireAddress, &pairToken.Token0, &pairToken.Token1)
		if err != nil {
			return err
		}
		repo.PairAddresses = append(repo.PairAddresses, *pairToken)
	}
	return nil
}
