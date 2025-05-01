package services

import (
	"github.com/matheusfmello/uptime-monitor/internal/repositories"
)

type ServiceFactory struct {
	UserService    *UserService
	MonitorService *MonitorService
}

func NewServiceFactory(repoFactory *repositories.RepositoryFactory) *ServiceFactory {
	return &ServiceFactory{
		UserService:    NewUserService(repoFactory.UserRepo),
		MonitorService: NewMonitorService(repoFactory.MonitorRepo),
	}
}
