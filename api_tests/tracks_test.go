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

func TestTrackEndpoints(t *testing.T) {
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
			Body(`{ "isSuccess": true,
					"message": "Success",
					"data": {
						"id": 1,
						"name": "test track #1",
						"state": "Active"
					},
					"validationResult": "" }`).
			Status(http.StatusCreated).
			End()

		apitest.New("create invalid track returns 400 Bad Request").
			Handler(router).
			Post("/tracks").
			Body(`{ "name": "" }`).
			Expect(t).
			Status(http.StatusBadRequest).
			End()

		// Track info
		apitest.New("get info returns 200 OK for existing track").
			Handler(router).
			Get("/tracks/1/info").
			Expect(t).
			Body(`{ "isSuccess": true,
					"message": "Success",
					"data": {
						"trackId": 1,
						"trackName": "test track #1",
						"state": "Active",
						"totalDistanceMeters": 0,
						"averageSpeedMps": 0
					},
					"validationResult": "" }`).
			Status(http.StatusOK).
			End()

		apitest.New("get info returns 404 Not Found for non-existing track").
			Handler(router).
			Get("/tracks/11111/info").
			Expect(t).
			Status(http.StatusNotFound).
			End()

		// Complete track
		apitest.New("complete returns 200 OK for existing track").
			Handler(router).
			Post("/tracks/1/complete").
			Expect(t).
			Status(http.StatusOK).
			End()

		apitest.New("complete returns 404 Not Found for non-existing track").
			Handler(router).
			Post("/tracks/111111/complete").
			Expect(t).
			Status(http.StatusNotFound).
			End()

		// Get all
		apitest.New("get all return 200 OK and list of tracks").
			Handler(router).
			Get("/tracks").
			Expect(t).
			Body(`{ "isSuccess": true,
					"message": "Success",
					"data": [
						{
							"id": 1,
							"name": "test track #1",
							"state": "Completed"
						}
					],
					"validationResult": "" }`).
			Status(http.StatusOK).
			End()

		// delete track
		apitest.New("delete returns 200 OK for existing track").
			Handler(router).
			Delete("/tracks/1").
			Expect(t).
			Status(http.StatusOK).
			End()

		apitest.New("delete returns 404 Not Found for existing track").
			Handler(router).
			Delete("/tracks/111111").
			Expect(t).
			Status(http.StatusNotFound).
			End()
	})
}
