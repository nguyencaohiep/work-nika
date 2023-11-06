package model

import (
	"native_service/pkg/db"
)

type TokenNativeRepo struct {
	Tokens []TokenNative
}

type TokenNative struct {
	ChainId string `json:"chain"`
	Symbol  string `json:"symbol"`
	Name    string `json:"name"`
}

func (repo *TokenNativeRepo) InsertTokens(chainId string) error {
	query := `insert into token_native (chainid, name, symbol) values ($1, $2, $3);`
	rows, err := db.PSQL.Query(query)
	if err != nil {
		return err
	}
	defer rows.Close()

	var token = &TokenNative{}
	for rows.Next() {
		err := rows.Scan(&token.ChainId, &token.Name, &token.Symbol)
		if err != nil {
			return err
		}
		repo.Tokens = append(repo.Tokens, *token)
	}
	return nil
}
