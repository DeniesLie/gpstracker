package service

import (
	"github.com/pkg/errors"

	"github.com/DeniesLie/gpstracker/internal/core/dto"
	"github.com/DeniesLie/gpstracker/internal/core/interfaces"
	"github.com/DeniesLie/gpstracker/internal/core/mapper"
	"github.com/DeniesLie/gpstracker/internal/core/model/enum"
	"github.com/DeniesLie/gpstracker/internal/core/validation/validator"
)

type WaypointService struct {
	waypointRepo interfaces.WaypointRepo
	trackRepo    interfaces.TrackRepo
}

func NewWaypointService(waypointRepo interfaces.WaypointRepo, trackRepo interfaces.TrackRepo) *WaypointService {
	return &WaypointService{waypointRepo, trackRepo}
}

func (s *WaypointService) GetByTrackId(trackId uint) (w []dto.WaypointGet, err error) {
	track, err := s.trackRepo.GetById(trackId)
	if err != nil {
		err = errors.Wrap(err, "failed at WaypointService.GetByTrackId(), some error occurred in repo")
		return
	}
	if track == nil {
		return w, NotFoundError{Resource: "Track"}
	}

	waypoints, err := s.waypointRepo.GetByTrackId(trackId)
	if err != nil {
		err = errors.Wrap(err, "failed at WaypointService.GetByTrackId(), some error occurred in repo")
		return
	}

	w = mapper.MapSlice(waypoints, mapper.ToWaypointGetDto)
	return
}

func (s *WaypointService) AddBatch(batch []dto.WaypointPost) (err error) {
	valRes, err := validator.ValidateWaypointBatchDto(batch)
	if !valRes.IsValid {
		return
	}

	track, err := s.trackRepo.GetById(batch[0].TrackID)
	if err != nil {
		err = errors.Wrap(err, "failed at WaypointService.AddBatch(), some error occurred in repo")
		return
	}
	if track == nil {
		return NotFoundError{Resource: "Track"}
	}
	if track.State == enum.TrackCompleted {
		return BusinessError{Message: "can't add waypoints to completed tracks"}
	}

	waypoints := mapper.MapSlice(batch, mapper.ToWaypointModel)
	_, err = s.waypointRepo.AddRange(waypoints)
	if err != nil {
		err = errors.Wrap(err, "failed at WaypointService.AddBatch(), some error occurred in repo")
		return
	}
	return
}
