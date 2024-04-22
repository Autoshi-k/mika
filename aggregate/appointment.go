package aggregate

import (
	"errors"
	"github.com/google/uuid"
	"something/domain/entity"
	"strings"
	"time"
)

var (
	ErrInvalidLabelx = errors.New("invalid service label - cannot be empty")
)

type Appointment struct {
	TimeBlock
	serviceIds string
	userId     string
	price      float64
	approved   time.Time
}

type TimeBlock struct {
	entity.TimeBlockTimings
	id       string
	canceled time.Time
}

func NewAppointment(userId string, serviceIds []string, startTime time.Time, durations entity.DurationDetails, pricing float64) (Appointment, error) {
	return Appointment{
		TimeBlock: TimeBlock{
			id: uuid.NewString(),
			TimeBlockTimings: entity.TimeBlockTimings{
				DurationDetails: durations,
				StartTime:       startTime,
				EndTime:         startTime.Add(durations.Duration),
			},
		},
		userId:     userId,
		serviceIds: strings.Join(serviceIds, ","),
		price:      pricing,
	}, nil
}

func (a Appointment) Approve(t time.Time) Appointment {
	a.approved = t
	return a
}

func (a Appointment) IsApproved() bool {
	return !a.approved.IsZero()
}

func (a Appointment) Cancel(t time.Time) Appointment {
	a.canceled = t
	return a
}

func (a Appointment) IsCanceled() bool {
	return !a.canceled.IsZero()
}
