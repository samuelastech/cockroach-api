package main

import (
	"fmt"
	"github.com/samulastech/cockroach/config"
	"github.com/samulastech/cockroach/server"
	"log"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	configs := config.GetConfig()
	app := server.NewChiServer(configs)
	log.Println(fmt.Sprintf("[msg: server is running][port:%s]", configs.Server.Port))
	return app.Start()
}
