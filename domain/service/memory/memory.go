package memory

import (
	"fmt"
	"github.com/google/uuid"
	"something/aggregate"
	"something/domain/service"
	"sync"
)

type memoryRepository struct {
	sync.Mutex
	services map[uuid.UUID]aggregate.Service
}

func NewServiceRepository() service.Repository {
	return &memoryRepository{services: make(map[uuid.UUID]aggregate.Service)}
}

func (repo memoryRepository) Add(s aggregate.Service) (err error) {
	repo.Lock()
	defer repo.Unlock()

	if _, ok := repo.services[s.Id]; ok {
		return fmt.Errorf("service already exist %w", service.ErrFailedToAddCustomer)
	} else {
		repo.services[s.Id] = s
		return nil
	}
}

func (repo memoryRepository) Update(id uuid.UUID, s aggregate.Service) (err error) {
	repo.Lock()
	defer repo.Unlock()

	if _, ok := repo.services[id]; !ok {
		return service.ErrServiceNotFound
	} else {
		repo.services[id] = s
		return nil
	}
}

func (repo memoryRepository) Remove(id uuid.UUID) (err error) {
	repo.Lock()
	defer repo.Unlock()

	if _, ok := repo.services[id]; !ok {
		return service.ErrServiceNotFound
	} else {
		delete(repo.services, id)
		return nil
	}
}

func (repo memoryRepository) Get(id uuid.UUID) (aggregate.Service, error) {
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
