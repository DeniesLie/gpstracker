package main

import (
	"log"
	"os"

	"github.com/DeniesLie/gpstracker/config"
	"github.com/DeniesLie/gpstracker/internal/api/app"
)

func main() {
	// Configuration
	env := os.Getenv("GPS_TRACKER_APP_ENV")
	cfg, err := config.LoadConfig("../../envs", env)
	if err != nil {
		log.Fatalf("Config error: %s", err)
	}

	// Run
	app.Run(cfg)
}
