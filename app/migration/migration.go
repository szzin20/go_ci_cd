package migration

import (
	"docker/models"

	"gorm.io/gorm"
)

func InitMigrationMysql(db *gorm.DB) {
	db.AutoMigrate(&models.User{})
	// auto migrate untuk table book
}
