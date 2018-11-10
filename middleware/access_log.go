package middleware

import (
	//"github.com/gogap/logrus"
	"time"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

/*
函数功能：增加响应基本信息
 */
func AddReqLogInMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		req := c.Request
		res := c.Writer
		start := time.Now()
		c.Next()
		latency := time.Since(start)
		log.WithFields(log.Fields{
			"method":    req.Method,
			"ip": c.ClientIP(),
			"uri":       req.RequestURI,
			"status":    res.Status(),
			"latency":   int64(latency),
			"outsize":     res.Size(),
			"ua": req.UserAgent(),
		}).Info("access_log")
	}
}
