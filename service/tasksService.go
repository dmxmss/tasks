package service

import (
	"github.com/dmxmss/tasks/internal/repository"
	u "github.com/dmxmss/tasks/internal/utils"
	"github.com/dmxmss/tasks/entities"
	"github.com/dmxmss/tasks/config"
)

type TasksService interface {
	GetAllTasks() ([]entities.GetTaskDto, error)
	CreateTask(entities.CreateTaskDto) (*entities.GetTaskDto, error)
	PatchTask(entities.PatchTaskDto) (*entities.GetTaskDto, error)
	DeleteTask(int) error
}

type TasksServiceImpl struct {
	tasksRepo repository.TasksRepository
}

func NewTasksService(conf *config.Config) (*TasksServiceImpl, error) {
	tasksRepo, err := repository.NewPgTasksRepository(conf)
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
		result = append(result, u.ToGetTaskDto(&task))
	}

	return result, err
}

func (ts *TasksServiceImpl) CreateTask(createTask entities.CreateTaskDto) (*entities.GetTaskDto, error) {
	task, err := ts.tasksRepo.CreateTask(createTask)
	if err != nil {
		return nil, err
	}

	result := u.ToGetTaskDto(task)

	return &result, nil
}


func (ts *TasksServiceImpl) PatchTask(patchTask entities.PatchTaskDto) (*entities.GetTaskDto, error) {
	task, err := ts.tasksRepo.PatchTask(patchTask)
	if err != nil {
		return nil, err
	}

	result := u.ToGetTaskDto(task)

	return &result, nil
}

func (ts *TasksServiceImpl) DeleteTask(id int) error {
	return ts.tasksRepo.DeleteTask(id)
}
