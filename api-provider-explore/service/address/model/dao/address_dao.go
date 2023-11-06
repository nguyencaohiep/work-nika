package dao

import (
	"explore_address/pkg/db"
)

type Address struct {
	Address   string `json:"address"`
	ChainName string `json:"chainName"`
	Decimal   int64  `json:"decimal"`
	Type      int64  `json:"type"`
}

func (address *Address) GetDecimal() error {
	query := `select decimal from address where address = $1 and chainname = $2;`
	err := db.PSQL.QueryRow(query, address.Address, address.ChainName).Scan(&address.Decimal)
	if err != nil {
		return err
	}
	return nil
}

func (address *Address) CheckExist() (bool, error) {
	var exist bool
	query := `SELECT exists(select 1 FROM address WHERE address = $1 and chainname = S2 and type = $3);`
	err := db.PSQL.QueryRow(query, address.Address, address.ChainName, address.Type).Scan(&exist)
	return exist, err
}
