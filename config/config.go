package config

/*
------------------------------
@Time : 5/4/2022 3:52 PM
@Author : Sariel Crescent
@File : config
@Software: GoLand
----------------------------
*/

type Config struct {
	App *App `yaml:"app"`
	Log *Log `yaml:"log"`
}

type App struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Database string `yaml:"database"`
}

type Log struct {
	Suffix  string `yaml:"suffix"`
	MaxSize int    `yaml:"maxSize"`
}
