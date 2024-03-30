package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"something/application/contracts"
	"something/application/services"
	"something/notsureyet"
)

type App struct {
	administration contracts.AdministrationService // todo change to api?
}

func main() {
	administration, err := services.NewAdministrationService(services.MemoryServicesRepository())
	if err != nil {
		panic(err)
	}

	app := App{administration: administration}
	r := mux.NewRouter()
	notsureyet.NewAdministrationController(r, app.administration)

	fmt.Println("program running")
	http.ListenAndServe(":80", r)
	fmt.Println("program stopped")
}
