package mikasapp

import (
	"context"
	"fmt"
	"something/aggregate"
	"something/application/contracts"
	"something/domain/entity"
	"something/infrastructure/serviceReply"
	"time"
)

type App struct {
	administration contracts.AdministrationService
	scheduler      contracts.SchedulerService
	services       contracts.ProductsService
}

func NewApp(administration contracts.AdministrationService, scheduler contracts.SchedulerService, services contracts.ProductsService) contracts.App {
	return App{
		administration: administration,
		scheduler:      scheduler,
		services:       services,
	}
}

func (app App) GetFreeSlotsForAppointments(ctx context.Context, services []string) (entity.Times, serviceReply.Reply) {
	overallDuration, sErr := app.services.GetServicesDurations(services)
	if sErr != nil {
		return entity.Times{}, sErr
	}
	availableTimeSlots := app.scheduler.GetFreeSlotsForTheWeek(overallDuration)
	return availableTimeSlots, nil
}

func (app App) GetAllAvailableServices(ctx context.Context) (aggregate.Services, serviceReply.Reply) {
	return app.services.GetAllAvailableServices()
}

func (app App) SetAppointment(ctx context.Context, user string, services []string, startTime time.Time) serviceReply.Reply {
	if availableTimeSlots, sErr := app.GetFreeSlotsForAppointments(ctx, services); sErr != nil {
		return sErr
	} else if finalDuration, finalPricing, sErr := app.services.GetCombinedServicesDetails(services); sErr != nil {
		return sErr
	} else if !availableTimeSlots.Contains(startTime) {
		return serviceReply.NewBadRequest("setAppointmentBadTime", fmt.Errorf("bad request - startTime is not available"))
	} else if _, sErr = app.scheduler.SetNewAppointment(entity.User{}, services, entity.NewTimeBlockTimings(startTime, finalDuration), finalPricing); sErr != nil {
		return sErr
	} else {
		return nil // todo need to return id
	}
}

func (app App) AddNewService(ctx context.Context, name, label string, durations entity.DurationDetails, pricing, orderBeforeMinutes float64, availableInDays int) (string, serviceReply.Reply) {
	return app.services.CreateNewService(name, label, pricing, durations)
}

func (app App) UpdateService(ctx context.Context, id, name, label string, durations entity.DurationDetails, pricing, orderBeforeMinutes float64, availableInDays int) serviceReply.Reply {
	//TODO implement me
	panic("implement me")
}

func (app App) RemoveService(ctx context.Context, id string) serviceReply.Reply {
	return app.services.RemoveService(id)
}
