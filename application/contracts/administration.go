package contracts

import (
	"something/aggregate"
	"something/domain/entity"
	"something/infrastructure/serviceReply"
	"time"
)

type AdministrationService interface {
	CreateNewService(name, label string, pricing float64) (string, serviceReply.Reply)
	EditService(id, name, label string, pricing float64) serviceReply.Reply
	RemoveService(id string) serviceReply.Reply
	GetServicesDurations(serviceIds []string) (time.Duration, serviceReply.Reply)
	GetAllAvailableServices() (aggregate.Services, serviceReply.Reply)
	GetCombinedServicesDetails(serviceIds []string) (entity.DurationDetails, float64, serviceReply.Reply)
}

// todo need to figure out where to set Req structs
