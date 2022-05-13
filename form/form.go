package form

import (
	"NoRewriting/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

/*
------------------------------
@Time : 5/12/2022 10:20 PM
@Author : Sariel Crescent
@File : form.go
@Software: GoLand
----------------------------
*/

func GetForm(c *gin.Context) {
	json := Forms{}
	c.ShouldBind(&json)
	insertForm(json)
	c.JSON(200, gin.H{
		"code": 200,
		"msg":  "success",
	})
}
func insertForm(json Forms) {
	name := json.GetName()
	cf := json.GetCodeforces()
	key := json.GetKey()
	secret := json.GetSecret()
	fmt.Println(secret)

	_, err := sql.Db.Exec("insert into code_forces.admin"+
		" (user_name, codeforces_id, `key`, secret)values (?,?,?,?)", name, cf, key, secret)
	if err != nil {
		log.Errorf("insert form fail %v", err)
	}
}
