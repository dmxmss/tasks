package service

import (
	"github.com/dmxmss/tasks/internal"
	"github.com/dmxmss/tasks/entities"
)

type TasksService interface {
	GetAllTasks() ([]entities.GetTaskDto, error)
}

type TasksServiceImpl struct {
	tasksRepo internal.TasksRepository
}

func (ts *TasksServiceImpl) GetAllTasks() ([]entities.GetTaskDto, error) {
	tasks, err := ts.tasksRepo.GetAllTasks()
	if err != nil {
		return nil, err
	}

	var result []entities.GetTaskDto

	for _, task := range tasks {
		result = append(result, entities.GetTaskDto{
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
