//go:build api_integration_test

package api_tests

import (
	"net/http"
	"testing"

	"github.com/DeniesLie/gpstracker/config"
	"github.com/DeniesLie/gpstracker/internal/api/app"
	"github.com/DeniesLie/gpstracker/internal/core/service"
	"github.com/DeniesLie/gpstracker/internal/data/db"
	"github.com/DeniesLie/gpstracker/internal/data/repo"
	"github.com/steinfletcher/apitest"
	"gorm.io/gorm"
)

func TestWaypointEndpoints(t *testing.T) {
	config, err := config.LoadConfig("../envs", "test")
	if err != nil {
		t.Errorf("Config error: %s", err)
	}

	db.UseAndDropDB(config.DBUrl, func(db *gorm.DB) {
		// init dependencies
		trackRepo := repo.NewTrackRepo(db)
		waypointsRepo := repo.NewWaypointRepo(db)
		trackSvc := service.NewTrackService(trackRepo, waypointsRepo)
		waypointSvc := service.NewWaypointService(waypointsRepo, trackRepo)

		// register routes
		router := app.RegisterRoutes(trackSvc, waypointSvc)

		// Create track
		apitest.New("create valid track returns 200 OK").
			Handler(router).
			Post("/tracks").
			Body(`{ "name": "test track #1" }`).
			Expect(t).
			Status(http.StatusCreated).
			End()

		// Add waypoints
		apitest.New("create valid waypoints returns 201 Created").
			Handler(router).
			Post("/waypoints/addBatch").
			Body(`[
					{
						"trackId": 1,
						"lat": 50.540586,
						"latHem": "N",
						"long": 30.244127,
						"longHem": "E",
						"time": 1680655436897
					},
					{
						"trackId": 1,
						"lat": 50.540694,
						"latHem": "N",
						"long": 30.244624,
						"longHem": "E",
						"time": 1680655437897
					},
					{
						"trackId": 1,
						"lat": 50.540799,
						"latHem": "N",
						"long": 30.244918,
						"longHem": "E",
						"time": 1680655438897
					}
				]`).
			Expect(t).
			Status(http.StatusCreated).
			End()

		apitest.New("create invalid waypoints returns 400 Bad Request").
			Handler(router).
			Post("/waypoints/addBatch").
			Body(`[
					{
						"trackId": 1,
						"lat": 50.540586,
						"latHem": "E",
						"long": 30.244127,
						"longHem": "J",
						"time": 1680655436897
					}
				]`).
			Expect(t).
			Status(http.StatusBadRequest).
			End()

		apitest.New("get by track id returns 200 OK for existing track").
			Handler(router).
			Get("/waypoints/1").
			Expect(t).
			Status(http.StatusOK).
			End()

		apitest.New("get by track id returns 404 Not Found for non-existing track").
			Handler(router).
			Get("/waypoints/111111").
			Expect(t).
			Status(http.StatusNotFound).
			End()
	})
}
