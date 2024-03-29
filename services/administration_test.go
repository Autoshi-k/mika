package services

import (
	"github.com/stretchr/testify/assert"
	"something/aggregate"
	"testing"
)

func TestAdministrationService(t *testing.T) {
	as, err := NewAdministrationService(
		MemoryServicesRepository(),
	)

	if err != nil {
		t.Error(err)
	}

	srv1, err := aggregate.NewService("my-first-service", "The Best Service", 5.0)
	if err != nil {
		t.Error(err)
	}

	err = as.services.Add(srv1)
	if err != nil {
		t.Error(err)
	}

	srv2, err := aggregate.NewService("my-second-service", "A great Service", 10.0)
	if err != nil {
		t.Error(err)
	}

	err = as.services.Add(srv2)
	if err != nil {
		t.Error(err)
	}

	srv3, err := as.services.Get(srv1.Id)
	if err != nil {
		t.Error(err)
	}

	assert.Equal(t, srv1, srv3)
	assert.NotEqual(t, srv1, srv2)
}
