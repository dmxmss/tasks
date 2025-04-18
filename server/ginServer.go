package server

import (
	"github.com/dmxmss/tasks/config"
	"github.com/dmxmss/tasks/service"
	e "github.com/dmxmss/tasks/error"
	"github.com/redis/go-redis/v9"
	"github.com/dmxmss/tasks/entities"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"fmt"
	"strings"
	"log"
	"context"
)

type GetUserTasksParams = entities.SearchTasksParams

type GinServer struct {
	app *gin.Engine
	conf *config.Config
	service *service.Service
}

func NewGinServer(conf *config.Config, db *gorm.DB, ctx context.Context, redisClient *redis.Client) (*GinServer, error) {
	r := gin.Default()

	service, err := service.NewService(conf, db, ctx, redisClient)
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
				path := c.Request.URL.Path
				if path == "/auth/me" || strings.HasPrefix(path, "/tasks") {
					s.JWTAccessMiddleware()(c)
				} else if path == "/auth/refresh" {
					s.JWTRefreshMiddleware()(c)
				}
				c.Next()
			},
		},
	})	
}

func (s *GinServer) Start() {
	s.app.Run(fmt.Sprintf("%s:%s", s.conf.App.Address, s.conf.App.Port))
}

func (s *GinServer) getClaims(c *gin.Context) (*entities.Claims, error) {
	v, exists := c.Get("claims")
	claims, ok := v.(*entities.Claims)

	if !ok || !exists {
		log.Printf("Error: missing claims")
		return nil, e.ErrAuthFailed
	}
	return claims, nil
}
