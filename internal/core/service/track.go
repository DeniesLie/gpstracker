package service

import (
	"math"

	"github.com/pkg/errors"

	"github.com/DeniesLie/gpstracker/internal/core/dto"
	"github.com/DeniesLie/gpstracker/internal/core/interfaces"
	"github.com/DeniesLie/gpstracker/internal/core/mapper"
	"github.com/DeniesLie/gpstracker/internal/core/model"
	"github.com/DeniesLie/gpstracker/internal/core/model/enum"
	"github.com/DeniesLie/gpstracker/internal/core/validation/validator"
)

type TrackService struct {
	trackRepo    interfaces.TrackRepo
	waypointRepo interfaces.WaypointRepo
}

func NewTrackService(trackRepo interfaces.TrackRepo, waypointRepo interfaces.WaypointRepo) *TrackService {
	return &TrackService{trackRepo, waypointRepo}
}

func (s *TrackService) GetAll() (t []dto.TrackGet, err error) {
	var tracks []model.Track
	tracks, err = s.trackRepo.GetAll()
	if err != nil {
		err = errors.Wrap(err, "failed at TrackService.GetAll(), some error occurred in repo")
		return
	}
	t = mapper.MapSlice(tracks, mapper.ToTrackDto)
	return
}

func (s *TrackService) Getinfo(trackId uint) (i *dto.TrackInfo, err error) {
	track, err := s.trackRepo.GetById(trackId)
	if err != nil {
		err = errors.Wrap(err, "failed at TrackService.Getinfo(), some error occurred in repo")
		return
	}
	if track == nil {
		return nil, NotFoundError{}
	}

	var waypoints []model.Waypoint
	waypoints, err = s.waypointRepo.GetByTrackId(trackId)
	if err != nil {
		err = errors.Wrap(err, "failed at TrackService.Getinfo(), some error occurred in repo")
		return
	}

	i = &dto.TrackInfo{}
	i.ID = trackId
	i.State = track.State.String()
	i.Name = track.Name
	i.TotalDistanceMtrs = getTotalDistanceMtrs(waypoints)
	i.AverageSpeedMps = getAverageSpeedMps(i.TotalDistanceMtrs, waypoints)
	return
}

func (s *TrackService) Create(trackDto dto.TrackPost) (t *dto.TrackGet, err error) {
	isValid, err := validator.ValidateTrackPostDto(&trackDto)
	if !isValid {
		return
	}

	existingTrack, err := s.trackRepo.GetByName(trackDto.Name)
	if err != nil {
		err = errors.Wrap(err, "failed at TrackService.Create(), some error occurred in repo")
		return
	}
	if existingTrack != nil {
		return nil, BusinessError{Message: "Case name is taken, try another name"}
	}

	track := model.Track{
		Name:  trackDto.Name,
		State: enum.TrackActive,
	}
	err = s.trackRepo.Add(&track)
	if err != nil {
		err = errors.Wrap(err, "failed at TrackService.Create(), some error occurred in repo")
		return
	}

	res := mapper.ToTrackDto(track)
	return &res, nil
}

func (s *TrackService) Complete(trackId uint) (err error) {
	track, err := s.trackRepo.GetById(trackId)
	if err != nil {
		err = errors.Wrap(err, "failed at TrackService.Complete(), some error occurred in repo")
		return
	}
	if track == nil {
		return NotFoundError{}
	}

	track.State = enum.TrackCompleted
	err = s.trackRepo.Update(track)
	if err != nil {
		err = errors.Wrap(err, "failed at TrackService.Complete(), some error occurred in repo")
	}
	return
}

func (s *TrackService) Delete(trackId uint) (err error) {
	track, err := s.trackRepo.GetById(trackId)
	if err != nil {
		err = errors.Wrap(err, "failed at TrackService.Delete(), some error occurred in repo")
		return
	}
	if track == nil {
		return NotFoundError{}
	}

	err = s.trackRepo.Delete(trackId)
	if err != nil {
		err = errors.Wrap(err, "failed at TrackService.Delete(), some error occurred in repo")
	}
	return
}

func getTotalDistanceMtrs(w []model.Waypoint) float64 {
	if len(w) <= 1 {
		return 0
	}
	var totalDistance float64
	for i := 0; i < len(w)-1; i++ {
		distance := haversine(w[i].Lat, w[i].Long, w[i+1].Lat, w[i+1].Long)
		totalDistance += distance
	}
	return totalDistance
}

func haversine(lat1, long1, lat2, long2 float64) float64 {
	const R = 6371000 // Earth radius in meters
	phi1 := lat1 * math.Pi / 180
	phi2 := lat2 * math.Pi / 180
	deltaPhi := (lat2 - lat1) * math.Pi / 180
	deltaLambda := (long2 - long1) * math.Pi / 180

	a := math.Sin(deltaPhi/2)*math.Sin(deltaPhi/2) + math.Cos(phi1)*math.Cos(phi2)*math.Sin(deltaLambda/2)*math.Sin(deltaLambda/2)
	c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))
	d := R * c

	return d
}

func getAverageSpeedMps(distance float64, w []model.Waypoint) float64 {
	if len(w) <= 1 {
		return 0
	}
	startedAt := w[0].Timestamp
	latestTime := w[len(w)-1].Timestamp
	durationSec := float64(latestTime-startedAt) / 1000

	return distance / durationSec
}
