package apis

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func ServerStatus(c *gin.Context) {
	c.String(http.StatusOK, "Hello,I am ok")
}
