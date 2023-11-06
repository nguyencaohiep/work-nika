package controller

import (
	"fmt"
	"math"
	"net/http"
	"score_gear5/pkg/db"
	"score_gear5/pkg/utils"
	"strconv"
	"sync"
	"time"
)

type CryptoScore struct {
	CryptoId         *string  `json:"cryptoId"`
	ThumbLogo        *string  `json:"thumbLogo"`
	SmallLogo        *string  `json:"smallLogo"`
	BigLogo          *string  `json:"bigLogo"`
	Score            *int64   `json:"score"`
	IsScam           *bool    `json:"isScam"`
	IsWarning        *bool    `json:"isWarning"`
	IsProxy          *bool    `json:"isProxy"`
	Reputation       *int64   `json:"reputation"`
	Website          *string  `json:"website"`
	MarketcapUSD     *float64 `json:"marketcapUSD"`
	TotalSupply      *string  `json:"totalSupply"`
	TotalLPUSD       *string  `json:"totalLPUSD"`
	PriceUSD         *float64 `json:"priceUSD"`
	Holders          *int64   `json:"holders"`
	Transfers        *int64   `json:"transfers"`
	IsCoingecko      *bool    `json:"isCoingecko"`
	IsCoinmarketcap  *bool    `json:"isCoinmarketcap"`
	IsBinance        *bool    `json:"isBinance"`
	IsCoinbase       *bool    `json:"isCoinbase"`
	IsDex            *bool    `json:"isDex"`
	ContractVerified *bool    `json:"contractVerified"`
}

func (cryptoScore *CryptoScore) CalculateCryptoScore() {
	var score int64 = 0

	reputation := cryptoScore.Reputation
	if reputation != nil {
		score += *reputation
	}

	isBinance := cryptoScore.IsBinance
	if isBinance != nil {
		if *isBinance {
			score += 15
		}
	}

	isCoinbase := cryptoScore.IsCoinbase
	if isCoinbase != nil {
		if *isCoinbase {
			score += 15
		}
	}

	isCoingecko := cryptoScore.IsCoingecko
	if isCoingecko != nil {
		if *isCoingecko {
			score += 5
		}
	}

	isCoinmarketcap := cryptoScore.IsCoinmarketcap
	if isCoinmarketcap != nil {
		if *isCoinmarketcap {
			score += 5
		}
	}

	isDex := cryptoScore.IsDex
	if isDex != nil {
		if *isDex {
			score += 2
		}
	}

	isContractVerified := cryptoScore.ContractVerified
	if isContractVerified != nil {
		if !*isContractVerified {
			score -= 5
		}
	}

	isProxy := cryptoScore.IsProxy
	if isProxy != nil {
		if *isProxy {
			score -= 20
		}
	}

	// TODO: total lp
	totalLP := cryptoScore.TotalLPUSD
	if totalLP != nil {
		floatLP, err := strconv.ParseFloat(*totalLP, 64)
		if err != nil {
			fmt.Println(err)
		}
		totalLPScore := math.Min(math.Max(0, math.Log10(floatLP)-2), 10)
		score += int64(totalLPScore)
	}
	// TODO: transactions in 24H

	isScam := cryptoScore.IsScam
	if isScam != nil {
		if *isScam {
			score -= 100
		}
	}

	isWarning := cryptoScore.IsWarning
	if isWarning != nil {
		if *isWarning {
			score -= 50
		}
	}

	haveLogo := cryptoScore.ThumbLogo != nil || cryptoScore.SmallLogo != nil || cryptoScore.BigLogo != nil
	if haveLogo {
		score += 3
	}

	haveWebsite := cryptoScore.Website != nil
	if haveWebsite {
		score += 10
	}

	// TODO: funded by BIG Invester

	totalHolders := cryptoScore.Holders
	if totalHolders != nil {
		if *totalHolders > 0 {
			holderScore := math.Min(math.Max(0, 2*math.Log10(float64(*totalHolders))-2), 10)
			score += int64(holderScore)

		}
	}

	totalTransfers := cryptoScore.Transfers
	if totalTransfers != nil {
		if *totalTransfers > 0 {
			transfersScore := math.Min(math.Max(0, math.Log10(float64(*totalTransfers))-2), 10)
			score += int64(transfersScore)

		}
	}
	cryptoScore.Score = &score
}

