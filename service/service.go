package service

import (
	"github.com/dmxmss/tasks/internal/repository"
	"github.com/dmxmss/tasks/config"

	"gorm.io/gorm"
)

type Service interface {
	TasksService
	AuthService
}

type ServiceImpl struct {
	tasksRepo repository.TasksRepository
	authRepo repository.AuthRepository
}

func NewService(conf *config.Config, db *gorm.DB) (Service, error) {
	tasksRepo, err := repository.NewTasksRepository(db)
	if err != nil {
		return nil, err
	}

	authRepo := repository.NewAuthRepository(conf.Auth)

	return &ServiceImpl{
		authRepo: authRepo,
		tasksRepo: tasksRepo,
	}, nil
}
