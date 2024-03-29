package repository

import (
	"something/domain/aggregate"
	"something/domain/service"
	"something/infrastructure/db"
)

type mySqlServiceRepo struct {
	db db.Connection
}

func NewServiceRepository() service.Repository {
	return &mySqlServiceRepo{db: db.NewDbConnection()}
}

func (repo mySqlServiceRepo) Add(service aggregate.Service) (err error) {
	return
}

func (repo mySqlServiceRepo) Update(id string, service aggregate.Service) (err error) {
	return
}

func (repo mySqlServiceRepo) Remove(id string) (err error) {
	return
}

func (repo mySqlServiceRepo) Get(id string) (service aggregate.Service, err error) {
	return
}

func (repo mySqlServiceRepo) GetAll() (services []aggregate.Service, err error) {
	return
}
