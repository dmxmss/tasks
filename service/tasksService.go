package service

import (
	u "github.com/dmxmss/tasks/internal/utils"
	"github.com/dmxmss/tasks/config"
	e "github.com/dmxmss/tasks/error"
	"github.com/dmxmss/tasks/internal/repository"
	"github.com/dmxmss/tasks/entities"

	"strconv"
	"encoding/json"
)

type TasksService interface {
	GetUserTasks(int, *entities.SearchTasksParams) ([]entities.GetTaskDto, error)
	CreateTask(int, string, entities.CreateTaskDto) (*entities.GetTaskDto, error)
	PatchTask(entities.PatchTaskDto) (*entities.GetTaskDto, error)
	DeleteTask(int) error
}

type taskService struct {
	conf *config.Redis
	tasksRepo repository.TasksRepository
	weatherRepo repository.WeatherRepository
	cachingRepo repository.CachingRepository
}

func NewTaskService(conf *config.Redis, tasksRepo repository.TasksRepository, weatherRepo repository.WeatherRepository, cachingRepo repository.CachingRepository) TasksService {
	return &taskService{
		conf: conf,
		tasksRepo: tasksRepo,
		weatherRepo: weatherRepo,
		cachingRepo: cachingRepo,
	}
}

func (ts *taskService) GetUserTasks(userId int, params *entities.SearchTasksParams) ([]entities.GetTaskDto, error) {
	if params != nil {
		return ts.GetTasksDirectly(userId, params)
	}

	tasks, err := ts.GetCachedTasks(userId)
	if err != nil {
		tasks, err = ts.GetTasksDirectly(userId, params)
		if err != nil {
			return tasks, err
		}

		id_str := strconv.Itoa(userId)
		ts.cachingRepo.SetCached(id_str, tasks, ts.conf.TaskExpiration)
		return tasks, nil
	}

	return tasks, nil
}

func (ts *taskService) GetTasksDirectly(userId int, params *entities.SearchTasksParams) ([]entities.GetTaskDto, error) {
	tasks, err := ts.tasksRepo.GetUserTasks(userId, params)
	if err != nil {
		return nil, err
	}

	var result []entities.GetTaskDto

	for _, task := range tasks {
		result = append(result, u.ToGetTaskDto(&task))
	}

	return result, err
}

func (s *taskService) GetCachedTasks(userId int) ([]entities.GetTaskDto, error) {
	id_str := strconv.Itoa(userId)

	data, err := s.cachingRepo.GetCached(id_str)
	if err != nil {
		return nil, err
	}

	var tasks []entities.GetTaskDto
	if err = json.Unmarshal(data, &tasks); err != nil {
		return nil, e.ErrJSONError
	}

	return tasks, nil
}

func (ts *taskService) CreateTask(userId int, city string, createTask entities.CreateTaskDto) (*entities.GetTaskDto, error) {
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

func (ts *taskService) PatchTask(patchTask entities.PatchTaskDto) (*entities.GetTaskDto, error) {
	task, err := ts.tasksRepo.PatchTask(patchTask)
	if err != nil {
		return nil, err
	}

	result := u.ToGetTaskDto(task)

	return &result, nil
}

func (ts *taskService) DeleteTask(id int) error {
	return ts.tasksRepo.DeleteTask(id)
}
