package configs

import (
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() (err error) {
	dsn, _ := os.ReadFile(os.Getenv("DSN_FILE"))
	DB, err = gorm.Open(postgres.Open(string(dsn)), &gorm.Config{})
	return
}
