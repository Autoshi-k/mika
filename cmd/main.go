package main

import (
	"fmt"
	"something/api"
	"something/application/mikasapp"
	"something/application/services"
	appointmentmemory "something/domain/repositories/appointment/memory"
	blockermemory "something/domain/repositories/blocker/memory"
	productmemory "something/domain/repositories/service/memory"
	timetablememory "something/domain/repositories/timetable/memory"
	"something/infrastructure/http"
)

func main() {
	fmt.Println("program running")
	administration := services.NewAdministrationService()
	scheduler := services.NewSchedulerService(
		appointmentmemory.NewAppointmentRepository(),
		blockermemory.NewBlockerRepository(),
		timetablememory.NewTimetableRepository(),
	)
	srvs := services.NewProductsService(productmemory.NewServiceRepository())

	app := mikasapp.NewApp(administration, scheduler, srvs)

	r := http.NewRouter()
	api.InitHandlers(r, app)
	r.Run("8002")

	fmt.Println("program stopped")
}
