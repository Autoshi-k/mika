package contracts

import (
	"something/aggregate"
	"something/domain/entity"
	"something/infrastructure/serviceReply"
	"time"
)

type App interface {
	GetFreeSlotsForAppointments(services []string) (entity.Times, serviceReply.Reply)
	GetAllAvailableServices() (aggregate.Services, serviceReply.Reply)
	SetAppointment(user string, services []string, startTime time.Time) serviceReply.Reply
}
