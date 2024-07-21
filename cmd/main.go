package main

import (
	"forum/internal/server"
	"forum/pkg/config"
	"log"
)

func main() {
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatalln(err)
	}

	app := server.NewApp(*cfg)

	if err := app.Run(); err != nil {
		log.Fatalln(err)
		return
	}
}
