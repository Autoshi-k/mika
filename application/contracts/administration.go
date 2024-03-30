package contracts

import "something/infrastructure/serviceReply"

type AdministrationService interface {
	CreateNewService(name, label string, pricing float64) (string, serviceReply.Reply)
	EditService(id, name, label string, pricing float64) serviceReply.Reply
	RemoveService(id string) serviceReply.Reply
}

// todo need to figure out where to set Req structs
