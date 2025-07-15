package main

import (
	"log"

	"github.com/asfung/ticus/internal/app"
	"github.com/asfung/ticus/internal/infrastructure/config"
)

func main() {
	cfg, err := config.NewAppConfig("./config/config.json")
	if err != nil {
		log.Fatal("Unable to read configuraion file", err)
	}
	application := app.NewApp(cfg)
	application.Run()
}
