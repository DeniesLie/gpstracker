//go:build db_integration_test

package repo

import (
	"testing"
	"time"

	"github.com/DeniesLie/gpstracker/config"
	"github.com/DeniesLie/gpstracker/internal/core/model"
	"github.com/DeniesLie/gpstracker/internal/data/db"
	"gorm.io/gorm"
)

func TestAddRangeGetByTrackId(t *testing.T) {
	track := model.Track{}
	track.Name = "NewTrack"

	config, err := config.LoadConfig("../../../envs", "test")
	if err != nil {
		t.Errorf("Config error: %s", err)
	}

	t.Run("add range and get by id", func(t *testing.T) {
		db.UseAndDropDB(config.DBUrl, func(db *gorm.DB) {
			trackRepo := NewTrackRepo(db)
			waypointRepo := NewWaypointRepo(db)

			err := trackRepo.Add(&track)
			if err != nil {
				t.Errorf("failed to add track: %v", err)
			}

			waypoints := []model.Waypoint{
				{
					Lat:       0.0,
					LatHem:    "W",
					Long:      0.0,
					LongHem:   "E",
					Timestamp: time.Now().UnixMilli(),
					TrackID:   track.ID,
				},
			}

			waypoints, err = waypointRepo.AddRange(waypoints)
			if err != nil {
				t.Errorf("failed to add range of waypoints: %v", err)
			}

			waypoints, err = waypointRepo.GetByTrackId(track.ID)
			if err != nil {
				t.Errorf("failed to get waypoints by track id: %v", err)
			}

			if len(waypoints) == 0 {
				t.Errorf("added entries were not found in db")
			}
		})
	})
}
