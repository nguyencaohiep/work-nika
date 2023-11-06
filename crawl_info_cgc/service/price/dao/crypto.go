package dao

type ListCrypto struct {
	CryptoSrc string   `json:"cryptosrc"`
	Cryptos   []Crypto `json:"cryptos"`
}

type Crypto struct {
	CryptoId                     string  `json:"cryptoid"`
	Name                         string  `json:"name"`
	CryptoSrc                    string  `json:"cryptosrc"`
	Cryptocode                   string  `json:"cryptocode"`
	Symbol                       string  `json:"symbol"`
	Address                      string  `json:"address"`
	MarketcapUSD                 float64 `json:"marketcapusd"`
	TotalSupply                  string  `json:"totalSupply"`
	PriceUSD                     float64 `json:"priceUSD"`
	PricePercentChange24h        float64 `json:"pricePercentChange24h"`
	TotalVolume                  float64 `json:"totalVolume"`
	High24h                      float64 `json:"high24h"`
	Low24h                       float64 `json:"low24h"`
	PriceChange24h               float64 `json:"priceChange24h"`
	PriceChangePercentage24h     float64 `json:"priceChangePercentage24h"`
	MarketcapChange24h           float64 `json:"marketcapChange24h"`
	MarketcapChangePercentage24h float64 `json:"marketcapChangePercentage24h"`
	ATH                          float64 `json:"ath"`
	ATHChangePercent             float64 `json:"athChangePercentage"`
	ATHDate                      string  `json:"athDate"`
	ATL                          float64 `json:"atl"`
	ATLChangePercentage          float64 `json:"atlChangePercentage"`
	ATLDate                      string  `json:"atlDate"`
}
