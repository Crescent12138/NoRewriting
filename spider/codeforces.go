package spider

import (
	"encoding/json"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
)

/*
------------------------------
@Time : 5/5/2022 1:12 PM
@Author : Sariel Crescent
@File : codeforces.go
@Software: GoLand
----------------------------
*/

func getSubmission(url string) (map[int64]Submission, error) {
	value, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	bytes, err := ioutil.ReadAll(value.Body)
	if err != nil {
		return nil, err
	}
	sub := &Status{}
	err = json.Unmarshal(bytes, sub)
	if err != nil {
		log.Infof("imply %v", err)
	}
	if sub.Status != "OK" {
		log.Errorf("fail to visit url ,check key or secret is right")
		log.Errorf(url)
		return nil, err
	}
	var mp map[int64]Submission
	mp = make(map[int64]Submission)
	problem := sub.GetResult()
	for _, val := range problem {
		if val.Verdict != "OK" {
			continue
		}
		mp[val.Id] = *val
	}
	return mp, nil
}
