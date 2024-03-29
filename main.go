package main

import (
	"fmt"
	"something/domain/service"
	"something/infrastructure/repository"
)

type App struct {
	// todo logger
	services service.Repository
}

func main() {
	app := App{
		services: repository.NewServiceRepository(),
	}
	fmt.Println("program running", app)
	fmt.Println("program stopped")
}
