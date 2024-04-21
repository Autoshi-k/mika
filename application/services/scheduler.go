package services

import (
	"something/aggregate"
	"something/application/contracts"
	"something/domain/entity"
	"something/domain/repositories/appointment"
	"something/domain/repositories/blocker"
	"something/domain/repositories/timetable"
	"something/infrastructure/serviceReply"
	"time"
)

type SchedulerService struct {
	appointments appointment.Repository
	blockers     blocker.Repository
	timetable    timetable.Repository
}

func NewSchedulerService(appointments appointment.Repository, blockers blocker.Repository, timetable timetable.Repository) contracts.SchedulerService {
	return SchedulerService{appointments: appointments, blockers: blockers, timetable: timetable}
}

func (s SchedulerService) ApproveAppointment(id string) serviceReply.Reply {
	//TODO implement me
	panic("implement me")
}

func (s SchedulerService) CancelAppointment(id string) serviceReply.Reply {
	//TODO implement me
	panic("implement me")
}

func (s SchedulerService) GetAllAppointmentsNotApproved() []aggregate.Appointment {
	//TODO implement me
	panic("implement me")
}

func (s SchedulerService) GetFreeSlotsForTheWeek(duration time.Duration) []time.Time {
	//TODO implement me
	panic("implement me")
}

func (s SchedulerService) SetNewAppointment(user entity.User, services []string, timings entity.TimeBlockTimings, price float64) (string, serviceReply.Reply) {
	//TODO implement me
	panic("implement me")
}

func (s SchedulerService) SetNewBlocker(durations entity.DurationDetails) (string, serviceReply.Reply) {
	//TODO implement me
	panic("implement me")
}
