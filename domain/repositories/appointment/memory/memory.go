package memory

import (
	"something/aggregate"
	"something/domain/repositories/appointment"
	"sync"
)

type memoryRepository struct {
	sync.Mutex
	appointments map[string]aggregate.Appointment
}

func NewAppointmentRepository() appointment.Repository {
	return memoryRepository{}
}

func (m memoryRepository) Add(id string, appointment aggregate.Appointment) (err error) {
	//TODO implement me
	panic("implement me")
}

func (m memoryRepository) Update(id string, appointment aggregate.Appointment) (err error) {
	//TODO implement me
	panic("implement me")
}

func (m memoryRepository) Remove(id string) (err error) {
	//TODO implement me
	panic("implement me")
}

func (m memoryRepository) Get(id string) (appointment aggregate.Appointment, err error) {
	//TODO implement me
	panic("implement me")
}

func (m memoryRepository) GetAll() (appointments []aggregate.Appointment, err error) {
	//TODO implement me
	panic("implement me")
}
