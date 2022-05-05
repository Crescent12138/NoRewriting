package spider

import (
	"NoRewriting/sql"
	log "github.com/sirupsen/logrus"
)

/*
------------------------------
@Time : 5/4/2022 3:38 PM
@Author : Sariel Crescent
@File : storage
@Software: GoLand
----------------------------
*/

func addDataToSubmission(userName string, mp map[int64]Submission) error {
	db := sql.Db
	query := "insert into submission(`Id`,user_Id,contest_id,problem_index,problem_name,rating)value(?,?,?,?,?,?)"
	stmt, err := db.Prepare(query)
	if err != nil {
		return err
	}
	for _, val := range mp {
		_, err = stmt.Exec(val.Id, userName, val.ContestId, val.Problem.Index, val.Problem.Name, val.Problem.Rating)
		if err != nil {
			log.Infof("add data %v", err)
		}
	}
	return nil
}
