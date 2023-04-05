package validator

import (
	"github.com/DeniesLie/gpstracker/internal/core/dto"
	"github.com/DeniesLie/gpstracker/internal/core/validation"
)

func ValidateWaypointPostDto(dto *dto.WaypointPost) (res validation.Result, err error) {
	res.IsValid = true
	res.Field("trackId").IsMoreThan(dto.TrackID, 0)
	res.Field("lat").ValueIsBetween(dto.Lat, 0, 90)
	res.Field("latHem").HasValues(dto.LatHem, []string{"S", "N"})
	res.Field("long").ValueIsBetween(dto.Long, 0, 180)
	res.Field("longHem").HasValues(dto.LongHem, []string{"W", "E"})
	if !res.IsValid {
		return res, validation.ValidationError{Res: res}
	}
	return res, nil
}

func ValidateWaypointBatchDto(waypoints []dto.WaypointPost) (res validation.Result, err error) {
	res.IsValid = true
	if len(waypoints) == 0 {
		res.Messages = append(res.Messages, "batch must contain at least one waypoint")
		return res, validation.ValidationError{Res: res}
	}
	trackId := waypoints[0].TrackID
	for _, w := range waypoints {
		r, _ := ValidateWaypointPostDto(&w)
		if !r.IsValid {
			res = validation.AggregateResult(res, r)
		}
		if w.TrackID != trackId {
			res.IsValid = false
			res.Messages = append(res.Messages, "trackId must be same for all waypoints")
		}
	}
	if !res.IsValid {
		return res, validation.ValidationError{Res: res}
	}
	return
}
