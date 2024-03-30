package sql

import (
	"github.com/google/uuid"
	"something/aggregate"
	"something/domain/repositories/service"
	"something/infrastructure/db"
)

type mySqlServiceRepo struct {
	db db.Connection
}

func NewServiceRepository() service.Repository {
	return &mySqlServiceRepo{db: db.NewDbConnection()}
}

func (repo mySqlServiceRepo) Add(id uuid.UUID, service aggregate.Service) (err error) {
	return
}

func (repo mySqlServiceRepo) Update(id uuid.UUID, service aggregate.Service) (err error) {
	return
}

func (repo mySqlServiceRepo) Remove(id uuid.UUID) (err error) {
	return
}

func (repo mySqlServiceRepo) Get(id uuid.UUID) (service aggregate.Service, err error) {
	return
}

func (repo mySqlServiceRepo) GetAll() (services []aggregate.Service, err error) {
	return
}
