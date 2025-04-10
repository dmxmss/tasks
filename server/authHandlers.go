package server

import (
	"github.com/dmxmss/tasks/entities"
	e "github.com/dmxmss/tasks/error"
	"github.com/gin-gonic/gin"
)

func (s *GinServer) SignUp(c *gin.Context) {
	var createUser entities.CreateUserDto

	if err := c.ShouldBindJSON(&createUser); err != nil {
		c.Error(e.ErrInvalidRequestBody)
		return
	}

	user, err := s.service.CreateUser(createUser)
	if err != nil {
		c.Error(err)
		return
	}

	accessToken, refreshToken, err := s.service.GenerateTokens(user.ID)
	if err != nil {
		c.Error(err)
		return
	}

	c.SetCookie(
    "access_token",
    *accessToken,
		s.conf.Auth.Access.ExpirationTime, 
		"/", 
		"", 
		true, 
		true,
	)

	c.SetCookie(
    "refresh_token",
    *refreshToken,
		s.conf.Auth.Refresh.ExpirationTime, 
		"/auth/refresh", 
		"", 
		true, 
		true,
	)
}
