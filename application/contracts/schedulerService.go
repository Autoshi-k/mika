package contracts

import (
	"something/aggregate"
	"something/infrastructure/serviceReply"
	"time"
)

type SchedulerService interface {
	SetNewAppointment(appointment aggregate.Appointment) (string, serviceReply.Reply) // todo how do i get the service details ?
	ApproveAppointment(id string) serviceReply.Reply
	CancelAppointment(id string) serviceReply.Reply
	GetFreeSlotsForTheWeek(duration time.Duration) []time.Time
}