func (cryptoScore *CryptoScore) UpdateCryptoScore() error {
	query := `update crypto set score = $1, updatedDate = $2 where cryptoid = $3`
	// fmt.Println(*cryptoScore.CryptoId, *cryptoScore.Score)
	// return nil
	_, err := db.PSQL.Exec(query, cryptoScore.Score, utils.TimeNowString(), cryptoScore.CryptoId)
	return err
}

type CryptoScoreRepo struct {
	CryptoScores []CryptoScore `json:"cryptoScores"`
}

func (repo *CryptoScoreRepo) GetCryptoScoresLimit(limit int64, offset int64) error {
	query := `select cryptoid, thumblogo, smalllogo, biglogo, 
			score, isscam, iswarning, isproxy, reputation, 
			website, marketcapusd, totalsupply, priceusd, 
			holders, transfers, iscoingecko , iscoinmarketcap, 
			isbinance, iscoinbase, isdex, contractverified, totallpusd
			from crypto where score is null order by cryptoid asc limit $1 offset $2;`

	rows, err := db.PSQL.Query(query, limit, offset)
	if err != nil {
		return err
	}
	defer rows.Close()

	for rows.Next() {
		var cryptoScore = CryptoScore{}
		err := rows.Scan(&cryptoScore.CryptoId, &cryptoScore.ThumbLogo, &cryptoScore.SmallLogo, &cryptoScore.BigLogo, &cryptoScore.Score, &cryptoScore.IsScam,
			&cryptoScore.IsWarning, &cryptoScore.IsProxy, &cryptoScore.Reputation, &cryptoScore.Website, &cryptoScore.MarketcapUSD, &cryptoScore.TotalSupply,
			&cryptoScore.PriceUSD, &cryptoScore.Holders, &cryptoScore.Transfers, &cryptoScore.IsCoingecko, &cryptoScore.IsCoinmarketcap, &cryptoScore.IsBinance,
			&cryptoScore.IsCoinbase, &cryptoScore.IsDex, &cryptoScore.ContractVerified, &cryptoScore.TotalLPUSD)
		if err != nil {
			return err
		}
		repo.CryptoScores = append(repo.CryptoScores, cryptoScore)

	}
	return nil
}

var (
	totalLimit   int64     = 2218550
	limit        int64     = 1000
	currentCount int64     = 0
	offset       int64     = 0
	startTime    time.Time = time.Now()
)

func Info(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "\tTotalLimit: %v,\n	CurrentCount: %v,\n	Limit: %v,\n	StartTime: %v\n	StartFor: %v h\n", totalLimit,
		currentCount+461000, limit, utils.TimeNowString(), time.Since(startTime).Hours())
}

func Score() {
	var threadNumber = 4
	var channel = make(chan int64, threadNumber+1)
	var crawlWaitGroup sync.WaitGroup
	crawlWaitGroup.Add(threadNumber)
	for thread := 0; thread < threadNumber; thread++ {
		fmt.Println("Thread ", thread+1)
		go func(thread int) {
			for {
				offset, stillRun := <-channel
				if !stillRun {
					crawlWaitGroup.Done()
					return
				}
				HandleCryptoScoresOffset(offset)
			}
		}(thread)
	}

	// Push blocknumer to goroutine
	for offset = 0; offset < totalLimit; offset += limit {
		currentOffset := offset
		channel <- currentOffset
	}

	close(channel)
	crawlWaitGroup.Wait()
	fmt.Println("(--------------------------DONE--------------------------------------)")
}

func HandleCryptoScoresOffset(offset int64) {
	cryptoScoreRepo := CryptoScoreRepo{}
	err := cryptoScoreRepo.GetCryptoScoresLimit(limit, offset)
	if err != nil {
		fmt.Println(err)
	}

	// fmt.Println("count -> ", len(cryptoScoreRepo.CryptoScores))
	for _, cryptoScore := range cryptoScoreRepo.CryptoScores {
		cryptoScore.CalculateCryptoScore()
		err = cryptoScore.UpdateCryptoScore()
		if err != nil {
			fmt.Println(err)
		}
		currentCount++
	}
}
