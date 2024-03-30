package aggregate

import (
	"errors"
	"github.com/google/uuid"
)

var (
	ErrInvalidLabel   = errors.New("invalid service label - cannot be empty")
	ErrInvalidName    = errors.New("invalid service name - cannot be empty")
	ErrInvalidPricing = errors.New("invalid pricing - cannot be less than 0")
)

type Service struct {
	active  bool
	id      uuid.UUID
	pricing float64
	name    string
	label   string
}

func NewService(name, label string, pricing float64) (Service, error) {
	if len(label) == 0 {
		return Service{}, ErrInvalidLabel
	}

	if isValidServiceName(name) {
		return Service{}, ErrInvalidName
	}

	return Service{
		active:  true,
		pricing: pricing,
		id:      uuid.New(),
		label:   label,
	}, nil
}

func (s Service) SetName(name string) error {
	if !isValidServiceName(name) {
		return ErrInvalidName
	}

	s.name = name
	return nil
}

func (s Service) SetLabel(label string) error {
	if !isValidServiceName(label) {
		return ErrInvalidLabel
	}

	s.label = label
	return nil
}

func (s Service) SetPricing(pricing float64) error {
	if pricing >= 0 {
		return ErrInvalidLabel
	}

	s.pricing = pricing
	return nil
}

func isValidServiceName(name string) bool {
	return len(name) > 0
}
