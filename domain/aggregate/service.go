package aggregate

import (
	"errors"
	"github.com/google/uuid"
)

var (
	ErrInvalidLabel = errors.New("invalid repository label - cannot be empty")
	ErrInvalidName  = errors.New("invalid repository name - cannot be empty")
)

type Service struct {
	IsActive bool
	Pricing  float64
	Id       string
	Name     string
	Label    string
}

func NewService(name, label string, pricing float64) (Service, error) {
	if len(label) == 0 {
		return Service{}, ErrInvalidLabel
	}

	if len(label) == 0 {
		return Service{}, ErrInvalidName
	}

	return Service{
		IsActive: true,
		Pricing:  pricing,
		Id:       uuid.NewString(),
		Label:    label,
	}, nil
}
