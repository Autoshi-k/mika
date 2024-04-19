package services

import (
	"something/aggregate"
	"something/domain/entity"
	"something/domain/repositories/service/memory"
	"testing"
)

func TestAdministrationService(t *testing.T) {
	as := NewAdministrationService(memory.NewServiceRepository())

	durations := entity.DurationDetails{
		Duration:           0,
		PreBufferDuration:  0,
		PostBufferDuration: 0,
	}

	_, sErr := as.CreateNewService("my-first-service", "The Best Service", 5.0, durations)
	if sErr != nil {
		t.Error(sErr)
	}

	srv2, err := aggregate.NewService("my-second-service", "A great Service", 10.0, durations)
	if err != nil {
		t.Error(err)
	}

	sErr = as.EditService(srv2.GetId(), "my-first-service", "The Best Service", 10.7)
	if sErr != nil {
		t.Error(sErr)
	}

	//assert.Equal(t, srv1, srv3)
	//assert.NotEqual(t, srv1, srv2)
}
