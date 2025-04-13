package repository

import (
	"github.com/dmxmss/tasks/entities"
	u "github.com/dmxmss/tasks/internal/utils"
	e "github.com/dmxmss/tasks/error"
	"gorm.io/gorm"

	"errors"
)

type TasksRepository interface {
	GetUserTasks(int, *entities.SearchTasksParams) ([]entities.Task, error)
	CreateTask(int, entities.CreateTaskDto) (*entities.Task, error)
	PatchTask(entities.PatchTaskDto) (*entities.Task, error)
	DeleteTask(int) error
	GetDb() *gorm.DB
}

type TasksPostgresRepo struct {
	db *gorm.DB
}

func NewTasksRepository(db *gorm.DB) TasksRepository {
	return &TasksPostgresRepo{db}
}

func (t *TasksPostgresRepo) GetUserTasks(id int, params *entities.SearchTasksParams) ([]entities.Task, error) {
	var tasks []entities.Task
	query := t.db.Where("user_id = ?", id)

	if params.Status != nil {
		query = query.Where("status = ?", *params.Status)
	}
	if params.Deadline != nil {
		query = query.Where("deadline <= ?", *params.Deadline)
	}

	if err := query.Find(&tasks).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, e.ErrUserTasksNotFound
		} else {
			return nil, e.ErrDbTransactionFailed
		}
	}

	return tasks, nil
}

func (t *TasksPostgresRepo) CreateTask(userId int, createTask entities.CreateTaskDto) (*entities.Task, error) {
	task := u.FromCreateTaskDto(createTask)
	task.UserID = userId

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

func (t *TasksPostgresRepo) PatchTask(patchTask entities.PatchTaskDto) (*entities.Task, error) {
	var task entities.Task

	if err := t.db.First(&task, patchTask.ID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, e.ErrDbTaskNotFound
		} else {
			return nil, e.ErrDbTransactionFailed
		}
	}

	if err := t.db.Model(&task).Updates(patchTask).Error; err != nil {
		return nil, e.ErrDbTransactionFailed
	}

	return &task, nil
}

func (t *TasksPostgresRepo) DeleteTask(id int) error {
	if err := t.db.Delete(&entities.Task{}, id).Error; err != nil {
		return e.ErrDbTransactionFailed
	}

	return nil
}

func (t *TasksPostgresRepo) GetDb() *gorm.DB {
	return t.db
}
