package aggregate

import (
	"errors"
	"github.com/google/uuid"
	"something/domain/entity"
)

var (
	ErrInvalidLabel   = errors.New("invalid service label - cannot be empty")
	ErrInvalidName    = errors.New("invalid service name - cannot be empty")
	ErrInvalidPricing = errors.New("invalid pricing - cannot be less than 0")
)

type Services []Service

type Service struct {
	active  bool
	id      string
	pricing float64
	name    string
	label   string
	entity.DurationDetails
}

func NewService(name, label string, pricing float64) (Service, error) {
	if !isValidServiceName(label) {
		return Service{}, ErrInvalidLabel
	}

	if !isValidServiceName(name) {
		return Service{}, ErrInvalidName
	}

	return Service{
		active:  true,
		pricing: pricing,
		id:      uuid.NewString(),
		name:    name,
		label:   label,
	}, nil
}

func (s *Service) GetId() string {
	return s.id
}

func (s *Service) IsActive() bool {
	return s.active
}

func (s *Service) SetName(name string) error {
	if !isValidServiceName(name) {
		return ErrInvalidName
	}

	s.name = name
	return nil
}

func (s *Service) SetLabel(label string) error {
	if !isValidServiceName(label) {
		return ErrInvalidLabel
	}

	s.label = label
	return nil
}

func (s *Service) SetPricing(pricing float64) error {
	if pricing >= 0 {
		return ErrInvalidLabel
	}

	s.pricing = pricing
	return nil
}

func (s *Service) SetAsInactive() {
	s.active = false
}

func isValidServiceName(name string) bool {
	return len(name) > 0
}

func (ss Services) GetDurations() entity.DurationDetails {
	durations := entity.DurationDetails{}

	for _, s := range ss {
		durations.Duration += s.Duration
		if s.PreBufferDuration > durations.PreBufferDuration {
			durations.PreBufferDuration = s.PreBufferDuration
		}

		if s.PostBufferDuration > durations.PostBufferDuration {
			durations.PostBufferDuration = s.PostBufferDuration
		}
	}

	return durations
}

func (ss Services) GetPricing() float64 {
	var pricing float64

	for _, s := range ss {
		pricing += s.pricing
	}

	return pricing
}
