package appointment

import "something/aggregate"

type Repository interface {
	Add(id string, appointment aggregate.Appointment) (err error)
	Update(id string, appointment aggregate.Appointment) (err error)
	Remove(id string) (err error)
	Get(id string) (appointment aggregate.Appointment, err error)
	GetAll() (appointment []aggregate.Appointment, err error)
}
