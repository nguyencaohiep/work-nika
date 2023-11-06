package dao

import (
	"convert-service/pkg/db"
	"convert-service/pkg/utils"
	"encoding/json"
	"fmt"
	"os"

	"github.com/google/uuid"
)

type Chain struct {
	Symbol    *string `json:"symbol"`
	Address   *string `json:"address"`
	Decimal   *int64  `json:"decimal"`
	CryptoId  *string `json:"cryptoId"`
	ChainName *string `json:"chainName"`
	CryptoSrc *string `json:"cryptoSrc"`
}

type RepoCryptos struct {
	Cryptos []Crypto `json:"cryptos"`
}

type Crypto struct {
	Id                           uuid.UUID          `json:"id"`
	CryptoId                     string             `json:"crypto"`
	CryptoCode                   string             `json:"cryptoCode"`
	Name                         string             `json:"name"`
	Symbol                       string             `json:"symbol"`
	Decimal                      *int64             `json:"decimal"`
	Address                      *string            `json:"address"`
	ThumbLogo                    string             `json:"thumbLogo"`
	SmallLogo                    string             `json:"smallLogo"`
	BigLogo                      string             `json:"bigLogo"`
	ChainId                      *string            `json:"chainId"`
	ChainName                    *string            `json:"chainName"`
	Description                  *string            `json:"description"`
	Score                        int64              `json:"score"`
	Socials                      map[string]*string `json:"socials"`
	Website                      string             `json:"website"`
	Explorer                     *string            `json:"explorer"`
	Multichain                   []Chain            `json:"multichain"`
	MarketcapUSD                 float64            `json:"marketcapusd"`
	TotalSupply                  string             `json:"totalSupply"`
	PriceUSD                     float64            `json:"priceUSD"`
	CreatedDate                  string             `json:"createdDate"`
	UpdatedDate                  string             `json:"updatedDate"`
	TotallpUSD                   *string            `json:"totalLpUSD"`
	Type                         string             `json:"type"`
	TotalVolume                  float64            `json:"totalVolume"`
	High24h                      float64            `json:"high24h"`
	Low24h                       float64            `json:"low24h"`
	PriceChange24h               float64            `json:"priceChange24h"`
	PriceChangePercentage24h     float64            `json:"priceChangePercentage24h"`
	MarketcapChange24h           float64            `json:"marketcapChange24h"`
	MarketcapChangePercentage24h float64            `json:"marketcapChangePercentage24h"`
	ATH                          float64            `json:"ath"`
	ATHChangePercent             float64            `json:"athChangePercentage"`
	ATHDate                      string             `json:"athDate"`
	ATL                          float64            `json:"atl"`
	ATLChangePercentage          float64            `json:"atlChangePercentage"`
	ATLDate                      string             `json:"atlDate"`
	CurrentPrice                 map[string]float64 `json:"currentPrice"`
}

func (repo *RepoCryptos) GetCryptos() error {
	query := `SELECT id, cryptoid, cryptocode, "name", symbol, 
	"decimal", address, thumblogo, smalllogo, biglogo, 
	chainid, chainname, description, score, socials, 
	website, explorer, multichain, marketcapusd, totalsupply, 
	priceusd, createddate, updateddate, totallpusd,
	type, totalVolume, high24h, low24h, priceChange24h,
	priceChangePercentage24h, marketcapChange24h, marketcapChangePercentage24h, ath, athChangePercentage,
	athDate, atl, atlChangePercentage, atlDate FROM public.crypto order by marketcapusd desc;`
	rows, err := db.PSQL.Query(query)
	if err != nil {
		return err
	}
	defer rows.Close()

	var multichainByte = []byte{}
	var socialByte = []byte{}
	for rows.Next() {
		var crypto = &Crypto{}
		err := rows.Scan(&crypto.Id, &crypto.CryptoId, &crypto.CryptoCode, &crypto.Name, &crypto.Symbol,
			&crypto.Decimal, &crypto.Address, &crypto.ThumbLogo, &crypto.SmallLogo, &crypto.BigLogo,
			&crypto.ChainId, &crypto.ChainName, &crypto.Description, &crypto.Score, &socialByte,
			&crypto.Website, &crypto.Explorer, &multichainByte, &crypto.MarketcapUSD, &crypto.TotalSupply,
			&crypto.PriceUSD, &crypto.CreatedDate, &crypto.UpdatedDate, &crypto.TotallpUSD,
			&crypto.Type, &crypto.TotalVolume, &crypto.High24h, &crypto.Low24h, &crypto.PriceChange24h,
			&crypto.PriceChangePercentage24h, &crypto.MarketcapChange24h, &crypto.MarketcapChangePercentage24h, &crypto.ATH, &crypto.ATLChangePercentage,
			&crypto.ATHDate, &crypto.ATL, &crypto.ATLChangePercentage, &crypto.ATLDate)
		if err != nil {
			return err
		}

		if len(multichainByte) > 0 {
			err = json.Unmarshal(multichainByte, &crypto.Multichain)
			if err != nil {
				return err
			}
		}

		if len(socialByte) > 0 {
			err = json.Unmarshal(socialByte, &crypto.Socials)
			if err != nil {
				return err
			}
		}

		repo.Cryptos = append(repo.Cryptos, *crypto)
	}
	return nil

}

func (repo *RepoCryptos) UpdateInfo() error {
	query := `UPDATE public.crypto SET  
	marketcapusd= $1, totalsupply= $2, priceusd= $3, totallpusd= $4, totalvolume= $5, 
	high24h= $6, low24h= $7, pricechange24h= $8, pricechangepercentage24h= $9, marketcapchange24h= $10, 
	marketcapchangepercentage24h= $11, ath= $12, athchangepercentage= $13, athdate= $14, atl= $15, 
	atlchangepercentage= $16, atldate= $17, updatedDate = $18 where cryptoid = $19;`

	for _, crypto := range repo.Cryptos {
		_, err := db.PSQL.Exec(query, crypto.MarketcapUSD, crypto.TotalSupply, crypto.PriceUSD, crypto.TotallpUSD, crypto.TotalVolume,
			crypto.High24h, crypto.Low24h, crypto.PriceChange24h, crypto.PriceChangePercentage24h, crypto.MarketcapChange24h,
			crypto.MarketcapChangePercentage24h, crypto.ATH, crypto.ATHChangePercent, crypto.ATHDate, crypto.ATL,
			crypto.ATLChangePercentage, crypto.ATLDate, utils.TimeNowVietNamString(), crypto.CryptoId)
		if err != nil {
			return err
		}
	}
	return nil
}

func (repo *RepoCryptos) GetDes() error {
	query := `select description, symbol from crypto order by marketcapusd desc limit 100;`
	rows, err := db.PSQL.Query(query)
	if err != nil {
		return err
	}
	defer rows.Close()

	file, err := os.Create("example.txt")
	if err != nil {
		fmt.Println(err)
		return nil
	}
	defer file.Close()
	for rows.Next() {
		var crypto = &Crypto{}
		err := rows.Scan(&crypto.Description, &crypto.Symbol)
		if err != nil {
			return err
		}
		dat := fmt.Sprintf(`update crypto set des = '%v' where symbol = '%v';`, *crypto.Description, crypto.Symbol)

		_, err = file.WriteString(dat)
		if err != nil {
			fmt.Println(err)
			return nil
		}
	}
	return nil
}
