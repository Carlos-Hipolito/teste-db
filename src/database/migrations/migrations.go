package migrations

import (
	"teste-db/src/models"

	"gorm.io/gorm"
)

func RunMigrations(db *gorm.DB) {
	db.AutoMigrate(&models.User{})
}
