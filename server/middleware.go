package server

import (
	e "github.com/dmxmss/tasks/error"
	"github.com/dmxmss/tasks/entities"
	"github.com/gin-gonic/gin"

	"net/http"
)

func ErrorMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		errors := c.Errors

		if len(errors) > 0 {
			var err entities.Error
			var statusCode int

			switch errors.Last().Err {
			case e.ErrDbUserForeignKeyViolation:
				err.Error = "invalid request: user not found"
				statusCode = http.StatusBadRequest
			case e.ErrDbTaskNotFound:
				err.Error = "task not found"
				statusCode = http.StatusNotFound
			default:
				err.Error = "internal server error"
				statusCode = http.StatusInternalServerError
			}

			c.JSON(statusCode, err)
		}
	}
}
