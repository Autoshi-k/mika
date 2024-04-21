package services

import (
	"something/aggregate"
	"something/application/contracts"
	"something/domain/entity"
	"something/domain/repositories/service"
	"something/infrastructure/serviceReply"
	"time"
)

type ProductsConfiguration func(s *ProductService) error

type ProductService struct {
	services service.Repository
}

func NewProductsService(services service.Repository) contracts.ProductsService {
	return ProductService{services: services}
}

func (s ProductService) CreateNewService(name, label string, pricing float64, durations entity.DurationDetails) (string, serviceReply.Reply) {
	srvc, err := aggregate.NewService(name, label, pricing, durations)
	if err != nil {
		return "", serviceReply.NewInternalError(err)
	}

	err = s.services.Add(srvc.GetId(), srvc)
	if err != nil {
		return "", serviceReply.NewDbError(err)
	}

	return srvc.GetId(), nil
}

func (s ProductService) EditService(id, name, label string, pricing float64) serviceReply.Reply {
	srvc, err := s.services.Get(id)
	if err != nil {
		return serviceReply.NewDbError(err)
	}

	srvc.SetName(name)
	srvc.SetLabel(label)
	srvc.SetPricing(pricing)

	err = s.services.Update(id, srvc)
	if err != nil {
		return serviceReply.NewDbError(err)
	}

	return nil
}

func (s ProductService) RemoveService(id string) serviceReply.Reply {
	err := s.services.Remove(id)
	if err != nil {
		return serviceReply.NewDbError(err)
	}

	return nil
}

func (s ProductService) GetServicesDurations(serviceIds []string) (time.Duration, serviceReply.Reply) {
	services, err := s.services.GetMany(serviceIds)
	if err != nil {
		return 0, serviceReply.NewDbError(err)
	}
	return services.GetDurations().CombinedDuration(), nil
}

func (s ProductService) GetAllAvailableServices() (aggregate.Services, serviceReply.Reply) {
	services, err := s.services.GetAll()
	if err != nil {
		return services, serviceReply.NewDbError(err)
	}
	return services, nil
}

func (s ProductService) GetCombinedServicesDetails(serviceIds []string) (entity.DurationDetails, float64, serviceReply.Reply) {
	services, err := s.services.GetMany(serviceIds)
	if err != nil {
		return entity.DurationDetails{}, 0, serviceReply.NewDbError(err)
	}

	return services.GetDurations(), services.GetPricing(), nil
}
