package login

import (
	"NoRewriting/sql"
	"github.com/gin-gonic/gin"
)

/*
------------------------------
@Time : 5/12/2022 9:18 PM
@Author : Sariel Crescent
@File : login
@Software: GoLand
----------------------------
*/

// LogIn LoginFunction  用于登陆界面路由，实现登陆匹配
func LogIn(c *gin.Context) {

	//处理报文
	json := Login{}
	c.ShouldBind(&json)
	//fmt.Println(c)
	userName := json.GetUserName()
	password := json.GetPassword()
	if judgeUser(userName, password) {
		c.JSON(200, gin.H{
			"code": 200,
			"msg":  "ok",
		})
	} else {
		c.JSON(401, gin.H{
			"code": 501,
			"msg":  "wrong userName or password",
		})
	}
}

func judgeUser(userName string, password string) bool {
	db := sql.Db

	row, err := db.Query("select password from admin_login where user_id=?", userName)
	if err != nil {
		return false
	}
	for row.Next() {
		var tmp string
		row.Scan(&tmp)
		if tmp != password {
			return false
		}
		if tmp == password {
			return true
		}
	}
	return false
}
