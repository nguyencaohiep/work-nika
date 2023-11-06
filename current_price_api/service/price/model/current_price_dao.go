package model

import (
	"price_service/pkg/db"
	"price_service/pkg/log"
	"price_service/pkg/utils"
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

func (price *CurrentPrice) InsertPrice() error {
	query := `insert into current_price (tokenaddress, symbol, name, chainId, price, source, createddate, updatedate) 
		values ($1, $2, $3, $4, $5, $6, $7, $8);`
	_, err := db.PSQL.Exec(query, price.TokenAddress, price.Symbol, price.Name, price.ChainId, price.Price, price.Source, utils.TimeNowString(), utils.TimeNowString())
	if err != nil {
		log.Println(log.LogLevelError, "InsertPrice", "")
		return err
	}
	return nil
}

func (price *CurrentPrice) CheckExistPrice() (bool, error) {
	var exist bool
	query := `SELECT exists(select 1 FROM current_price WHERE tokenaddress = $1 and chainId = $2);`
	err := db.PSQL.QueryRow(query, price.TokenAddress, price.ChainId).Scan(&exist)
	return exist, err
}
