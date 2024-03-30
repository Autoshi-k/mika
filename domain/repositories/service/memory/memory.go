package memory

import (
	"fmt"
	"something/aggregate"
	"something/domain/repositories/service"
	"sync"
)

type memoryRepository struct {
	sync.Mutex
	services map[string]aggregate.Service
}

func NewServiceRepository() service.Repository {
	return &memoryRepository{services: make(map[string]aggregate.Service)}
}

func (repo memoryRepository) Add(id string, s aggregate.Service) (err error) {
	repo.Lock()
	defer repo.Unlock()

	if _, ok := repo.services[id]; ok {
		return fmt.Errorf("service already exist %w", service.ErrFailedToAddCustomer)
	} else {
		repo.services[id] = s
		return nil
	}
}

func (repo memoryRepository) Update(id string, s aggregate.Service) (err error) {
	repo.Lock()
	defer repo.Unlock()

	if _, ok := repo.services[id]; !ok {
		return service.ErrServiceNotFound
	} else {
		repo.services[id] = s
		return nil
	}
}

func (repo memoryRepository) Remove(id string) (err error) {
	repo.Lock()
	defer repo.Unlock()

	if _, ok := repo.services[id]; !ok {
		return service.ErrServiceNotFound
	} else {
		delete(repo.services, id)
		return nil
	}
}

func (repo memoryRepository) Get(id string) (aggregate.Service, error) {
	if s, ok := repo.services[id]; !ok {
		return aggregate.Service{}, service.ErrServiceNotFound
	} else {
		return s, nil
	}
}

func (repo memoryRepository) GetAll() (services []aggregate.Service, err error) {
	ss := make([]aggregate.Service, len(repo.services))
	for _, v := range services {
		ss = append(ss, v)
	}
	return ss, nil
}
