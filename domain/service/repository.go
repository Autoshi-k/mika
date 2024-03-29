package service

import "something/domain/aggregate"

type Repository interface {
	Add(service aggregate.Service) (err error)
	Update(id string, service aggregate.Service) (err error)
	Remove(id string) (err error)
	Get(id string) (service aggregate.Service, err error)
	GetAll() (services []aggregate.Service, err error)
}
