package main

import (
	"log"

	"github.com/joho/godotenv"
	database "github.com/matheusfmello/uptime-monitor/internal/config"
	"github.com/matheusfmello/uptime-monitor/internal/jobs"
	"github.com/matheusfmello/uptime-monitor/internal/repositories"
	"github.com/matheusfmello/uptime-monitor/internal/services"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Failed to load .env file")
	}

	db := database.InitDb()

	repoFactory := repositories.NewRepositoryFactory(db)
	serviceFactory := services.NewServiceFactory(repoFactory)

	checker := jobs.NewMonitorChecker(serviceFactory.MonitorService)
	checker.Run()
}
