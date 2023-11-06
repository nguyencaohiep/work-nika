package model

import (
	"encoding/json"
	"info_project_service/pkg/db"
)

type ProjectRepo struct {
	Projects []Project
}

type Project struct {
	Id            int64  `json:"id"`
	CoinId        string `json:"coinId"`
	Type          string `json:"type"`
	Address       string `json:"address"`
	ChainId       int8   `json:"chainId"`
	Symbol        string `json:"symbol"`
	Name          string `json:"name"`
	Tag           string `json:"tag"`
	Decimals      string `json:"decimals"`
	MarketCap     string `json:"marketCap"`
	TotalSupply   string `json:"totalSupply"`
	MaxSupply     string `json:"maxSupply"`
	Image         string `json:"image"`
	VolumeTrading string `json:"volumeTrading"`
	IsVerify      bool   `json:"isVerify"`
	Detail        Detail `json:"detail"`
}

type Detail struct {
	Description string    `json:"description"`
	Founder     string    `json:"founder"`
	Holders     string    `json:"holders"`
	Websites    string    `json:"websites"`
	SourceCode  string    `json:"sourceCode"`
	Community   Community `json:"community"`
	MoreInfo    []More    `json:"moreInfo"`
}
type More struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type Community struct {
	Facebook string `json:"facebook"`
	Twitter  string `json:"twitter"`
	Discord  string `json:"discord"`
	Telegram string `json:"telegram"`
}

func (project *Project) InsertInfo() error {
	query := `insert into project_coin (coinid, type, address, symbol, chainId, name, tag, marketcap, maxsupply, isverify,  
		totalSupply, image, volumeTrading, detail) values ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14);`

	detailJson, err := json.Marshal(project.Detail)
	if err != nil {
		return err
	}
	_, err = db.PSQL.Exec(query, project.CoinId, project.Type, project.Address, project.Symbol, project.ChainId, project.Name,
		project.Tag, project.MarketCap, project.MaxSupply, false, project.TotalSupply, project.Image, project.VolumeTrading, detailJson)
	if err != nil {
		return err
	}
	return nil
}

func (project *Project) UpdateInfo() error {
	detailJson, err := json.Marshal(project.Detail)
	if err != nil {
		return err
	}
	query := `update project_coin set coinid = $1, type = $2, address = $3, symbol = $4, chainId = $5, name = $6, tag = $7,
		marketcap = $8, maxSupply = $9, totalSupply = $10, isverify = $11, image = $12, volumeTrading = $13, detail = $14;`
	_, err = db.PSQL.Exec(query, project.CoinId, project.Type, project.Address, project.Symbol, project.ChainId, project.Name,
		project.Tag, project.MarketCap, project.MaxSupply, project.TotalSupply, project.IsVerify, project.Image, project.VolumeTrading,
		detailJson)
	if err != nil {
		return err
	}
	return nil
}

func (project *Project) DeleteProject() error {
	query := `delete from project_coin where id = $1;`
	_, err := db.PSQL.Exec(query, project.Id)
	if err != nil {
		return err
	}
	return nil
}
