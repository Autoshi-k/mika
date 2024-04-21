package blocker

import "something/aggregate"

type Repository interface {
	Add(id string, appointment aggregate.TimeBlock) (err error)
	Update(id string, appointment aggregate.TimeBlock) (err error)
	Remove(id string) (err error)
	Get(id string) (appointment aggregate.TimeBlock, err error)
	GetAll() (appointment []aggregate.TimeBlock, err error)
}
