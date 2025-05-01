package repositories

import "gorm.io/gorm"

type RepositoryFactory struct {
	UserRepo    *UserRepository
	MonitorRepo *MonitorRepository
}

func NewRepositoryFactory(db *gorm.DB) *RepositoryFactory {
	return &RepositoryFactory{
		UserRepo:    NewUserRepository(db),
		MonitorRepo: NewMonitorRepository(db),
	}
}
