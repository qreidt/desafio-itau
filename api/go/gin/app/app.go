package app

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"itau-api/app/controllers"
	"itau-api/app/middlewares"
	"itau-api/app/repositories"
	"itau-api/app/services"
)

type App struct {
	db           *gorm.DB
	router       *gin.Engine
	repositories Repositories
	services     Services
	middlewares  Middlewares
	controllers  Controllers
}

type Repositories struct {
	UserRepository  *repositories.UserRepository
	TokenRepository *repositories.TokenRepository
}

type Services struct {
	AuthService *services.AuthService
	UserService *services.UserService
}

type Middlewares struct {
	AuthMiddleware *middlewares.AuthMiddleware
}

type Controllers struct {
	AuthController        *controllers.AuthController
	TransactionController *controllers.TransactionController
}

func NewApp(db *gorm.DB) *App {
	r := Repositories{
		UserRepository:  repositories.NewUserRepository(db),
		TokenRepository: repositories.NewTokenRepository(db),
	}

	userService := services.NewUserService(r.UserRepository)
	s := Services{
		UserService: userService,
		AuthService: services.NewAuthService(userService, r.UserRepository, r.TokenRepository),
	}

	m := Middlewares{AuthMiddleware: middlewares.NewAuthMiddeware(
		r.UserRepository,
		r.TokenRepository,
	)}

	c := Controllers{
		AuthController: controllers.NewAuthController(s.UserService, s.AuthService),
	}

	return &App{
		db:           db,
		router:       gin.Default(),
		repositories: r,
		services:     s,
		middlewares:  m,
		controllers:  c,
	}
}

func (app *App) SetupRoutes() {
	appControllers := app.controllers

	//app.router.GET('/auth', )
	app.router.POST("/register", appControllers.AuthController.Register)
	app.router.POST("/login", appControllers.AuthController.Login)
	//app.router.POST('/logout', )

	transfers := app.router.Group("/transfers").Use(app.middlewares.AuthMiddleware.UseAuth)
	{
		transfers.GET("/", appControllers.TransactionController.Index)
		transfers.POST("/", appControllers.TransactionController.Store)
		transfers.GET("/:transfer", appControllers.TransactionController.Show)
		transfers.PATCH("/:transfer", appControllers.TransactionController.Update)
		transfers.DELETE("/:transfer", appControllers.TransactionController.Delete)
	}
}

func (app *App) RunServer(address string) {
	err := app.router.Run(address)
	if err != nil {
		panic(err)
	}
}
