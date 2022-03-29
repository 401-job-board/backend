package healthcheck

import (
	"github.com/gin-gonic/gin"
)

// Routes sets up resource specific routes on the engine instance
// http://localhost:8080/v1/sns/
func Routes(r *gin.RouterGroup) {
	router := r.Group("/healthcheck")
	router.GET("/", performCheck())
}
