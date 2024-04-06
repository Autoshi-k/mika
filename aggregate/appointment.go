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
	appointmentTimings
	id         string
	serviceIds string
	userId     string
	price      float64
	approved   bool
	canceled   bool
}

type appointmentTimings struct {
	entity.DurationDetails
	StartTime time.Time
	EndTime   time.Time
}

func NewAppointment(userId string, serviceIds []string, startTime time.Time, durations entity.DurationDetails, pricing float64) (Appointment, error) {
	return Appointment{
		id:         uuid.NewString(),
		userId:     userId,
		serviceIds: strings.Join(serviceIds, ","),
		appointmentTimings: appointmentTimings{
			DurationDetails: durations,
			StartTime:       startTime,
			EndTime:         startTime.Add(durations.Duration),
		},
		price: pricing,
	}, nil
}
