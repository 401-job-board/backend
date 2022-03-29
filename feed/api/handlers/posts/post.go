package posts

import (
	"fmt"
	"net/http"

	"github.com/fanfit/feed/api/views"
	"github.com/fanfit/feed/models/posts/repository"
	"github.com/fanfit/feed/models/posts/service"
	"github.com/gin-gonic/gin"
)

func post(service service.Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		var input repository.Posting
		if err := c.ShouldBind(&input); err != nil {
			views.Wrap(err, c)
			return
		}
		fmt.Println("About to create")
		response, err := service.CreatePosting(c.Request.Context(), input)
		if err != nil {
			views.Wrap(err, c)
			return
		}
		c.JSON(http.StatusAccepted, response)
	}
}
