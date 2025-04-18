package service

import (
	"github.com/dmxmss/tasks/entities"
	"github.com/dmxmss/tasks/config"
	"github.com/dmxmss/tasks/internal/repository"
	u "github.com/dmxmss/tasks/internal/utils"
)

type AuthService interface {
	ValidateToken(string) (*entities.Claims, error)
	GenerateToken(int, string, int) (*string, error)
	GenerateTokens(int, string) (*string, *string, error)
}

type authService struct {
	conf *config.Auth
	authRepo repository.AuthRepository
}

func NewAuthService(conf *config.Auth, authRepo repository.AuthRepository) AuthService {
	return &authService{
		conf: conf,
		authRepo: authRepo,
	}
}

func (s *authService) ValidateToken(rawToken string) (*entities.Claims, error) {
	claims, err := s.authRepo.ValidateToken(rawToken)
	if err != nil {
		return nil, err
	}

	return claims, nil
}

func (s *authService) GenerateToken(userId int, city string, expirationTime int) (*string, error) {
	claims := u.GetClaims(userId, city, expirationTime)
	token, err := s.authRepo.GenerateAndSignToken(claims)

	return token, err
}

func (s *authService) GenerateTokens(userId int, city string) (*string, *string, error) {
	access, err := s.GenerateToken(userId, city, s.conf.Access.ExpirationTime)
	if err != nil {
		return nil, nil, err
	}

	refresh, err := s.GenerateToken(userId, city, s.conf.Refresh.ExpirationTime)
	if err != nil {
		return nil, nil, err
	}

	return access, refresh, nil
}

