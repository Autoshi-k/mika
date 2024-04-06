package mikasapp

import (
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
}

func NewApp(administration contracts.AdministrationService, scheduler contracts.SchedulerService) contracts.App {
	return App{
		administration: administration,
		scheduler:      scheduler,
	}
}

func (app App) GetFreeSlotsForAppointments(services []string) (entity.Times, serviceReply.Reply) {
	overallDuration, sErr := app.administration.GetServicesDurations(services)
	if sErr != nil {
		return entity.Times{}, sErr
	}
	availableTimeSlots := app.scheduler.GetFreeSlotsForTheWeek(overallDuration)
	return availableTimeSlots, nil
}

func (app App) GetAllAvailableServices() (aggregate.Services, serviceReply.Reply) {
	return app.administration.GetAllAvailableServices()
}

func (app App) SetAppointment(user string, services []string, startTime time.Time) serviceReply.Reply {
	if availableTimeSlots, sErr := app.GetFreeSlotsForAppointments(services); sErr != nil {
		return sErr
	} else if finalDuration, finalPricing, sErr := app.administration.GetCombinedServicesDetails(services); sErr != nil {
		return sErr
	} else if !availableTimeSlots.Contains(startTime) {
		return serviceReply.NewBadRequest("setAppointmentBadTime", fmt.Errorf("bad request - startTime is not available"))
	} else {
		appointment, _ := aggregate.NewAppointment(user, services, startTime, finalDuration, finalPricing)
		app.scheduler.SetNewAppointment(appointment)
		return nil
	}
}
