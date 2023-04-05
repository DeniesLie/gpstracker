package validator

import (
	"github.com/DeniesLie/gpstracker/internal/core/dto"
	"github.com/DeniesLie/gpstracker/internal/core/validation"
)

func ValidateTrackPostDto(dto *dto.TrackPost) (isValid bool, err error) {
	result := validation.Result{}
	result.Field("name").IsRequired(dto.Name).LengthIsBetween(dto.Name, 1, 100)
	if !result.IsValid {
		return false, &validation.ValidationError{Res: result}
	}
	return true, nil
}
