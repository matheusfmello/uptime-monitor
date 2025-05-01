package services

import (
	"github.com/matheusfmello/uptime-monitor/internal/models"
)

type UserServiceInterface interface {
	Register(name string, email string, password string) (*models.User, error)
	Authenticate(email string, password string) (*models.User, error)
}
