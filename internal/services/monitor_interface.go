package services

import "github.com/matheusfmello/uptime-monitor/internal/models"

type MonitorServiceInterface interface {
	Create(userID uint, url string, frequencyMinutes int) (*models.Monitor, error)
	ListByUser(userID uint) ([]models.Monitor, error)
}
