package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	database "github.com/matheusfmello/uptime-monitor/internal/config"
	"github.com/matheusfmello/uptime-monitor/internal/repositories"
	"github.com/matheusfmello/uptime-monitor/internal/services"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	db := database.InitDb()

	repos := repositories.NewRepositoryFactory(db)
	services := services.NewServiceFactory(repos)

	app := fiber.New()

	SetupRoutes(app, services)

	log.Fatal(app.Listen(":" + os.Getenv("PORT")))
}
