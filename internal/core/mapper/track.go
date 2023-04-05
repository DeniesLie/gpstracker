package mapper

import (
	"github.com/DeniesLie/gpstracker/internal/core/dto"
	"github.com/DeniesLie/gpstracker/internal/core/model"
)

func ToTrackDto(from model.Track) dto.TrackGet {
	t := dto.TrackGet{
		ID:    from.Model.ID,
		Name:  from.Name,
		State: from.State.String(),
	}
	return t
}
