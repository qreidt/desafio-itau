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
	db := MigrateDatabase()

	// Run Server
	server := app.NewApp(db)
	server.SetupRoutes()
	server.RunServer(":4000")
}

func MigrateDatabase() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("./tmp/database.sqlite"), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database.")
	}

	db.AutoMigrate(
		&models.User{},
		&models.ApiToken{},
		&models.BankAccount{},
		&models.Transfer{},
	)
	return db
}
