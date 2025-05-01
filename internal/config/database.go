package database

import (
	"fmt"
	"log"
	"os"

	"github.com/matheusfmello/uptime-monitor/internal/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDb() *gorm.DB {
	dsn := buildDSN()
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Error connecting to the database")
	}

	err = db.AutoMigrate(models.AllModels...)
	if err != nil {
		log.Fatal("Error migrating models")
	}

	return db
}

func buildDSN() (dsn string) {
	var (
		HOST     string = os.Getenv("DB_HOST")
		USER     string = os.Getenv("DB_USER")
		PASSWORD string = os.Getenv("DB_PASSWORD")
		DB_NAME  string = os.Getenv("DB_NAME")
		PORT     string = os.Getenv("DB_PORT")
	)

	dsn = fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		HOST, USER, PASSWORD, DB_NAME, PORT,
	)

	return
}
