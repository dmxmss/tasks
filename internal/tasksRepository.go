package internal

import (
	"github.com/dmxmss/tasks/entities"
	e "github.com/dmxmss/tasks/error"
	"gorm.io/gorm"

	"errors"
)

type TasksRepository interface {
	GetAllTasks() ([]entities.Task, error)
}

type TasksPostgresRepo struct {
	gorm.DB
}

func (t *TasksPostgresRepo) GetAllTasks() ([]entities.Task, error) {
	var tasks []entities.Task
	result := t.Find(&tasks)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		} else {
			return nil, e.ErrDbTransactionFailed
		}
	}

	return tasks, nil
}
