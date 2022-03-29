package posts

import (
	"github.com/fanfit/feed/models/posts/service"
	"github.com/gin-gonic/gin"
)

// Routes sets up resource specific routes on the engine instance
func Routes(r *gin.RouterGroup, service service.Service) {
	router := r.Group("/posts")
	router.GET("/", getAllPostings(service))
	router.GET("/jobs/", getJobs(service))
	router.GET("/applications/", getApplicants(service))

	router.POST("/", post(service))
}
