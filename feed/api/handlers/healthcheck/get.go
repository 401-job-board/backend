package healthcheck

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func performCheck() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, "\n\n*** Service running successfully ***\n\n")
	}
}
