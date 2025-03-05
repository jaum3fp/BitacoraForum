package db

import (
	"log"
	"os"

	"github.com/jaum3fp/bitacora-forum/internal/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)


var instance *gorm.DB


func Init() {
	var err error
	dsn := os.Getenv("DSN")
  	instance, err = gorm.Open(mysql.Open(dsn))
	if err != nil {
		log.Fatal(err)
	}
	instance.AutoMigrate(&models.Country{}, &models.User{}, &models.Post{})
}

func GetInstance() *gorm.DB {
	if instance == nil {
		log.Fatal("Inicializa la base de datos con Init() antes de usar su instancia!")
	}
	return instance
}
