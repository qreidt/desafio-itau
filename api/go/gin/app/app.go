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
	router       *gin.Engine
	Db           *gorm.DB
	Repositories Repositories
	Services     Services
	Middlewares  Middlewares
	Controllers  Controllers
}

type Repositories struct {
	UserRepository        *repositories.UserRepository
	TokenRepository       *repositories.TokenRepository
	BankAccountRepository *repositories.BankAccountRepository
}

type Services struct {
	AuthService        *services.AuthService
	UserService        *services.UserService
	BankAccountService *services.BankAccountService
}

type Middlewares struct {
	AuthMiddleware *middlewares.AuthMiddleware
}

type Controllers struct {
	AuthController        *controllers.AuthController
	TransactionController *controllers.TransactionController
}

func NewApp(db *gorm.DB) *App {
	app := &App{router: gin.Default(), Db: db}

	app.Repositories = Repositories{
		UserRepository:        repositories.NewUserRepository(db),
		TokenRepository:       repositories.NewTokenRepository(db),
		BankAccountRepository: repositories.NewBankAccountRepository(db),
	}

	app.Services = Services{
		UserService:        services.NewUserService(db, app.Repositories.UserRepository),
		AuthService:        services.NewAuthService(app.Repositories.UserRepository, app.Repositories.TokenRepository),
		BankAccountService: services.NewBankAccountService(app.Repositories.BankAccountRepository),
	}

	app.Middlewares = Middlewares{
		AuthMiddleware: middlewares.NewAuthMiddeware(
			app.Repositories.UserRepository,
			app.Repositories.TokenRepository,
		),
	}

	app.Controllers = Controllers{
		AuthController: controllers.NewAuthController(
			app.Services.UserService,
			app.Services.AuthService,
			app.Services.BankAccountService,
			app.Repositories.TokenRepository,
		),
		TransactionController: controllers.NewTransactionController(),
	}

	return app
}

func (app *App) SetupRoutes() {
	appControllers := app.Controllers

	app.router.POST("/register", appControllers.AuthController.Register)
	app.router.POST("/login", appControllers.AuthController.Login)
	//app.router.POST("/logout", appControllers.AuthController.Logout)

	useAuth := app.router.Use(app.Middlewares.AuthMiddleware.UseAuth)
	{
		useAuth.GET("/auth", appControllers.AuthController.Auth)
		transfers := app.router.Group("/transfers")
		{
			transfers.GET("/", appControllers.TransactionController.Index)
			transfers.POST("/", appControllers.TransactionController.Store)
			transfers.GET("/:transfer", appControllers.TransactionController.Show)
		}
	}
}

func (app *App) RunServer(address string) {
	err := app.router.Run(address)
	if err != nil {
		panic(err)
	}
}
