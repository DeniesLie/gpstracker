package db

import (
	"log"

	"github.com/DeniesLie/gpstracker/internal/core/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connect(url string) *gorm.DB {
	db, err := gorm.Open(postgres.Open(url), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}

	db.AutoMigrate(&model.Track{}, &model.Waypoint{})

	return db
}
