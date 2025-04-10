package service

import (
	u "github.com/dmxmss/tasks/internal/utils"
	"github.com/dmxmss/tasks/entities"
)

type TasksService interface {
	GetAllTasks() ([]entities.GetTaskDto, error)
	CreateTask(entities.CreateTaskDto) (*entities.GetTaskDto, error)
	PatchTask(entities.PatchTaskDto) (*entities.GetTaskDto, error)
	DeleteTask(int) error
}

func (ts *ServiceImpl) GetAllTasks() ([]entities.GetTaskDto, error) {
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

func (ts *ServiceImpl) CreateTask(createTask entities.CreateTaskDto) (*entities.GetTaskDto, error) {
	task, err := ts.tasksRepo.CreateTask(createTask)
	if err != nil {
		return nil, err
	}

	result := u.ToGetTaskDto(task)

	return &result, nil
}


func (ts *ServiceImpl) PatchTask(patchTask entities.PatchTaskDto) (*entities.GetTaskDto, error) {
	task, err := ts.tasksRepo.PatchTask(patchTask)
	if err != nil {
		return nil, err
	}

	result := u.ToGetTaskDto(task)

	return &result, nil
}

func (ts *ServiceImpl) DeleteTask(id int) error {
	return ts.tasksRepo.DeleteTask(id)
}
