package server

import (
	e "github.com/dmxmss/tasks/error"
	"github.com/dmxmss/tasks/entities"
	"github.com/gin-gonic/gin"

	"net/http"
	"strings"
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

func (s *GinServer) JWTMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, entities.Error{Error: "unauthorized"})
			return
		}

		rawToken := strings.TrimPrefix(authHeader, "Bearer")
		if rawToken == authHeader {
			c.AbortWithStatusJSON(http.StatusUnauthorized, entities.Error{Error: "unauthorized"})
			return
		}

		token, err := s.service.ValidateToken(rawToken)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, entities.Error{Error: "unauthorized"})
			return
		}


		claims, ok := token.Claims.(entities.Claims)
		if !ok {
			c.AbortWithStatusJSON(http.StatusUnauthorized, entities.Error{Error: "unauthorized"})
			return
		}
		c.Set("user_id", claims.UserID)
		
		c.Next()
	}
}
