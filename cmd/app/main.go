package main

import (
	"log"

	"github.com/xasannosirov/online-media-service/config"
	"github.com/xasannosirov/online-media-service/internal/app"
)

func main() {
	// Configuration
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatalf("Config error: %s", err)
	}

	// Run
	app.Run(cfg)
}
