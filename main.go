package main

import (
	"github.com/dmxmss/tasks/server"
	"github.com/dmxmss/tasks/config"
	"github.com/redis/go-redis/v9"

	"gorm.io/gorm"
	"gorm.io/driver/postgres"

	"fmt"
	"context"
)

func main() {
	conf := config.GetConfig()

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		conf.Database.Host,
		conf.Database.User,
		conf.Database.Password,
		conf.Database.Name,
		conf.Database.Port,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{TranslateError: true})	
	if err != nil {
		panic(err)
	}

	db = db.Debug()

	ctx := context.Background()
	client := redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%s", conf.Redis.Host, conf.Redis.Port),
		DB: conf.Redis.DB,
	})

	s, err := server.NewGinServer(conf, db, ctx, client)
	if err != nil {
		panic(err)
	}

	s.RegisterHandlers()

	s.Start()
}
