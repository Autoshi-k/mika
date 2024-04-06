package services

import (
	"something/aggregate"
	"something/application/contracts"
	"something/domain/repositories/appointment"
	"something/infrastructure/serviceReply"
	"time"
)

type SchedulerService struct {
	appointments appointment.Repository
	// todo blockers blockers.Repository
}

func NewSchedulerService(appointments appointment.Repository) contracts.SchedulerService {
	return SchedulerService{appointments: appointments}
}

func (s SchedulerService) SetNewAppointment(appointment aggregate.Appointment) (string, serviceReply.Reply) {
	//TODO implement me
	panic("implement me")
}

func (s SchedulerService) ApproveAppointment(id string) serviceReply.Reply {
	//TODO implement me
	panic("implement me")
}

func (s SchedulerService) CancelAppointment(id string) serviceReply.Reply {
	//TODO implement me
	panic("implement me")
}

func (s SchedulerService) GetFreeSlots(duration time.Duration) []time.Time {
	//TODO implement me
	panic("implement me")
}

func (s SchedulerService) GetFreeSlotsForTheWeek(duration time.Duration) []time.Time {
	//TODO implement me
	panic("implement me")
}
