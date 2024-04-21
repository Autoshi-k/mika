package memory

import (
	"something/aggregate"
	"something/domain/repositories/blocker"
	"sync"
)

type memoryRepository struct {
	sync.Mutex
	appointments map[string]aggregate.Appointment
}

func NewBlockerRepository() blocker.Repository {
	return memoryRepository{}
}

func (m memoryRepository) Add(id string, appointment aggregate.TimeBlock) (err error) {
	//TODO implement me
	panic("implement me")
}

func (m memoryRepository) Update(id string, appointment aggregate.TimeBlock) (err error) {
	//TODO implement me
	panic("implement me")
}

func (m memoryRepository) Remove(id string) (err error) {
	//TODO implement me
	panic("implement me")
}

func (m memoryRepository) Get(id string) (appointment aggregate.TimeBlock, err error) {
	//TODO implement me
	panic("implement me")
}

func (m memoryRepository) GetAll() (appointments []aggregate.TimeBlock, err error) {
	//TODO implement me
	panic("implement me")
}
