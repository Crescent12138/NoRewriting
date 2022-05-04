package main

import (
	"fmt"
	"io/ioutil"
)

/*
------------------------------
@Time : 5/4/2022 3:35 PM
@Author : Sariel Crescent
@File : main
@Software: GoLand
----------------------------
*/

func main() {
	files, err := ioutil.ReadDir(".")
	if err != nil {
		fmt.Println(err)
	}
	for _, f := range files {
		fmt.Println(f.Name())
	}

}
