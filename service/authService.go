package service

import (
	"github.com/dmxmss/tasks/entities"
	"github.com/golang-jwt/jwt/v5"
	
	"time"
)

type AuthService interface {
	ValidateToken(string) (*entities.Claims, error)
	GenerateToken(int, int) (*string, error)
	GenerateTokens(int) (*string, *string, error)
}

func (s *service) ValidateToken(rawToken string) (*entities.Claims, error) {
	claims, err := s.authRepo.ValidateToken(rawToken)
	if err != nil {
		return nil, err
	}

	return claims, nil
}

func (s *service) GenerateToken(userId, expirationTime int) (*string, error) {
	claims := s.getClaims(userId, expirationTime)
	token, err := s.authRepo.GenerateAndSignToken(claims)

	return token, err
}

func (s *service) GenerateTokens(userId int) (*string, *string, error) {
	access, err := s.GenerateToken(userId, s.conf.Auth.Access.ExpirationTime)
	if err != nil {
		return nil, nil, err
	}

	refresh, err := s.GenerateToken(userId, s.conf.Auth.Refresh.ExpirationTime)
	if err != nil {
		return nil, nil, err
	}

	return access, refresh, nil
}

func (s *service) getClaims(userId, expirationTime int) entities.Claims {
	return entities.Claims{
		UserID: userId,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Duration(expirationTime)*time.Second)),
		},
	}
}
