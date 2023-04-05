package repo

import (
	"github.com/DeniesLie/gpstracker/internal/core/model"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type WaypointRepo struct {
	db *gorm.DB
}

func NewWaypointRepo(db *gorm.DB) *WaypointRepo {
	return &WaypointRepo{db}
}

func (r *WaypointRepo) GetByTrackId(trackId uint) (waypoints []model.Waypoint, err error) {
	waypoints = []model.Waypoint{}
	result := r.db.Where(&model.Waypoint{TrackID: trackId}).Order("timestamp").Find(&waypoints)
	if result.Error != nil {
		err = errors.Wrap(result.Error, "failed at WaypointRepo.GetByTrackId(), some db error occurred")
	}
	return
}

func (r *WaypointRepo) AddRange(w []model.Waypoint) (waypoints []model.Waypoint, err error) {
	result := r.db.Create(w)
	if result.Error != nil {
		err = errors.Wrap(result.Error, "failed at WaypointRepo.AddRange(), some db error occurred")
	}
	return
}
