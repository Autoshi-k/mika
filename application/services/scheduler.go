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

func (s SchedulerService) ApproveAppointment(id string, now time.Time) serviceReply.Reply {
	if ap, err := s.appointments.Get(id); err != nil {
		return serviceReply.NewDbError(err)
	} else if err = s.appointments.Update(id, ap.Approve(now)); err != nil {
		return serviceReply.NewDbError(err)
	} else {
		return nil
	}
}

func (s SchedulerService) CancelAppointment(id string, now time.Time) serviceReply.Reply {
	if ap, err := s.appointments.Get(id); err != nil {
		return serviceReply.NewDbError(err)
	} else if err = s.appointments.Update(id, ap.Cancel(now)); err != nil {
		return serviceReply.NewDbError(err)
	} else {
		return nil
	}
}

func (s SchedulerService) GetAllAppointmentsNotApproved() []aggregate.Appointment {
	var notApprovedAppointments []aggregate.Appointment

	appointments, _ := s.appointments.GetAll() // err ?
	for _, ap := range appointments {
		if !ap.IsApproved() && !ap.IsCanceled() {
			notApprovedAppointments = append(notApprovedAppointments, ap)
		}
	}

	return notApprovedAppointments
}

func (s SchedulerService) GetFreeSlotsForTheWeek(workday entity.WorkDay, duration time.Duration, now time.Time) (map[int][]time.Time, serviceReply.Reply) {
	if week, err := s.timetable.GetWeek(now); err != nil {
		return nil, serviceReply.NewDbError(err)
	} else if blockers, err := s.blockers.GetWeek(now); err != nil {
		return nil, serviceReply.NewDbError(err)
	} else {
		var freeSlots map[int][]time.Time
		for i, day := range week {
			freeSlots[i] = workday.Reduce(day.AllEventsBlockDurations(), blockers[i]).ReduceByDuration(duration).AvailableTimeSlots()
		}
	}
}

func (s SchedulerService) SetNewAppointment(user entity.User, services []string, timings entity.TimeBlockTimings, price float64) (string, serviceReply.Reply) {
	//TODO implement me
	panic("implement me")
}

func (s SchedulerService) SetNewBlocker(durations entity.DurationDetails) (string, serviceReply.Reply) {
	//TODO implement me
	panic("implement me")
}
