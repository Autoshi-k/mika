package contracts

import "something/infrastructure/serviceReply"

type AdministrationService interface {
	CreateNewService(name, label string, pricing float64) serviceReply.Reply
	EditService(id, name, label string, pricing float64) serviceReply.Reply
	RemoveService(id string) serviceReply.Reply
}
