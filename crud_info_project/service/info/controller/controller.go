package controller

import (
	"encoding/json"
	"info_project_service/pkg/log"
	"info_project_service/pkg/router"
	"info_project_service/service/info/model"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
)

type Info struct {
	Type          string       `json:"type"`
	ChainId       int8         `json:"chainId"`
	Address       string       `json:"address"`
	Symbol        string       `json:"symbol"`
	Name          string       `json:"name"`
	Decimals      string       `json:"decimals"`
	Tag           string       `json:"tag"`
	MarketCap     string       `json:"marketCap"`
	TotalSupply   string       `json:"totalSupply"`
	MaxSupply     string       `json:"maxSupply"`
	Image         string       `json:"image"`
	VolumeTrading string       `json:"volumeTrading"`
	Description   string       `json:"description"`
	Founder       string       `json:"founder"`
	Holders       string       `json:"holders"`
	Websites      string       `json:"websites"`
	SourceCode    string       `json:"sourceCode"`
	Socail        Social       `json:"socail"`
	MoreInfo      []model.More `json:"moreInfo"`
}

type Social struct {
	Facebook string `json:"facebook"`
	Twitter  string `json:"twitter"`
	Discord  string `json:"discord"`
	Telegram string `json:"telegram"`
}

func AddProjectInfo(w http.ResponseWriter, r *http.Request) {
	var infoForm = &Info{}
	err := json.NewDecoder(r.Body).Decode(infoForm)
	if err != nil {
		log.Println(log.LogLevelDebug, "AddProjectInfo: json.NewDecoder(r.Body).Decode(infoForm)", err)
		router.ResponseInternalError(w, err.Error())
		return
	}

	project := &model.Project{
		ChainId:       infoForm.ChainId,
		Type:          infoForm.Type,
		Address:       infoForm.Address,
		Symbol:        infoForm.Symbol,
		Name:          infoForm.Name,
		Tag:           infoForm.Tag,
		MarketCap:     infoForm.MarketCap,
		TotalSupply:   infoForm.TotalSupply,
		MaxSupply:     infoForm.MaxSupply,
		Image:         infoForm.Image,
		VolumeTrading: infoForm.VolumeTrading,
	}

	detail := model.Detail{
		Description: infoForm.Description,
		Founder:     infoForm.Founder,
		Holders:     infoForm.Holders,
		Websites:    infoForm.Websites,
		SourceCode:  infoForm.SourceCode,
		MoreInfo:    infoForm.MoreInfo,
	}

	community := model.Community{
		Facebook: infoForm.Socail.Facebook,
		Twitter:  infoForm.Socail.Twitter,
		Telegram: infoForm.Socail.Telegram,
		Discord:  infoForm.Socail.Discord,
	}
	detail.Community = community
	project.Detail = detail

	err = project.InsertInfo()
	if err != nil {
		log.Println(log.LogLevelError, "project.InsertInfo()", err)
		router.ResponseBadRequest(w, "Add project info failed!", "")
	}

	router.ResponseCreatedWithData(w, "Add project info successfully!", "")
}

func UpdateProjectInfo(w http.ResponseWriter, r *http.Request) {
	var infoForm = &Info{}
	err := json.NewDecoder(r.Body).Decode(infoForm)
	if err != nil {
		log.Println(log.LogLevelDebug, "UpdateProjectInfo: json.NewDecoder(r.Body).Decode(infoForm)", err)
		router.ResponseInternalError(w, err.Error())
		return
	}

	paramProjectId := chi.URLParam(r, "id")
	projectId, err := strconv.Atoi(paramProjectId)
	if err != nil {
		router.ResponseInternalError(w, err.Error())
		log.Println(log.LogLevelDebug, "UpdateProjectInfo: strconv.Atoi(paramProjectId)", err)
		return
	}

	project := &model.Project{
		Id:            int64(projectId),
		ChainId:       infoForm.ChainId,
		Type:          infoForm.Type,
		Address:       infoForm.Address,
		Symbol:        infoForm.Symbol,
		Name:          infoForm.Name,
		Tag:           infoForm.Tag,
		MarketCap:     infoForm.MarketCap,
		TotalSupply:   infoForm.TotalSupply,
		MaxSupply:     infoForm.MaxSupply,
		Image:         infoForm.Image,
		VolumeTrading: infoForm.VolumeTrading,
	}

	detail := model.Detail{
		Description: infoForm.Description,
		Founder:     infoForm.Founder,
		Holders:     infoForm.Holders,
		Websites:    infoForm.Websites,
		SourceCode:  infoForm.SourceCode,
		MoreInfo:    infoForm.MoreInfo,
	}

	community := model.Community{
		Facebook: infoForm.Socail.Facebook,
		Twitter:  infoForm.Socail.Twitter,
		Telegram: infoForm.Socail.Telegram,
		Discord:  infoForm.Socail.Discord,
	}
	detail.Community = community
	project.Detail = detail

	err = project.UpdateInfo()
	if err != nil {
		log.Println(log.LogLevelError, "project.UpdateInfo()", err)
		router.ResponseBadRequest(w, "Update project info failed!", "")
	}

	router.ResponseCreatedWithData(w, "Update project info successfully!", "")
}

func RemoveProject(w http.ResponseWriter, r *http.Request) {
	paramProjectId := chi.URLParam(r, "id")
	projectId, err := strconv.Atoi(paramProjectId)
	if err != nil {
		log.Println(log.LogLevelDebug, "UpdateProjectInfo: strconv.Atoi(paramProjectId)", err)
		router.ResponseInternalError(w, err.Error())
		return
	}

	project := &model.Project{
		Id: int64(projectId),
	}
	err = project.DeleteProject()
	if err != nil {
		log.Println(log.LogLevelError, "project.UpdateInfo()", err)
		router.ResponseBadRequest(w, "Update project info failed!", "")
	}

	router.ResponseCreatedWithData(w, "Update project info successfully!", "")
}
