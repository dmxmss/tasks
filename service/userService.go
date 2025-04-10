package service

import (
	"github.com/dmxmss/tasks/entities"
)

type UserService interface {
	CreateUser(entities.CreateUserDto) (*entities.GetUserDto, error)
}

func (us *service) CreateUser(createUser entities.CreateUserDto) (*entities.GetUserDto, error) {
	user, err := us.userRepo.CreateUser(createUser)
	if err != nil {
		return nil, err
	}

	return &entities.GetUserDto{
		ID: user.ID,
		FullName: user.FullName,
		Email: user.Email,
		City: user.City,
	}, nil
}
