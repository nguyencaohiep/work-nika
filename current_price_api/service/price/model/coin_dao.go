package model

import (
	"price_service/pkg/db"
)

type CoinRepo struct {
	Coins []Coin
}

type Coin struct {
	ChainId  string `json:"chain"`
	Address  string `json:"address"`
	Symbol   string `json:"symbol"`
	Name     string `json:"name"`
	Decimals int64  `json:"decimals"`
}

func (repo *CoinRepo) GetAllTokens(chainId string, defi string) error {
	query := `select name, symbol, address, decimals from coin where chainid = $1 and src = $2;`
	rows, err := db.PSQL.Query(query, chainId, defi)
	if err != nil {
		return err
	}
	defer rows.Close()

	var coin = &Coin{}
	for rows.Next() {
		err := rows.Scan(&coin.Name, &coin.Symbol, &coin.Address, &coin.Decimals)
		if err != nil {
			return err
		}
		repo.Coins = append(repo.Coins, *coin)
	}
	return nil
}
