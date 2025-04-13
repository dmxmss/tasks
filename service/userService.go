package service

import (
	"github.com/dmxmss/tasks/entities"
	e "github.com/dmxmss/tasks/error"
)

type UserService interface {
	CreateUser(entities.CreateUserDto) (*entities.GetUserDto, error)
	LogIn(entities.LoginUserDto) (*entities.GetUserDto, error)
	GetUserInfo(int) (*entities.GetUserDto, error)
}

func (s *service) CreateUser(createUser entities.CreateUserDto) (*entities.GetUserDto, error) {
	hashedPassword, err := s.hashRepo.HashPassword(createUser.Password)
	if err != nil {
		return nil, err
	}

	createUser.Password = hashedPassword
	user, err := s.userRepo.CreateUser(createUser)
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

func (s *service) LogIn(login entities.LoginUserDto) (*entities.GetUserDto, error) {
	user, err := s.userRepo.GetUserBy(entities.SearchUserDto{Email: &login.Email})
	if err != nil {
		return nil, err
	}

	if !s.hashRepo.IsPasswordValid(login.Password, user.Password) {
		return nil, e.ErrAuthInvalidCredentials
	}

	return &entities.GetUserDto{
		ID: user.ID,
		FullName: user.FullName,
		Email: user.Email,
	}, nil
}

func (s *service) GetUserInfo(userId int) (*entities.GetUserDto, error) {
	user, err := s.userRepo.GetUserBy(entities.SearchUserDto{ID: &userId})
	if err != nil {
		return nil, err
	}

	return &entities.GetUserDto{
		ID: user.ID,
		FullName: user.FullName,
		Email: user.Email,
	}, nil
}
