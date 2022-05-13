package submission

import (
	"NoRewriting/sql"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

/*
------------------------------
@Time : 5/12/2022 10:55 PM
@Author : Sariel Crescent
@File : submission
@Software: GoLand
----------------------------
*/

// PostSubmission 暂时没写相关token，就只能凑合把所有人的都输出出来，反正现在也只有我一个的
func PostSubmission(c *gin.Context) {
	var rank, num []int
	for i := 800; i <= 3500; i += 100 {
		res := getRankNumber(i)
		if res == 0 {
			continue
		}
		rank = append(rank, i)
		num = append(num, res)
	}
	c.JSON(200, gin.H{
		"xx": rank,
		"yy": num,
	})
}
func getRankNumber(num int) int {
	row, err := sql.Db.Query("select count(*) from code_forces.submission where rating=?", num)
	if err != nil {
		log.Errorf("fail to get count %v", err)
		return 0
	}
	for row.Next() {
		var tmp int
		row.Scan(&tmp)
		return tmp
	}
	return 0
}
func GetPassedProblems(c *gin.Context) {
	json := make(map[string]interface{})
	c.ShouldBind(&json)
	num := res
}

func getPassedProblem() {

}
