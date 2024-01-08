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
	//app.router.GET('/auth', )
	//app.router.POST('/register', )
	//app.router.POST('/login', )
	//app.router.POST('/logout', )

	//transfers := app.router.Group("/transfers")
	//{
	//	//transfers.GET('/', )
	//	//transfers.POST('/', )
	//	//transfers.GET('/:transfer', )
	//	//transfers.PATCH('/:transfer', )
	//	//transfers.DELETE('/:transfer', )
	//}
}

func (app *App) RunServer(address string) {
	app.router.Run(address)
}
