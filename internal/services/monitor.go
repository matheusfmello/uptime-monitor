package services

import (
	"github.com/matheusfmello/uptime-monitor/internal/models"
	"github.com/matheusfmello/uptime-monitor/internal/repositories"
)

type MonitorService struct {
	repo *repositories.MonitorRepository
}

func NewMonitorService(repo *repositories.MonitorRepository) *MonitorService {
	return &MonitorService{repo}
}

var _ MonitorServiceInterface = &MonitorService{}

func (s *MonitorService) Create(userId uint, url string, frequencyMinutes int) (*models.Monitor, error) {
	monitor := &models.Monitor{
		UserID:   userId,
		URL:      url,
		Interval: frequencyMinutes,
	}

	err := s.repo.Create(monitor)
	if err != nil {
		return nil, err
	}

	return monitor, nil
}

func (s *MonitorService) ListByUser(userId uint) ([]models.Monitor, error) {
	return s.repo.FindByUserId(uint(userId))
}
