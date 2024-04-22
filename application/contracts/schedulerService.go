package contracts

import (
	"something/aggregate"
	"something/domain/entity"
	"something/infrastructure/serviceReply"
	"time"
)

type SchedulerService interface {
	ApproveAppointment(id string, now time.Time) serviceReply.Reply
	CancelAppointment(id string, now time.Time) serviceReply.Reply
	GetAllAppointmentsNotApproved() []aggregate.Appointment
	GetFreeSlotsForTheWeek(duration time.Duration, now time.Time) []time.Time
	SetNewAppointment(user entity.User, services []string, timings entity.TimeBlockTimings, price float64) (string, serviceReply.Reply)
	SetNewBlocker(durations entity.DurationDetails) (string, serviceReply.Reply)
}
