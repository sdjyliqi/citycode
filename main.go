package main

import (
	"citycode/apis"
	"citycode/middleware"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.Use(middleware.AddReqIDInMiddleware())
	router.Use(middleware.AddReqLogInMiddleware())
	router.GET("/", apis.ServerStatus)
	router.Run(":11000")
}
