package model

import (
	"info_project_service/pkg/db"
	"info_project_service/pkg/utils"
)

type ProjectStatus struct {
	IsScam      bool   `json:"isScam"`
	Reason      string `json:"reason"`
	Source      string `json:"source"`
	UserId      string `json:"userId"`
	Createddate string `json:"createddate"`
}

func (projectStatus *ProjectStatus) InsertStatus() error {
	query := `insert into project project_sstatus ( isScam, reason, source, user_id, createddate) values 
	($1, $2, $3, $4, $5);`
	_, err := db.PSQL.Exec(query, projectStatus.IsScam, projectStatus.Reason, projectStatus.Source, projectStatus.UserId, utils.TimeNowString())
	if err != nil {
		return err
	}
	return nil
}
