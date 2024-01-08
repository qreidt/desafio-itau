package app

import (
	"github.com/gin-gonic/gin"
)

type App struct {
	router *gin.Engine
}

func NewApp() *App {
	return &App{
		router: gin.Default(),
	}
}

func (app *App) SetupRoutes() {
	//
}

func (app *App) RunServer(address string) {
	app.router.Run(address)
}
