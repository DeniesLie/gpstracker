package db

import (
	"log"

	"github.com/DeniesLie/gpstracker/internal/core/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func GetModels() []interface{} {
	return []interface{}{
		&model.Track{},
		&model.Waypoint{},
	}
}

func Connect(url string) *gorm.DB {
	db, err := gorm.Open(postgres.Open(url), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}

	err = db.AutoMigrate(GetModels()...)
	if err != nil {
		log.Fatalln(err)
	}
	return db
}

func UseAndDropDB(dbUrl string, action func(*gorm.DB)) {
	db := Connect(dbUrl)

	action(db)

	sqlDb, _ := db.DB()
	defer sqlDb.Close()

	for _, model := range GetModels() {
		err := db.Migrator().DropTable(model)
		if err != nil {
			log.Fatalf("Could not drop the table: %v", err)
		}
	}
}
