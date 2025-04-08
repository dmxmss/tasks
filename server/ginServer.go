package server

import (
	"github.com/dmxmss/tasks/config"
	"github.com/dmxmss/tasks/service"
	e "github.com/dmxmss/tasks/error"
	"github.com/dmxmss/tasks/entities"
	"github.com/gin-gonic/gin"

	"net/http"
	"fmt"
)

type GinServer struct {
	app *gin.Engine
	conf *config.Config
	tasksService service.TasksService
}

func NewGinServer(conf *config.Config) (*GinServer, error) {
	r := gin.Default()
	tasksService, err := service.NewTasksService(conf)
	if err != nil {
		return nil, err
	}

	return &GinServer{app: r, tasksService: tasksService, conf: conf}, nil
}

func (s *GinServer) RegisterHandlers() {
	s.app.Use(ErrorMiddleware())
	RegisterHandlers(s.app, s)	
}

func (s *GinServer) Start() {
	s.app.Run(fmt.Sprintf("%s:%s", s.conf.App.Address, s.conf.App.Port))
}

func (s *GinServer) GetAllTasks(c *gin.Context) {
	tasks, err := s.tasksService.GetAllTasks()
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, tasks)
}

func (s *GinServer) CreateTask(c *gin.Context) {
	var createTask entities.CreateTaskDto
	err := c.ShouldBindJSON(&createTask)
	if err != nil {
		c.Error(e.ErrInvalidRequestBody)
		return
	}

	task, err := s.tasksService.CreateTask(createTask)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusCreated, task)
}
