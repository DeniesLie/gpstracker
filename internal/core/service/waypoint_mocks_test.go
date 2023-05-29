package service

import "github.com/DeniesLie/gpstracker/internal/core/model"

type WaypointRepoMock struct {
	GetByTrackIdResult []model.Waypoint
	GetByTrackIdError  error
	AddRangeResult     []model.Waypoint
	AddRangeError      error
}

func (r *WaypointRepoMock) GetByTrackId(trackId uint) (waypoints []model.Waypoint, err error) {
	return r.GetByTrackIdResult, r.GetByTrackIdError
}

func (r *WaypointRepoMock) AddRange(w []model.Waypoint) (waypoints []model.Waypoint, err error) {
	return r.AddRangeResult, r.AddRangeError
}
