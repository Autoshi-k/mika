package main

import (
	"fmt"
	"something/api"
	"something/application/mikasapp"
	"something/application/services"
	appointmentmemory "something/domain/repositories/appointment/memory"
	servicememory "something/domain/repositories/service/memory"
	"something/infrastructure/http"
)

func main() {
	fmt.Println("program running")
	administration := services.NewAdministrationService(servicememory.NewServiceRepository())
	scheduler := services.NewSchedulerService(appointmentmemory.NewAppointmentRepository())
	app := mikasapp.NewApp(administration, scheduler)
	r := http.NewRouter()
	api.InitHandlers(r, app)
	r.Run("8002")
	fmt.Println("program stopped")
}
