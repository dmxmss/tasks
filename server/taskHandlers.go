package server

import (
	"github.com/dmxmss/tasks/entities"
	e "github.com/dmxmss/tasks/error"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (s *GinServer) GetAllTasks(c *gin.Context) {
	tasks, err := s.service.GetAllTasks()
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

	task, err := s.service.CreateTask(createTask)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusCreated, task)
}

func (s *GinServer) PatchTask(c *gin.Context, id int) {
	var patchTask entities.PatchTaskDto
	if err := c.ShouldBindJSON(&patchTask); err != nil {
		c.Error(e.ErrInvalidRequestBody)
		return
	}
	patchTask.ID = id
	
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
