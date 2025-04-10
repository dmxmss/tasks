package server

import (
	"github.com/dmxmss/tasks/config"
	"github.com/dmxmss/tasks/service"
	"github.com/gin-gonic/gin"

	"fmt"
)

type GinServer struct {
	app *gin.Engine
	conf *config.Config
	service service.Service
}

func NewGinServer(conf *config.Config) (*GinServer, error) {
	r := gin.Default()

	service, err := service.NewService(conf)
	if err != nil {
		return nil, err
	}

	return &GinServer{app: r,
										service: service,
										conf: conf}, nil
}

func (s *GinServer) RegisterHandlers() {
	s.app.Use(ErrorMiddleware())
	RegisterHandlers(s.app, s)	
}

func (s *GinServer) Start() {
	s.app.Run(fmt.Sprintf("%s:%s", s.conf.App.Address, s.conf.App.Port))
}
