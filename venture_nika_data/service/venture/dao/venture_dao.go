package dao

import (
	"encoding/json"
	"fmt"
	"venture-data-service/pkg/db"
)

type VentureRepo struct {
	Ventures []Venture
}

type CompactVentureRepo struct {
	Name     string           `json:"name"`
	Count    int              `json:"count"`
	Ventures []CompactVenture `json:"ventures"`
}

type Venture struct {
	VentureId          string            `json:"ventureId"`
	VentureSrc         string            `json:"ventureSrc"`
	VentureCode        string            `json:"ventureCode"`
	VentureName        string            `json:"ventureName"`
	VentureLogo        string            `json:"ventureLogo"`
	YearFounded        *int              `json:"yearFounded"`
	Location           *string           `json:"location"`
	Description        string            `json:"description"`
	Socials            map[string]string `json:"socials"`
	SourceUrl          string            `json:"sourceUrl"`
	Website            *string           `json:"website"`
	TotalFund          *int64            `json:"totalFund"`
	Subcategory        *string           `json:"subcategory"`
	Score              float64           `json:"score"`
	StatisticsCategory map[string]int    `json:"statisticsCategory"`
	Contact            Contact
}

type CompactVenture struct {
	VentureId   *string  `json:"ventureId"`
	VentureName *string  `json:"ventureName"`
	VentureLogo *string  `json:"ventureLogo"`
	YearFounded *string  `json:"yearFounded"`
	Location    *string  `json:"location"`
	Score       *float64 `json:"score"`
}

type Contact struct {
	Normal   string `json:"normal"`
	LinkedIn string `json:"linkedIn"`
}

func (repo *VentureRepo) GetVentureCodes() error {
	socialByte := []byte{}
	query := `SELECT venturecode, venturename, socials FROM public.venture where socials is not null;`
	rows, err := db.PSQL.Query(query)
	if err != nil {
		fmt.Println(1)

		return err
	}
	defer rows.Close()

	for rows.Next() {
		venture := &Venture{}
		err := rows.Scan(&venture.VentureCode, &venture.VentureName, &socialByte)
		if err != nil {
			fmt.Println(2)
			return err
		}

		if len(socialByte) > 0 {
			err = json.Unmarshal(socialByte, &venture.Socials)
			if err != nil {
				fmt.Println(3)

				return err
			}
		}

		repo.Ventures = append(repo.Ventures, *venture)
	}
	return nil
}

func (venture *Venture) UpdateContact() error {
	var contactByte []byte
	contactByte, err := json.Marshal(venture.Contact)
	if err != nil {
		return err
	}
	query := `UPDATE public.venture SET  contact = $1 where venturecode = $2;`
	_, err = db.PSQL.Exec(query, contactByte, venture.VentureCode)
	if err != nil {
		return err
	}
	return nil
}
