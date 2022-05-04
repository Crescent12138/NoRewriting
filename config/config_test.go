package config

import "testing"

/*
------------------------------
@Time : 5/4/2022 10:08 PM
@Author : Sariel Crescent
@File : config_test.go
@Software: GoLand
----------------------------
*/

func Test(t *testing.T) {
	if UserMySQL.App.Port != 3306 {
		t.Errorf("no")
	}
}
