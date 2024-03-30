package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"something/application/contracts"
	"something/application/services"
	"something/domain/repositories/service/memory"
	"something/notsureyet"
)

type App struct {
	administration contracts.AdministrationService // todo change to api?
}

func main() {
	administration := services.NewAdministrationService(memory.NewServiceRepository())

	app := App{administration: administration}
	r := mux.NewRouter()
	notsureyet.NewAdministrationController(r, app.administration)

	fmt.Println("program running")
	err := http.ListenAndServe(":8002", r)
	if err != nil {
		panic(err)
	}
	fmt.Println("program stopped")
}
