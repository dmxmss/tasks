package service

import (
	"github.com/dmxmss/tasks/internal/repository"
	"github.com/dmxmss/tasks/config"
)

type Service interface {
	TasksService
	AuthService
}

type ServiceImpl struct {
	tasksRepo repository.TasksRepository
	authRepo repository.AuthRepository
}

func NewService(conf *config.Config) (Service, error) {
	tasksRepo, err := repository.NewTasksRepository(conf)
	if err != nil {
		return nil, err
	}

	authRepo := repository.NewAuthRepository(conf.Auth)

	return &ServiceImpl{
		authRepo: authRepo,
		tasksRepo: tasksRepo,
	}, nil
}
