package main

import (
	"log"

	"github.com/DeniesLie/gpstracker/config"
	"github.com/DeniesLie/gpstracker/internal/api/app"
)

func main() {
	// Configuration
	cfg, err := config.LoadConfig("../../envs")
	if err != nil {
		log.Fatalf("Config error: %s", err)
	}

	// Run
	app.Run(cfg)
}
