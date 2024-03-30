package notsureyet

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"something/application/contracts"
)

type controller struct {
	service contracts.AdministrationService
}

func NewAdministrationController(r *mux.Router, service contracts.AdministrationService) {
	c := controller{service: service}
	r.HandleFunc("/hello", c.CreateNewService)
	r.HandleFunc("/editService", c.EditService)
	r.HandleFunc("/removeService", c.RemoveService)
}

func (c controller) CreateNewService(w http.ResponseWriter, r *http.Request) {
	req := struct {
	}{}

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	c.service.CreateNewService()
}

func (c controller) EditService(w http.ResponseWriter, r *http.Request) {
	c.service.EditService()
}

func (c controller) RemoveService(w http.ResponseWriter, r *http.Request) {
	c.service.RemoveService()
}
