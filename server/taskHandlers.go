package server

import (
	"net/http"

	"github.com/dmxmss/tasks/entities"
	e "github.com/dmxmss/tasks/error"
	"github.com/gin-gonic/gin"
)

func (s *GinServer) GetUserTasks(c *gin.Context, params GetUserTasksParams) {
	claims, err := s.getClaims(c)
	if err != nil {
		c.Error(err)
		return
	}

	tasks, err := s.service.GetUserTasks(claims.UserID, &params)
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

	task, err := s.service.CreateTask(claims.UserID, claims.City, createTask)
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
	
	task, err := s.service.PatchTask(patchTask)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, task)
}

func (s *GinServer) DeleteTask(c *gin.Context, id int) {
	if err := s.service.DeleteTask(id); err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusNoContent, nil)
}
