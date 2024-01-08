package main

import (
	"fmt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"itau-api/app"
	"itau-api/app/models"
)

func main() {
	fmt.Println("Starting Web Server on Port 4000")

	// Migrate Database
	MigrateDatabase()

	// Run Server
	app := app.NewApp()
	app.SetupRoutes()
	app.RunServer(":4000")
}

func MigrateDatabase() {
	db, err := gorm.Open(sqlite.Open("./tmp/database.sqlite"), &gorm.Config{})

	if err != nil {
		panic("Failed to connect to database.")
	}

	var users models.User
	db.AutoMigrate(&users)
}
