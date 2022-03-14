package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main(){
	r:=gin.Default()
	r.LoadHTMLGlob("passcode/index.html")

	r.Static("/passcode","./passcode")

	//r.StaticFS("/passcode",http.Dir("./passcode"))

	r.GET("index",func(c*gin.Context){
		c.HTML(http.StatusOK,"index.html",nil)
	})
	r.Run()
}