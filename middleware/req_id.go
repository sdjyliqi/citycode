package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/satori/go.uuid"
	log "github.com/sirupsen/logrus"
)

func AddReqIDInMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		u1, err := uuid.NewV4()
		if err != nil {
			fmt.Printf("Something went wrong: %s", err)
			return
		}
		strUUID := fmt.Sprint(u1)
		c.Writer.Header().Set("X-Request-Id", strUUID)
		c.Next()
	}
}

func AddReqLogInMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		//打印请求体的log
		log.Println("REQ:", c.Request)
		c.Next()

	}
}
