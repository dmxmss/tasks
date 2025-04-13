package service

import (
	"github.com/dmxmss/tasks/internal/repository"
	"github.com/dmxmss/tasks/config"

	"gorm.io/gorm"
)

type Service interface {
	TasksService
	AuthService
	UserService
}

type service struct {
	conf *config.Config
	tasksRepo repository.TasksRepository
	authRepo repository.AuthRepository
	userRepo repository.UserRepository
	hashRepo repository.HashRepository
}

func NewService(conf *config.Config, db *gorm.DB) (Service, error) {
	tasksRepo := repository.NewTasksRepository(db)

	authRepo := repository.NewAuthRepository(conf.Auth)
	userRepo := repository.NewUserRepository(db)
	hashRepo := repository.NewHashRepository(conf.Hash)

	return &service{
		conf: conf,
		authRepo: authRepo,
		tasksRepo: tasksRepo,
		userRepo: userRepo,
		hashRepo: hashRepo,
	}, nil
}
