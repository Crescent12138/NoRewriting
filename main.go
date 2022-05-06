package main

import "net/http"

/*
------------------------------
@Time : 5/4/2022 3:35 PM
@Author : Sariel Crescent
@File : main
@Software: GoLand
----------------------------
*/

func main() {
	
	http.ListenAndServe(":8109", nil)
}
