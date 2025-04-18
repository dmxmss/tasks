package server

import (
	"github.com/dmxmss/tasks/entities"
	e "github.com/dmxmss/tasks/error"
	"github.com/gin-gonic/gin"

	"net/http"
)

func (s *GinServer) GetUserTasks(c *gin.Context, params GetUserTasksParams) {
	claims, err := s.getClaims(c)
	if err != nil {
		c.Error(err)
		return
	}
	var p *GetUserTasksParams
	if params.Status != nil && params.Deadline != nil {
		p = &params
	}

	tasks, err := s.service.TasksService.GetUserTasks(claims.UserID, p)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, tasks)
}

func (s *GinServer) CreateTask(c *gin.Context) {
	claims, err := s.getClaims(c)
	if err != nil {
		c.Error(err)
		return
	}

	var createTask entities.CreateTaskDto

	err = c.ShouldBindJSON(&createTask)
	if err != nil {
		c.Error(e.ErrInvalidRequestBody)
		return
	}

	task, err := s.service.TasksService.CreateTask(claims.UserID, claims.City, createTask)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusCreated, task)
}

func (s *GinServer) PatchTask(c *gin.Context, id int) {
	var patchTask entities.PatchTaskDto
	patchTask.ID = id
	if err := c.ShouldBindJSON(&patchTask); err != nil {
		c.Error(e.ErrInvalidRequestBody)
		return
	}
	
	task, err := s.service.TasksService.PatchTask(patchTask)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, task)
}

func (s *GinServer) DeleteTask(c *gin.Context, id int) {
	if err := s.service.TasksService.DeleteTask(id); err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusNoContent, nil)
}
