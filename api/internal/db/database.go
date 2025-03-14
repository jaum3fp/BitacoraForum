package db

import (
	"log"
	"os"

	"github.com/jaum3fp/bitacora-forum/internal/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func Connect() (db *gorm.DB) {
	var err error
	dsn := os.Getenv("DSN")

	logger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{LogLevel: logger.Info},
	)

	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger,
	})

	if err != nil {
		log.Fatal(err)
	}

	db.AutoMigrate(&models.Country{}, &models.User{}, &models.Tag{}, &models.Post{})

	return
}
