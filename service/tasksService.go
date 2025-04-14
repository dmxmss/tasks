package service

import (
	u "github.com/dmxmss/tasks/internal/utils"
	"github.com/dmxmss/tasks/entities"
)

type TasksService interface {
	GetUserTasks(int, *entities.SearchTasksParams) ([]entities.GetTaskDto, error)
	CreateTask(int, string, entities.CreateTaskDto) (*entities.GetTaskDto, error)
	PatchTask(entities.PatchTaskDto) (*entities.GetTaskDto, error)
	DeleteTask(int) error
}

func (ts *service) GetUserTasks(id int, params *entities.SearchTasksParams) ([]entities.GetTaskDto, error) {
	tasks, err := ts.tasksRepo.GetUserTasks(id, params)
	if err != nil {
		return nil, err
	}

	var result []entities.GetTaskDto

	for _, task := range tasks {
		result = append(result, u.ToGetTaskDto(&task))
	}

	return result, err
}

func (ts *service) CreateTask(userId int, city string, createTask entities.CreateTaskDto) (*entities.GetTaskDto, error) {
	var weather string

	if city != "" {
		weatherResponse, err := ts.weatherRepo.GetCurrentWeatherAt(city)
		if err != nil {
			return nil, err
		}

		weather = weatherResponse.String()
	}

	task, err := ts.tasksRepo.CreateTask(userId, weather, createTask)
	if err != nil {
		return nil, err
	}

	result := u.ToGetTaskDto(task)

	return &result, nil
}

func (ts *service) PatchTask(patchTask entities.PatchTaskDto) (*entities.GetTaskDto, error) {
	task, err := ts.tasksRepo.PatchTask(patchTask)
	if err != nil {
		return nil, err
	}

	result := u.ToGetTaskDto(task)

	return &result, nil
}

func (ts *service) DeleteTask(id int) error {
	return ts.tasksRepo.DeleteTask(id)
}
