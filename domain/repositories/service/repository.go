package service

import (
	"errors"
	"github.com/google/uuid"
	"something/aggregate"
)

var (
	ErrServiceNotFound     = errors.New("the service was not found in the repository")
	ErrFailedToAddCustomer = errors.New("failed to add the service to the repository")
	ErrUpdateCustomer      = errors.New("failed to update the service in the repository")
)

type Repository interface {
	Add(id uuid.UUID, service aggregate.Service) (err error)
	Update(id uuid.UUID, service aggregate.Service) (err error)
	Remove(id uuid.UUID) (err error)
	Get(id uuid.UUID) (service aggregate.Service, err error)
	GetAll() (services []aggregate.Service, err error)
}
