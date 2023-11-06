package model

import (
	"native_service/pkg/db"
)

type CoinRepo struct {
	Coins []Coin
}

type Coin struct {
	ChainId string `json:"chain"`
	Address string `json:"address"`
	Symbol  string `json:"symbol"`
	Name    string `json:"name"`
	Price   string `json:"price"`
	Source  string `json:"source"`
}

func (repo *CoinRepo) GetAllCoins(address string) error {
	query := `select name, symbol from coin where address = $1;`
	rows, err := db.PSQL.Query(query, address)
	if err != nil {
		return err
	}
	defer rows.Close()

	var coin = &Coin{}
	for rows.Next() {
		err := rows.Scan(&coin.Name, &coin.Symbol)
		if err != nil {
			return err
		}
		repo.Coins = append(repo.Coins, *coin)
	}
	return nil
}
