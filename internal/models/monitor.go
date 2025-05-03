package models

import "time"

type Monitor struct {
	ID               uint      `gorm:"primaryKey" json:"id"`
	UserID           uint      `json:"user_id"`
	URL              string    `json:"url"`
	FrequencyMinutes int       `json:"frequency_minutes"`
	LastCheckedAt    time.Time `json:"last_checked_at"`
	CreatedAt        time.Time
	UpdatedAt        time.Time
}
