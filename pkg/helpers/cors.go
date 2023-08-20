package helpers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Cors(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, PATCH")
	if c.Request.Method == "OPTIONS" {
		if len(c.GetHeader("Access-Control-Request-Headers")) > 0 {
			c.Header("Access-Control-Allow-Headers", c.GetHeader("Access-Control-Request-Headers"))
		}
		c.JSON(http.StatusOK, "ok")
	}
	c.Next()

}
