package server

import (
	"github.com/dmxmss/tasks/config"
	"github.com/gin-gonic/gin"

	"fmt"
)

type GinServer struct {
	app *gin.Engine
	conf *config.Config
}

func NewGinServer(conf *config.Config) Server {
	r := gin.Default()
	return &GinServer{app: r, conf: conf}
}

func (s *GinServer) Start() {
	s.app.Run(
		fmt.Sprintf("%s:%s", s.conf.App.Address, s.conf.App.Port),
	)
}
