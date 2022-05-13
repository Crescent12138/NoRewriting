package login

import (
	"NoRewriting/sql"
	"github.com/gin-gonic/gin"
)

/*
------------------------------
@Time : 5/12/2022 10:52 PM
@Author : Sariel Crescent
@File : register
@Software: GoLand
----------------------------
*/

// RegisterTo  这个分页用于实现注册，因为注册的前端没写，暂时懒得搞
func RegisterTo(c *gin.Context) {
	json := Register{}
	c.ShouldBind(&json)
	password := json.GetPassword()
	userName := json.GetUserName()
	_, err := sql.Db.Exec("insert into code_forces.admin_login"+
		"(user_id, password)values(?,?)", userName, password)
	if err != nil {
		c.JSON(507, gin.H{
			"code": 507,
			"msg":  "repeat userName",
		})
	} else {
		c.JSON(200, gin.H{
			"code": 200,
			"msg":  "ok",
		})
	}

}
