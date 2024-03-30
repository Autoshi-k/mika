package services

import (
	"github.com/google/uuid"
	"something/aggregate"
	"something/application/contracts"
	"something/domain/repositories/service"
	"something/infrastructure/serviceReply"
)

type AdministrationConfiguration func(s *AdministrationService) error

type AdministrationService struct {
	services service.Repository
}

func NewAdministrationService(services service.Repository) contracts.AdministrationService {
	return AdministrationService{services: services}
}

func (s AdministrationService) CreateNewService(name, label string, pricing float64) serviceReply.Reply {
	srvc, err := aggregate.NewService(name, label, pricing)
	if err != nil {
		return serviceReply.NewInternalError(err)
	}

	err = s.services.Add(srvc.GetId(), srvc)
	if err != nil {
		return serviceReply.NewDbError(err)
	}

	return nil
}

func (s AdministrationService) EditService(id, name, label string, pricing float64) serviceReply.Reply {
	uid, err := uuid.Parse(id)
	if err != nil {
		return serviceReply.NewInternalError(err)
	}

	srvc, err := s.services.Get(uid)
	if err != nil {
		return serviceReply.NewDbError(err)
	}

	srvc.SetName(name)
	srvc.SetLabel(label)
	srvc.SetPricing(pricing)

	err = s.services.Update(uid, srvc)
	if err != nil {
		return serviceReply.NewDbError(err)
	}

	return nil
}

func (s AdministrationService) RemoveService(id string) serviceReply.Reply {
	uid, err := uuid.Parse(id)
	if err != nil {
		return serviceReply.NewInternalError(err)
	}

	err = s.services.Remove(uid)
	if err != nil {
		return serviceReply.NewDbError(err)
	}

	return nil
}
