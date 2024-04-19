package api

import (
	"something/domain/entity"
)

type AddNewServiceRequest struct {
	Durations       entity.DurationDetails
	Name            string
	Label           string
	Pricing         float64
	PreOrder        float64
	AvailableInDays int
}

type AddNewServiceResponse string

//type AddNewServiceResponse struct {
//	Id string `json:"id"`
//}

func (req AddNewServiceRequest) Validate() bool {
	return req.Name != "" && req.Label != ""
}

type UpdateServiceRequest struct {
	Id string `json:"id"`
	AddNewServiceRequest
}

type UpdateServiceResponse struct {
}

func (req UpdateServiceRequest) Validate() bool {
	return req.Id != "" && req.Name != "" && req.Label != ""
}

type RemoveServiceRequest struct {
	Id string `json:"id"`
}

type RemoveServiceResponse struct {
}

func (req RemoveServiceRequest) Validate() bool {
	return req.Id != ""
}
