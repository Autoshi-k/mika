package services

import (
	"something/aggregate"
	"something/domain/repositories/service/memory"
	"testing"
)

func TestAdministrationService(t *testing.T) {
	as := NewAdministrationService(memory.NewServiceRepository())

	sErr := as.CreateNewService("my-first-service", "The Best Service", 5.0)
	if sErr != nil {
		t.Error(sErr)
	}

	srv2, err := aggregate.NewService("my-second-service", "A great Service", 10.0)
	if err != nil {
		t.Error(err)
	}

	sErr = as.EditService(srv2.GetId().String(), "my-first-service", "The Best Service", 10.7)
	if sErr != nil {
		t.Error(sErr)
	}

	//assert.Equal(t, srv1, srv3)
	//assert.NotEqual(t, srv1, srv2)
}
