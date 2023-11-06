package dao

type CryptoRepo struct {
	Cryptos []Crypto `json:"cryptos"`
}

type Crypto struct {
	CryptoId         string
	Contractverified bool  `json:"contractVerified"`
	Holders          int64 `json:"holders"`
}
