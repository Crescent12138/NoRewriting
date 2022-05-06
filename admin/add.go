package admin

import "NoRewriting/sql"

/*
------------------------------
@Time : 5/5/2022 1:57 PM
@Author : Sariel Crescent
@File : add.go
@Software: GoLand
----------------------------
*/

//用于添加新用户，添加管理员相关
func addAdmin(users User) error {
	db := sql.Db
	sql := "insert into admin (`user_id`,`password`,`key`,`secret`)value( ?,?,?,?)"
	_, err := db.Exec(sql, users.UserName, users.Password, users.Key, users.Secret)
	return err
}
