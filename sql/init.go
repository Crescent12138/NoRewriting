package sql

/*
------------------------------
@Time : 5/4/2022 3:39 PM
@Author : Sariel Crescent
@File : init.go
@Software: GoLand
----------------------------
*/

import (
	"NoRewriting/config"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	log "github.com/sirupsen/logrus"
)

var (
	Db *sql.DB
)

func init() {

	msg := config.UserMySQL.App
	password := msg.Password
	host := msg.Host
	dbname := msg.Username
	user := msg.Username
	dsn := fmt.Sprintf("%v:%v@tcp(%v)/%v", user, password, host, dbname)
	var err error
	Db, err = sql.Open("mysql", dsn)
	if err != nil {
		log.Errorf("Open Sql Error, %v", err)
	}
}
