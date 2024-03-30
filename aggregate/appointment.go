package aggregate

import (
	"errors"
	"github.com/google/uuid"
	"something/domain/entity"
	"time"
)

var (
	ErrInvalidLabelx = errors.New("invalid service label - cannot be empty")
)

type Appointment struct {
	approved   bool
	canceled   bool
	id         uuid.UUID
	price      float64
	serviceIds string
	user       entity.User
	appointmentTimings
}

type appointmentTimings struct {
	Duration           time.Duration
	PreBufferDuration  time.Duration
	PostBufferDuration time.Duration
	StartTime          time.Time
	EndTime            time.Time
}

func NewAppointment(user entity.User, serviceIds string, price float64) (Appointment, error) {
	return Appointment{
		id:         uuid.New(),
		user:       user,
		serviceIds: serviceIds,
		price:      price,
	}, nil
}
