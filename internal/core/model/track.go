package model

import (
	"gorm.io/gorm"

	"github.com/DeniesLie/gpstracker/internal/core/model/enum"
)

type Track struct {
	gorm.Model
	Name  string
	State enum.TrackState
}
