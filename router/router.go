package router

import (
	"NoRewriting/form"
	"NoRewriting/login"
	"NoRewriting/submission"
)

/*
------------------------------
@Time : 5/12/2022 9:05 PM
@Author : Sariel Crescent
@File : router.go
@Software: GoLand
----------------------------
*/
import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func ListenRouter() {
	r := gin.Default()
	r.Use(cors.Default())
	r.POST("/api/login", login.LogIn)
	r.POST("/api/form", form.GetForm)
	r.POST("/api/submission", submission.PostSubmission)
	r.POST("/api/register", login.RegisterTo)
	err := r.Run(":8888")
	if err != nil {
		return
	}
}
