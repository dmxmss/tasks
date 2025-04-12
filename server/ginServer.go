package server

import (
	"github.com/dmxmss/tasks/config"
	"github.com/dmxmss/tasks/service"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"fmt"
)

type GinServer struct {
	app *gin.Engine
	conf *config.Config
	service service.Service
}

func NewGinServer(conf *config.Config, db *gorm.DB) (*GinServer, error) {
	r := gin.Default()

	service, err := service.NewService(conf, db)
	if err != nil {
		return nil, err
	}

	return &GinServer{app: r,
										service: service,
										conf: conf}, nil
}

func (s *GinServer) RegisterHandlers() {
	s.app.Use(ErrorMiddleware())

	RegisterHandlersWithOptions(s.app, s, GinServerOptions{
		Middlewares: []MiddlewareFunc{
			func(c *gin.Context) {
				if c.Request.URL.Path == "/auth/me" {
					s.JWTMiddleware()(c)
				}
				c.Next()
			},
		},
	})	
}

func (s *GinServer) Start() {
	s.app.Run(fmt.Sprintf("%s:%s", s.conf.App.Address, s.conf.App.Port))
}
