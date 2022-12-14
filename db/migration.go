package db

import (
	m "city/models"

	"gorm.io/gorm"
)

func migration(db *gorm.DB) {
	models := getModels()
	db.AutoMigrate(models...)
}

func getModels() []interface{} {
	models := []interface{}{&m.Auth{}}
	return models
}
