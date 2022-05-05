package spider

import (
	"net/http"
	"testing"
)

/*
------------------------------
@Time : 5/5/2022 11:37 AM
@Author : Sariel Crescent
@File : public_test.go
@Software: GoLand
----------------------------
*/

func TestStatus(t *testing.T) {
	userName := "Sariel_snow"
	key := "26827a9121c91af42b0ef134896f899dc2ecf146"
	secret := "b7674f19476573d743758bb75a5f548ed6cbfbec"
	wget := statusWithoutKey(userName)
	_, err := http.Get(wget)
	if err != nil {
		t.Errorf("wget status without key %v", err)
	}
	wget = statusWithKey(userName, key, secret)
	_, err = http.Get(wget)
	if err != nil {
		t.Errorf("wget status with key %v", err)
	}
}
