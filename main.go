package main

import (
	"github.com/dmxmss/tasks/server"
	"github.com/dmxmss/tasks/config"
)

func main() {
	conf := config.GetConfig()
	s := server.NewGinServer(conf)

	s.Start()
}
