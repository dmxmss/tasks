package internal

import (
	"github.com/dmxmss/tasks/entities"
	"github.com/dmxmss/tasks/config"
	e "github.com/dmxmss/tasks/error"
	"gorm.io/gorm"
	"gorm.io/driver/postgres"

	"fmt"
	"errors"
)

type TasksRepository interface {
	GetAllTasks() ([]entities.Task, error)
	CreateTask(entities.CreateTaskDto) (*entities.Task, error)
	GetDb() *gorm.DB
}

type TasksPostgresRepo struct {
	db *gorm.DB
}

func NewPgTasksRepository(config *config.Config) (TasksRepository, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		config.Database.Host,
		config.Database.User,
		config.Database.Password,
		config.Database.Name,
		config.Database.Port,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{TranslateError: true})	
	if err != nil {
		return nil, e.ErrDbInitError
	}
	
	return &TasksPostgresRepo{db}, nil
}

func (t *TasksPostgresRepo) GetAllTasks() ([]entities.Task, error) {
	var tasks []entities.Task
	result := t.db.Find(&tasks)
	if result.Error != nil {
		return nil, e.ErrDbTransactionFailed
	}

	return tasks, nil
}

func (t *TasksPostgresRepo) CreateTask(createTask entities.CreateTaskDto) (*entities.Task, error) {
	task := entities.Task{
		Name: createTask.Name,
		Description: createTask.Description,
		Deadline: createTask.Deadline,
		Tags: createTask.Tags,
		UserID: createTask.UserID,
	}

	err := t.db.Create(&task).Error
	if err != nil {
		if errors.Is(err, gorm.ErrForeignKeyViolated) {
			return nil, e.ErrDbUserForeignKeyViolation
		} else {
			return nil, e.ErrDbTransactionFailed
		}
	}

	return &task, nil
}

func (t *TasksPostgresRepo) GetDb() *gorm.DB {
	return t.db
}
