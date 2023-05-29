package mapper

import (
	"github.com/DeniesLie/gpstracker/internal/core/dto"
	"github.com/DeniesLie/gpstracker/internal/core/model"
)

func ToWaypointModel(from dto.WaypointPost) (to model.Waypoint) {
	to = model.Waypoint{
		Lat:       from.Lat,
		LatHem:    from.LatHem,
		Long:      from.Long,
		LongHem:   from.LongHem,
		Timestamp: from.Time,
		TrackID:   from.TrackID,
	}
	return
}

func ToWaypointGetDto(from model.Waypoint) (to dto.WaypointGet) {
	to = dto.WaypointGet{
		ID:      from.ID,
		Lat:     from.Lat,
		LatHem:  from.LatHem,
		Long:    from.Long,
		LongHem: from.LongHem,
		Time:    from.Timestamp,
		TrackID: from.TrackID,
	}
	return
}
