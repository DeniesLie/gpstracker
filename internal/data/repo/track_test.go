//go:build db_integration_test

package repo

import (
	"reflect"
	"testing"
	"time"

	"github.com/DeniesLie/gpstracker/config"
	"github.com/DeniesLie/gpstracker/internal/core/model"
	"github.com/DeniesLie/gpstracker/internal/core/model/enum"
	"github.com/DeniesLie/gpstracker/internal/data/db"
	"gorm.io/gorm"
)

func TestAddUpdateGetByName(t *testing.T) {
	newTrack := model.Track{}
	newTrack.Name = "NewTrack"
	newTrack.State = enum.TrackCreated
	newTrack.CreatedAt = time.Now()
	newTrack.UpdatedAt = time.Now()

	config, err := config.LoadConfig("../../../envs", "test")
	if err != nil {
		t.Errorf("Config error: %s", err)
	}

	t.Run("add update and get by name a track entity", func(t *testing.T) {
		db.UseAndDropDB(config.DBUrl, func(db *gorm.DB) {
			r := NewTrackRepo(db)

			err := r.Add(&newTrack)
			if err != nil {
				t.Errorf("failed to create track: %v", err)
			}

			newTrack.Name = "NewTrackUpdated"
			newTrack.State = enum.TrackActive
			err = r.Update(&newTrack)
			if err != nil {
				t.Errorf("failed to update track: %v", err)
			}

			track, err := r.GetByName(newTrack.Name)
			if err != nil {
				t.Errorf("failed to get track by name")
			}

			track.CreatedAt = newTrack.CreatedAt
			track.UpdatedAt = newTrack.UpdatedAt
			if !reflect.DeepEqual(newTrack, *track) {
				t.Errorf("track has been saved incorrectly; actual: %v, expected: %v", track, newTrack)
			}
		})
	})
}

func TestAddDeleteGetById(t *testing.T) {
	newTrack := model.Track{}
	newTrack.Name = "NewTrack"
	newTrack.State = enum.TrackCreated

	config, err := config.LoadConfig("../../envs", "test")
	if err != nil {
		t.Errorf("Config error: %s", err)
	}

	t.Run("add delete and get track by id", func(t *testing.T) {
		db.UseAndDropDB(config.DBUrl, func(db *gorm.DB) {
			r := NewTrackRepo(db)

			err := r.Add(&newTrack)
			if err != nil {
				t.Errorf("failed to create track: %v", err)
			}

			track, err := r.GetById(newTrack.ID)
			if err != nil {
				t.Errorf("failed to get track by id: %v", err)
			}
			if track == nil {
				t.Errorf("GetById returned nil for added track, trackId: %v", newTrack.ID)
			}

			err = r.Delete(track.ID)
			if err != nil {
				t.Errorf("failed to delete track: %v", err)
			}

			track, err = r.GetById(track.ID)
			if err != nil {
				t.Errorf("failed to get track by id: %v", err)
			}
			if track != nil {
				t.Errorf("track entry hasn't been deleted")
			}
		})
	})
}

func TestAddGetAll(t *testing.T) {
	newTrack := model.Track{}
	newTrack.Name = "NewTrack"
	newTrack.State = enum.TrackCreated

	config, err := config.LoadConfig("../../envs", "test")
	if err != nil {
		t.Errorf("Config error: %s", err)
	}

	t.Run("add track and get all", func(t *testing.T) {
		db.UseAndDropDB(config.DBUrl, func(db *gorm.DB) {
			r := NewTrackRepo(db)

			err := r.Add(&newTrack)
			if err != nil {
				t.Errorf("failed to create track: %v", err)
			}

			tracks, err := r.GetAll()
			if err != nil {
				t.Errorf("failed to get all tracks: %v", err)
			}
			if tracks == nil || len(tracks) == 0 {
				t.Errorf("get all doesn't return all entries")
			}
		})
	})
}
