package crawler

import (
	"venture-data-service/pkg/log"
	"venture-data-service/service/venture/dao"
)

var (
	VentureRepo dao.VentureRepo
)

func init() {
	GetVentureCode()
}

func GetVentureCode() {
	repo := dao.VentureRepo{}
	err := repo.GetVentureCodes()
	if err != nil {
		log.Println(log.LogLevelError, "GetVentureCode: repo.GetVentureCodes()", err)
	}

	VentureRepo = repo
}
