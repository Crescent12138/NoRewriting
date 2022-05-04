package config

import (
	log "github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

/*
------------------------------
@Time : 5/4/2022 11:51 PM
@Author : Sariel Crescent
@File : init.go
@Software: GoLand
----------------------------
*/
var (
	UserMySQL = &Config{}
	path      = "C:\\Users\\Saiel Crescent\\Desktop\\NoRewriting\\config.yml"
)

func init() {
	yamlFile, err := ioutil.ReadFile(path)
	if err != nil {
		log.Errorf("yml open %v", err)
	}
	//var UserMySQL *config.Config
	err = yaml.Unmarshal(yamlFile, &UserMySQL)
	if err != nil {
		log.Errorf("yml unmarshal %v", err)
	}
}
