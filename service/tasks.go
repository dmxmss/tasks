package service

import (
	"github.com/dmxmss/tasks/internal"
	"github.com/dmxmss/tasks/entities"
	"github.com/dmxmss/tasks/config"
)

type TasksService interface {
	GetAllTasks() ([]entities.GetTaskDto, error)
	CreateTask(entities.CreateTaskDto) (*uint, error)
}

type TasksServiceImpl struct {
	tasksRepo internal.TasksRepository
}

func NewTasksService(conf *config.Config) (*TasksServiceImpl, error) {
	tasksRepo, err := internal.NewPgTasksRepository(conf)
	if err != nil {
		return nil, err
	}
	return &TasksServiceImpl{tasksRepo: tasksRepo}, nil
}

func (ts *TasksServiceImpl) GetAllTasks() ([]entities.GetTaskDto, error) {
	tasks, err := ts.tasksRepo.GetAllTasks()
	if err != nil {
		return nil, err
	}

	var result []entities.GetTaskDto

	for _, task := range tasks {
		result = append(result, entities.GetTaskDto{
			ID: task.ID,
			Name: task.Name,
			Description: task.Description,
			Status: task.Status,
			Deadline: task.Deadline,
			Tags: task.Tags,
			Weather: task.Weather,
			UserID: task.UserID,
		})
	}

	return result, err
}

func (ts *TasksServiceImpl) CreateTask(createTask entities.CreateTaskDto) (*uint, error) {
	task, err := ts.tasksRepo.CreateTask(createTask)
	if err != nil {
		return nil, err
	}

	return &task.ID, nil
}
