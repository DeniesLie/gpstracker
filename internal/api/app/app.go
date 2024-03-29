package app

import (
	"log"

	"github.com/DeniesLie/gpstracker/config"
	"github.com/DeniesLie/gpstracker/internal/api/controllers"
	"github.com/DeniesLie/gpstracker/internal/api/middleware"
	"github.com/DeniesLie/gpstracker/internal/core/service"
	"github.com/DeniesLie/gpstracker/internal/data/db"
	"github.com/DeniesLie/gpstracker/internal/data/repo"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(
	trackSvc controllers.TrackService,
	waypointSvc controllers.WaypointService) *gin.Engine {

	router := gin.Default()
	router.Use(middleware.ErrorHandler)

	controllers.AddTrackRoutes(router, trackSvc)
	controllers.AddWaypointRoutes(router, waypointSvc)

	return router
}

func Run(cfg *config.Config) {
	// Create repos
	db := db.Connect(cfg.DBUrl)
	trackRepo := repo.NewTrackRepo(db)
	waypointRepo := repo.NewWaypointRepo(db)

	// Create services
	trackSvc := service.NewTrackService(trackRepo, waypointRepo)
	waypointSvc := service.NewWaypointService(waypointRepo, trackRepo)

	router := RegisterRoutes(trackSvc, waypointSvc)

	err := router.Run(cfg.Port)
	if err != nil {
		log.Fatalln(err)
	}
}
