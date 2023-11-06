package model

import (
	"native_service/pkg/db"
	"native_service/pkg/utils"
)

type CurrentPrice struct {
	TokenAddress string `json:"tokenAddress"`
	Symbol       string `json:"symbol"`
	Name         string `json:"name"`
	ChainId      string `json:"chainId"`
	Price        string `json:"price"`
	Source       string `json:"source"`
	Createddate  string `json:"createddate"`
	Updateddate  string `json:"updateddate"`
}

func (coin *Coin) InsertPrice() error {
	query := `insert into current_price (tokenaddress, symbol, name, chainId, price, source, createddate, updatedate) 
	values ($1, $2, $3, $4, $5, $6, $7, $8);`
	_, err := db.PSQL.Exec(query, coin.Address, coin.Symbol, coin.Name, coin.ChainId, coin.Price, coin.Source, utils.TimeNowString(), utils.TimeNowString())
	if err != nil {
		return err
	}
	return nil
}
