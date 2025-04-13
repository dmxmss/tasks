package server

import (
	e "github.com/dmxmss/tasks/error"
	"github.com/dmxmss/tasks/entities"
	"github.com/gin-gonic/gin"

	"net/http"
	"strings"
	"errors"
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
			case e.ErrUserAlreadyExists:
				err.Error = "user already exists"
				statusCode = http.StatusConflict
			case e.ErrAuthInvalidCredentials:
				err.Error = "invalid credentials"
				statusCode = http.StatusUnauthorized
			default:
				err.Error = "internal server error"
				statusCode = http.StatusInternalServerError
			}

			c.JSON(statusCode, err)
		}
	}
}

func (s *GinServer) JWTAccessMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, entities.Error{Error: "missing authorization header"})
			return
		}

		rawToken := strings.TrimPrefix(authHeader, "Bearer ")
		if rawToken == authHeader {
			c.AbortWithStatusJSON(http.StatusUnauthorized, entities.Error{Error: "invalid authentication type"})
			return
		}

		claims, err := s.service.ValidateToken(rawToken)
		if err != nil {
			if errors.Is(err, e.ErrAuthTokenExpired) {
				c.AbortWithStatusJSON(http.StatusUnauthorized, entities.Error{Error: "token expired"})
			} else if errors.Is(err, e.ErrAuthTokenInvalid) {
				c.AbortWithStatusJSON(http.StatusUnauthorized, entities.Error{Error: "token is invalid"})
			} else {
				c.AbortWithStatusJSON(http.StatusUnauthorized, entities.Error{Error: "token expired"})
			}
			return
		}

		c.Set("claims", claims)
		
		c.Next()
	}
}

func (s *GinServer) JWTRefreshMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		refresh, err := c.Cookie("refresh_token")
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, entities.Error{Error: "missing refresh token"})
			return
		}

		claims, err := s.service.ValidateToken(refresh)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, entities.Error{Error: "token is invalid"})
			return
		}

		c.Set("claims", claims)
		c.Next()
	}
}
