package contracts

import (
	"context"
	"something/aggregate"
	"something/domain/entity"
	"something/infrastructure/serviceReply"
	"time"
)

type App interface {
	GetAllAvailableServices(ctx context.Context) (aggregate.Services, serviceReply.Reply)
	GetFreeSlotsForAppointments(ctx context.Context, services []string) (entity.Times, serviceReply.Reply)
	SetAppointment(ctx context.Context, user string, services []string, startTime time.Time) serviceReply.Reply
	AddNewService(ctx context.Context, name, label string, durations entity.DurationDetails, pricing, orderBeforeMinutes float64, availableInDays int) (string, serviceReply.Reply)
	UpdateService(ctx context.Context, id, name, label string, durations entity.DurationDetails, pricing, orderBeforeMinutes float64, availableInDays int) serviceReply.Reply
	RemoveService(ctx context.Context, id string) serviceReply.Reply
}
