package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/satori/go.uuid"
)

/*
函数功能：为了方便方便后续查找问题，针对请求添加reqID
*/
func AddReqIDInMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		requestID := c.Request.Header.Get("X-Request-ID")
		if requestID == "" {
			requestID = uuid.NewV4().String()
		}
		c.Writer.Header().Set("X-Request-ID", requestID)
		c.Next()

	}
}
