package configs

import (
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDB() (db *gorm.DB, err error) {
	dsn := os.Getenv("DSN")
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	return
}
