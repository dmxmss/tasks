package service

import (
	"github.com/dmxmss/tasks/entities"
	u "github.com/dmxmss/tasks/internal/utils"
)

type AuthService interface {
	ValidateToken(string) (*entities.Claims, error)
	GenerateToken(int, string, int) (*string, error)
	GenerateTokens(int, string) (*string, *string, error)
}

func (s *service) ValidateToken(rawToken string) (*entities.Claims, error) {
	claims, err := s.authRepo.ValidateToken(rawToken)
	if err != nil {
		return nil, err
	}

	return claims, nil
}

func (s *service) GenerateToken(userId int, city string, expirationTime int) (*string, error) {
	claims := u.GetClaims(userId, city, expirationTime)
	token, err := s.authRepo.GenerateAndSignToken(claims)

	return token, err
}

func (s *service) GenerateTokens(userId int, city string) (*string, *string, error) {
	access, err := s.GenerateToken(userId, city, s.conf.Auth.Access.ExpirationTime)
	if err != nil {
		return nil, nil, err
	}

	refresh, err := s.GenerateToken(userId, city, s.conf.Auth.Refresh.ExpirationTime)
	if err != nil {
		return nil, nil, err
	}

	return access, refresh, nil
}

