package models

import "time"

type Monitor struct {
	ID            uint      `gorm:"primaryKey" json:"id"`
	Name          string    `json:"name" validate:"required"`
	URL           string    `json:"url" validate:"required,url"`
	Interval      int       `json:"interval" validate:"required,min=30"` // seconds
	Timeout       int       `json:"timeout" validate:"required,min=1"`   // seconds
	Status        string    `json:"status,omitempty"`
	LastCheckedAt time.Time `json:"last_checked_at,omitempty"`
	UserID        uint      `json:"user_id"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}
