package spider

import (
	_ "NoRewriting/sql"
	"testing"
)

/*
------------------------------
@Time : 5/5/2022 1:23 PM
@Author : Sariel Crescent
@File : storage_test.go
@Software: GoLand
----------------------------
*/

func TestAdd(t *testing.T) {
	userName := "Sariel_snow"
	key := "26827a9121c91af42b0ef134896f899dc2ecf146"
	secret := "b7674f19476573d743758bb75a5f548ed6cbfbec"
	mp, err := getSubmission(statusWithKey(userName, key, secret))
	if err != nil {
		t.Errorf("get submission error %v", err)
	}
	err = addDataToSubmission(userName, mp)
	if err != nil {
		t.Errorf("add data to submission error %v", err)
	}

}
func TestGym(t *testing.T) {
	userName := "Sariel_snow"
	key := "26827a9121c91af42b0ef134896f899dc2ecf146"
	secret := "b7674f19476573d743758bb75a5f548ed6cbfbec"
	gymWithKey(userName, key, secret)
}

func TestProblem_Reset(t *testing.T) {
	mp, err := getProblemSet(problemSetWithoutKey())
	if err != nil {
		t.Errorf("%v", err)
	}
	addProblem(mp)
}
