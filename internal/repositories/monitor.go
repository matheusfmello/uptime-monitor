package repositories

import (
	"time"

	"github.com/matheusfmello/uptime-monitor/internal/models"
	"gorm.io/gorm"
)

type MonitorRepository struct {
	db *gorm.DB
}

func NewMonitorRepository(db *gorm.DB) *MonitorRepository {
	return &MonitorRepository{db}
}

func (r *MonitorRepository) Create(monitor *models.Monitor) error {
	return r.db.Create(monitor).Error
}

func (r *MonitorRepository) FindByUserId(userId uint) ([]models.Monitor, error) {
	var monitors []models.Monitor
	err := r.db.Where("user_id = ?", userId).Find(&monitors).Error

	return monitors, err
}

func (r *MonitorRepository) GetDueMonitors() ([]models.Monitor, error) {
	var monitors []models.Monitor
	err := r.db.Where("last_checked_at IS NULL OR NOW() - last_checked_at >= (frequency_minutes || ' minutes')::interval").
		Find(&monitors).Error
	return monitors, err
}

func (r *MonitorRepository) UpdateCheckResult(id uint, status string, responseTimeMs int64) error {
	return r.db.Model(&models.Monitor{}).Where("id = ?", id).Updates(map[string]interface{}{
		"last_checked_at": time.Now(),
	}).Error
}
