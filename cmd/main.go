package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"something/application/mikasapp"
	"something/application/services"
	appointmentmemory "something/domain/repositories/appointment/memory"
	servicememory "something/domain/repositories/service/memory"
	"something/notsureyet"
)

func main() {
	administration := services.NewAdministrationService(servicememory.NewServiceRepository())
	scheduler := services.NewSchedulerService(appointmentmemory.NewAppointmentRepository())

	app := mikasapp.NewApp(administration, scheduler)
	r := mux.NewRouter()
	notsureyet.NewAdministrationController(r, app) // todo router should be on app not single services

	fmt.Println("program running")
	err := http.ListenAndServe(":8002", r)
	if err != nil {
		panic(err)
	}
	fmt.Println("program stopped")
}
