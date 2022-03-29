package posts

import (
	"net/http"

	"github.com/fanfit/feed/api/views"
	"github.com/fanfit/feed/models/posts/repository"
	"github.com/fanfit/feed/models/posts/service"

	"github.com/gin-gonic/gin"
)

func getAllPostings(service service.Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		resource, err := service.GetPosts(c.Request.Context())
		if err != nil {
			views.Wrap(err, c)
			return
		}
		c.JSON(http.StatusOK, resource)
	}
}

func getJobs(service service.Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		var input repository.GetJobsParams
		if err := c.ShouldBind(&input); err != nil {
			views.Wrap(err, c)
			return
		}
		resource, err := service.GetJobs(c.Request.Context(), input)
		if err != nil {
			views.Wrap(err, c)
			return
		}
		c.JSON(http.StatusOK, resource)
	}
}
