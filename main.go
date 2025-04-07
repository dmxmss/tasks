package main

import (
	"github.com/dmxmss/tasks/server"
	"github.com/dmxmss/tasks/config"
)

func main() {
	conf := config.GetConfig()
	s, err := server.NewGinServer(conf)
	if err != nil {
		panic(err)
	}

	s.RegisterHandlers()

	s.Start()
}
