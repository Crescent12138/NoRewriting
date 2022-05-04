package main

import (
	"NoRewriting/sql"
	"testing"
)

/*
------------------------------
@Time : 5/4/2022 10:57 PM
@Author : Sariel Crescent
@File : init_test.go
@Software: GoLand
----------------------------
*/

func Test(t *testing.T) {

	if sql.Db == nil {
		t.Errorf("open db fail")
	}
}
