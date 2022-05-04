package sql

import "testing"

/*
------------------------------
@Time : 5/4/2022 11:54 PM
@Author : Sariel Crescent
@File : init_test.go
@Software: GoLand
----------------------------
*/

func Test(t *testing.T) {
	if Db == nil {
		t.Errorf("wrong in open Db")
	}
}
