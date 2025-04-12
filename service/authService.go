package service

import (
	"github.com/dmxmss/tasks/entities"
)

type AuthService interface {
	ValidateToken(string) (*entities.Claims, error)
	GenerateTokens(int) (*string, *string, error)
}

func (s*service) ValidateToken(rawToken string) (*entities.Claims, error) {
	claims, err := s.authRepo.ValidateToken(rawToken)
	if err != nil {
		return nil, err
	}

	return claims, nil
}

func (s *service) GenerateTokens(userId int) (*string, *string, error) {
	accessToken, refreshToken, err := s.authRepo.GenerateTokens(userId)
	if err != nil {
		return nil, nil, err
	}

	return accessToken, refreshToken, nil
}
