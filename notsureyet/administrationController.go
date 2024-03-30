package notsureyet

import (
	"encoding/json"
	"fmt"
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
		Name    string
		Label   string
		Pricing float64
	}{}

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	res, sErr := c.service.CreateNewService(req.Name, req.Label, req.Pricing)
	if sErr != nil {
		// todo can't get sErr.Error(), why
		http.Error(w, "sErr", http.StatusInternalServerError)
	}

	//w.WriteHeader(http.StatusOK)
	_, err = w.Write([]byte(res))
	if err != nil {
		fmt.Println(err)
	}
}

func (c controller) EditService(w http.ResponseWriter, r *http.Request) {
	req := struct {
		Id      string
		Name    string
		Label   string
		Pricing float64
	}{}

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	sErr := c.service.EditService(req.Id, req.Name, req.Label, req.Pricing)
	if sErr != nil {
		// todo can't get sErr.Error(), why
		http.Error(w, "sErr", http.StatusInternalServerError)
	}
}

func (c controller) RemoveService(w http.ResponseWriter, r *http.Request) {
	req := struct {
		Id string
	}{}

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	sErr := c.service.RemoveService(req.Id)
	if sErr != nil {
		// todo can't get sErr.Error(), why
		http.Error(w, "sErr", http.StatusInternalServerError)
	}
}
