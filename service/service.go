package service

import (
	"github.com/dmxmss/tasks/internal/repository"
	"github.com/dmxmss/tasks/config"
	"github.com/redis/go-redis/v9"

	"gorm.io/gorm"
	"context"
)

type Service struct {
	conf *config.Config
	TasksService TasksService
	AuthService AuthService
	UserService UserService
}

func NewService(conf *config.Config, db *gorm.DB, ctx context.Context, redisClient *redis.Client) (*Service, error) {
	tasksRepo := repository.NewTasksRepository(db)
	authRepo := repository.NewAuthRepository(conf.Auth)
	userRepo := repository.NewUserRepository(db)
	hashRepo := repository.NewHashRepository(conf.Hash)
	weatherRepo := repository.NewWeatherRepository(conf.Weather)
	cachingRepo := repository.NewCachingRepository(ctx, redisClient)

	tasksService := NewTaskService(conf.Redis, tasksRepo, weatherRepo, cachingRepo)
	authService := NewAuthService(conf.Auth, authRepo)
	userService := NewUserService(userRepo, hashRepo)

	return &Service{
		conf: conf,
		TasksService: tasksService,
		AuthService: authService,
		UserService: userService,
	}, nil
}
