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
		UserID:           userId,
		URL:              url,
		FrequencyMinutes: frequencyMinutes,
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

func (s *MonitorService) GetDueMonitors() ([]MonitorInfo, error) {
	monitors, err := s.repo.GetDueMonitors()
	if err != nil {
		return nil, err
	}
	resp := make([]MonitorInfo, len(monitors))
	for i, monitor := range monitors {
		resp[i] = MonitorInfo{
			ID:  monitor.ID,
			URL: monitor.URL,
		}
	}

	return resp, nil
}

func (s *MonitorService) SaveCheckResult(id uint, status string, responseTime int64) error {
	return s.repo.UpdateCheckResult(id, status, responseTime)
}
