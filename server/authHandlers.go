package server

import (
	"github.com/dmxmss/tasks/entities"
	u "github.com/dmxmss/tasks/internal/utils"
	e "github.com/dmxmss/tasks/error"
	"github.com/gin-gonic/gin"

	"net/http"
)

func (s *GinServer) SignUp(c *gin.Context) {
	var createUser entities.CreateUserDto

	if err := c.ShouldBindJSON(&createUser); err != nil {
		c.Error(e.ErrInvalidRequestBody)
		return
	}

	user, err := s.service.UserService.CreateUser(createUser)
	if err != nil {
		c.Error(err)
		return
	}

	accessToken, refreshToken, err := s.service.AuthService.GenerateTokens(user.ID, user.City)
	if err != nil {
		c.Error(err)
		return
	}

	u.WriteTokenToCookies(c, *accessToken, 
													 *refreshToken, 
													 s.conf.Auth.Access.ExpirationTime,
													 s.conf.Auth.Refresh.ExpirationTime)
}

func (s *GinServer) LogIn(c *gin.Context) {
	var loginData entities.LoginUserDto
	if err := c.ShouldBindJSON(&loginData); err != nil {
		c.Error(e.ErrInvalidRequestBody)
		return
	}

	user, err := s.service.UserService.LogIn(loginData)
	if err != nil {
		c.Error(err)
		return
	}
	
	access, refresh, err := s.service.AuthService.GenerateTokens(user.ID, user.City)
	if err != nil {
		c.Error(err)
		return
	}

	u.WriteTokenToCookies(c, *access, 
													 *refresh, 
													 s.conf.Auth.Access.ExpirationTime,
													 s.conf.Auth.Refresh.ExpirationTime)
}

func (s *GinServer) GetUserInfo(c *gin.Context) {
	claims, err := s.getClaims(c)
	if err != nil {
		c.Error(err)
		return
	}

	user, err := s.service.UserService.GetUserInfo(claims.UserID)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, user)
}

func (s *GinServer) UpdateTokens(c *gin.Context) {
	claims, err := s.getClaims(c)
	if err != nil {
		c.Error(err)
		return
	}

	access, refresh, err := s.service.AuthService.GenerateTokens(claims.UserID, claims.City)
	if err != nil {
		c.Error(err)
		return
	}
	
	u.WriteTokenToCookies(c, *access, 
													 *refresh, 
													 s.conf.Auth.Access.ExpirationTime,
													 s.conf.Auth.Refresh.ExpirationTime)
}

