package service

import (
	"github.com/dmxmss/tasks/entities"
)

type AuthService interface {
	ValidateToken(string) (*entities.Token, error)
}

func (service *ServiceImpl) ValidateToken(rawToken string) (*entities.Token, error) {
	token, err := service.authRepo.ValidateToken(rawToken)
	if err != nil {
		return nil, err
	}

	return &entities.Token{Token: token}, nil
}
