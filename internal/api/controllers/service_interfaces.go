package controllers

import (
	"github.com/DeniesLie/gpstracker/internal/core/dto"
)

type (
	TrackService interface {
		GetAll() (tracks []dto.TrackGet, err error)
		Getinfo(trackId uint) (trackInfo *dto.TrackInfo, err error)
		Create(dto.TrackPost) (track *dto.TrackGet, err error)
		Complete(trackId uint) (err error)
		Delete(trackId uint) (err error)
	}

	WaypointService interface {
		GetByTrackId(trackId uint) (waypoints []dto.WaypointGet, err error)
		AddBatch([]dto.WaypointPost) (err error)
	}
)