package service

import "github.com/DeniesLie/gpstracker/internal/core/model"

type (
	TrackRepo interface {
		GetAll() (tracks []model.Track, err error)
		GetById(id uint) (track *model.Track, err error)
		GetByName(name string) (t *model.Track, err error)
		Add(t *model.Track) error
		Update(t *model.Track) error
		Delete(id uint) error
	}

	WaypointRepo interface {
		GetByTrackId(trackId uint) (waypoints []model.Waypoint, err error)
		AddRange(w []model.Waypoint) (waypoints []model.Waypoint, err error)
	}
)
