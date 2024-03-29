package services

import (
	"something/domain/service"
	"something/domain/service/memory"
	"something/domain/service/sql"
)

type AdministrationConfiguration func(s *AdministrationService) error

type AdministrationService struct {
	services service.Repository
}

func NewAdministrationService(cfgs ...AdministrationConfiguration) (*AdministrationService, error) {
	s := &AdministrationService{}

	// Apply all Configurations passed in
	for _, cfg := range cfgs {
		err := cfg(s)
		if err != nil {
			return nil, err
		}
	}
	return s, nil
}

func MemoryServicesRepository() AdministrationConfiguration {
	return func(s *AdministrationService) error {
		s.services = memory.NewServiceRepository()
		return nil
	}
}

func SqlServicesRepository() AdministrationConfiguration {
	return func(s *AdministrationService) error {
		s.services = sql.NewServiceRepository()
		return nil
	}
}
