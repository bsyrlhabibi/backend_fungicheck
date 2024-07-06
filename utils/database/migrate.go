package database

import (
	"fungicheck/module/entities"

	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) {
	err := db.AutoMigrate(
		entities.UserModels{},
		entities.ArticleModels{},
	)

	if err != nil {
		return
	}
	return
}
