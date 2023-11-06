package controller

import (
	"explore_address/pkg/log"
	"explore_address/pkg/router"
	"explore_address/service/address/model/dao"
	"net/http"
	"strconv"
)

type Status struct {
	Status bool `json:"status"`
}

func CheckExistAddress(w http.ResponseWriter, r *http.Request) {
	address := r.URL.Query().Get("address")
	if len(address) <= 0 {
		router.ResponseBadRequest(w, "B.400", "Missing address param!")
		return
	}

	chainName := r.URL.Query().Get("chainName")
	if len(chainName) <= 0 {
		router.ResponseBadRequest(w, "B.400", "Missing chainName param!")
		return
	}

	typeAddress := r.URL.Query().Get("type")
	if len(chainName) <= 0 {
		router.ResponseBadRequest(w, "B.400", "Missing type param!")
		return
	}

	typeAddressInt, err := strconv.Atoi(typeAddress)
	if err != nil {
		log.Println(log.LogLevelError, "CheckExistAddress: strconv.Atoi(typeAddress)", err.Error())
		router.ResponseInternalError(w, "Missing chainName param!", err)
	}

	addressDAO := dao.Address{
		ChainName: chainName,
		Address:   address,
		Type:      int64(typeAddressInt),
	}

	exist, err := addressDAO.CheckExist()
	if err != nil {
		log.Println(log.LogLevelError, "CheckExistAddress: addressDAO.CheckExist()", err.Error())
		exist = false
	}

	status := Status{
		Status: exist,
	}

	router.ResponseSuccessWithData(w, "B.200", "Check exist address", status)
}

func GetDecimal(w http.ResponseWriter, r *http.Request) {
	
}
