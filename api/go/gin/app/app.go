package app

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"itau-api/app/controllers"
	"itau-api/app/repositories"
	"itau-api/app/services"
)

type App struct {
	db           *gorm.DB
	router       *gin.Engine
	repositories Repositories
	services     Services
	controllers  Controllers
}

type Repositories struct {
	UserRepository *repositories.UserRepository
}

type Services struct {
	AuthService *services.AuthService
	UserService *services.UserService
}

type Controllers struct {
	AuthController *controllers.AuthController
}

func NewApp(db *gorm.DB) *App {
	r := Repositories{
		UserRepository: repositories.NewUserRepository(db),
	}

	userService := services.NewUserService(r.UserRepository)
	s := Services{
		UserService: userService,
		AuthService: services.NewAuthService(userService, r.UserRepository),
	}

	c := Controllers{
		AuthController: controllers.NewAuthController(s.UserService, s.AuthService),
	}

	return &App{
		db:           db,
		router:       gin.Default(),
		repositories: r,
		services:     s,
		controllers:  c,
	}
}

func (app *App) SetupRoutes() {
	appControllers := app.controllers

	//app.router.GET('/auth', )
	app.router.POST("/register", appControllers.AuthController.Register)
	app.router.POST("/login", appControllers.AuthController.Login)
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
	err := app.router.Run(address)
	if err != nil {
		panic(err)
	}
}
