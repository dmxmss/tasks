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

	user, err := s.service.LogIn(loginData)
	if err != nil {
		c.Error(err)
		return
	}
	
	accessToken, refreshToken, err := s.service.GenerateTokens(user.ID)
	if err != nil {
		c.Error(err)
		return
	}

	u.WriteTokenToCookies(c, *accessToken, 
													 *refreshToken, 
													 s.conf.Auth.Access.ExpirationTime,
													 s.conf.Auth.Refresh.ExpirationTime)
}

func (s *GinServer) GetUserInfo(c *gin.Context) {
	v, exists := c.Get("claims")
	claims, ok := v.(*entities.Claims)

	if !ok || !exists {
		c.Error(e.ErrAuthFailed)
		return
	}

	user, err := s.service.GetUserInfo(claims.UserID)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, user)
}
