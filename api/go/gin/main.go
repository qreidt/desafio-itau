package main

import (
	"fmt"
	"itau-api/app"
)

func main() {
	fmt.Println("Starting Web Server on Port 4000")

	app := app.NewApp()
	app.SetupRoutes()
	app.RunServer(":4000")
}
