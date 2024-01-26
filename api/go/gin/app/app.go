package app

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"itau-api/app/controllers"
)

type App struct {
	db     *gorm.DB
	router *gin.Engine
}

func NewApp(db *gorm.DB) *App {
	return &App{
		db:     db,
		router: gin.Default(),
	}
}

func (app *App) SetupRoutes() {
	authController := controllers.NewUserController(app.db)
	//app.router.GET('/auth', )
	app.router.POST("/register", authController.Register)
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
