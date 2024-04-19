package api

import (
	"context"
	"fmt"
	"something/application/contracts"
	"something/infrastructure/http"
	"something/infrastructure/serviceReply"
)

type controller struct {
	app contracts.App
}

func InitHandlers(r *http.Router, app contracts.App) {
	c := controller{app: app}
	r.POST("/addNewService", c.AddNewService)
	r.POST("/editService", c.UpdateService)
	r.POST("/removeService", c.RemoveService)
}

func (c controller) AddNewService(ctx context.Context, req AddNewServiceRequest) (AddNewServiceResponse, serviceReply.Reply) {
	if ok := req.Validate(); !ok {
		return "", fmt.Errorf("bad") // service reply err
	}
	res, err := c.app.AddNewService(ctx, req.Name, req.Label, req.Durations, req.Pricing, req.PreOrder, req.AvailableInDays)
	return AddNewServiceResponse(res), err // todo should return without converting
}

func (c controller) UpdateService(ctx context.Context, req UpdateServiceRequest) (UpdateServiceResponse, serviceReply.Reply) {
	if ok := req.Validate(); !ok {
		return UpdateServiceResponse{}, fmt.Errorf("bad") // service reply err
	}
	err := c.app.UpdateService(ctx, req.Id, req.Name, req.Label, req.Durations, req.Pricing, req.PreOrder, req.AvailableInDays)
	return UpdateServiceResponse{}, err
}

func (c controller) RemoveService(ctx context.Context, req RemoveServiceRequest) (RemoveServiceResponse, serviceReply.Reply) {
	if ok := req.Validate(); !ok {
		return RemoveServiceResponse{}, fmt.Errorf("bad") // service reply err
	}
	err := c.app.RemoveService(ctx, req.Id)
	return RemoveServiceResponse{}, err
}
