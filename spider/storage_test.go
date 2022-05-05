package spider

import "testing"

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
	mp, err := getSubmission(statusWithoutKey(userName))
	if err != nil {
		t.Errorf("get submission error %v", err)
	}
	err = addDataToSubmission(userName, mp)
	if err != nil {
		t.Errorf("add data to submission error %v", err)
	}

}
