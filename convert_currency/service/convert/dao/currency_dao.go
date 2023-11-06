package dao

import (
	"convert-service/pkg/db"
)

type RepoCurrency struct {
	Currencies []Currency `json:"currencies"`
}

type Currency struct {
	Rate        float64
	Symbol      string `json:"symbol"`
	Name        string `json:"name"`
	Image       string `json:"image"`
	Createddate string `json:"createdDate"`
	Updateddate string `json:"updatedDate"`
}

func (repo *RepoCurrency) GetCurrencies() error {
	query := `SELECT symbol, name, image, createddate, updateddate FROM public.currency;`
	rows, err := db.PSQL.Query(query)
	if err != nil {
		return err
	}
	defer rows.Close()

	for rows.Next() {
		var currency = &Currency{}
		err := rows.Scan(&currency.Symbol, &currency.Name, &currency.Image, &currency.Createddate, &currency.Updateddate)
		if err != nil {
			return err
		}

		repo.Currencies = append(repo.Currencies, *currency)
	}
	return nil

}
