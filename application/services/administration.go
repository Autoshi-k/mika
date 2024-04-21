package services

import (
	"something/application/contracts"
)

type AdministrationConfiguration func(s *AdministrationService) error

type AdministrationService struct {
}

func NewAdministrationService() contracts.AdministrationService {
	return AdministrationService{}
}
