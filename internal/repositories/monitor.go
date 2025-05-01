package repositories

import (
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
