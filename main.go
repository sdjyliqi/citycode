package main

import (
	"citycode/middleware"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.Use(middleware.AddReqIDInMiddleware())
	router.Use(middleware.AddReqLogInMiddleware())
	router.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello World")
	})
	router.Run(":11000")
}
