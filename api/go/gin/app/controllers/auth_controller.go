package controllers

import (
	"github.com/gin-gonic/gin"
	"itau-api/app/exceptions"
	"itau-api/app/models"
	"itau-api/app/repositories"
	"itau-api/app/requests"
	"itau-api/app/services"
)

type AuthController struct {
	userService        *services.UserService
	authService        *services.AuthService
	bankAccountService *services.BankAccountService
	tokenRepository    *repositories.TokenRepository
}

func NewAuthController(
	userService *services.UserService,
	authService *services.AuthService,
	bankAccountService *services.BankAccountService,
	tokenRepository *repositories.TokenRepository,
) *AuthController {
	return &AuthController{
		userService:        userService,
		authService:        authService,
		bankAccountService: bankAccountService,
		tokenRepository:    tokenRepository,
	}
}

func (c *AuthController) Auth(ctx *gin.Context) {
	user := ctx.MustGet("user").(models.User)

	// Find User Bank Accounts
	var bankAccounts []models.BankAccount
	if err := c.bankAccountService.FindBankAccountsByUserId(&bankAccounts, user.ID); err != nil {
		exceptions.NewUnexpectedError(err, ctx)
		return
	}

	// Set Bank Accounts
	user.BankAccounts = bankAccounts

	// Return created user
	ctx.JSON(200, user)
}

func (c *AuthController) Register(ctx *gin.Context) {
	// Start Request Struct
	var form RegisterRequest

	// Validate JSON and fill request data
	if err := ctx.ShouldBindJSON(&form); err != nil {
		exceptions.NewInvalidRequestBody(err, ctx)
		return
	}

	// Validate Request Struct
	if err := requests.Validate(form); err != nil {
		requests.NewUnprocessableEntityException(err, form, ctx)
		return
	}

	// Validate if user document is already in use
	if user, _ := c.userService.FindUserByDocument(form.Document); user != nil {
		ctx.JSON(422, gin.H{
			"document.unique": "Este documento já está em uso.",
		})
		return
	}

	// Create User
	user := models.User{
		Name:     form.Name,
		Document: form.Document,
		Password: form.Password,
		Balance:  0,
	}

	if err := c.userService.CreateUser(&user); err != nil {
		exceptions.NewUnexpectedError(err, ctx)
		return
	}

	response := make(map[string]any)
	response["user"] = user

	if publicToken, err := c.tokenRepository.Create(&models.ApiToken{}, user.ID); err != nil {
		exceptions.NewUnexpectedError(err, ctx)
		return
	} else {
		response["token"] = publicToken
	}

	// Return created user
	ctx.JSON(201, response)
}

func (c *AuthController) Login(ctx *gin.Context) {
	// Start Request Struct
	var form LoginRequest

	// Validate JSON and fill request data
	if err := ctx.ShouldBindJSON(&form); err != nil {
		exceptions.NewInvalidRequestBody(err, ctx)
		return
	}

	// Validate Request Struct
	if err := requests.Validate(form); err != nil {
		requests.NewUnprocessableEntityException(err, form, ctx)
		return
	}

	user, token, err := c.authService.AuthenticateUser(form.Document, form.Password)
	if err != nil {
		exceptions.NewUnexpectedError(err, ctx)
		return
	}

	ctx.JSON(201, gin.H{
		"user":  user,
		"token": token,
	})
}

type RegisterRequest struct {
	requests.BaseRequest
	Name                 string `json:"name" validate:"required,min=1,max=255"`
	Document             string `json:"document" validate:"required,min=1,max=255"`
	Password             string `json:"password" validate:"required,min=8,max=72"`
	PasswordConfirmation string `json:"password_confirmation" validate:"required,eqfield=Password"`
}

type LoginRequest struct {
	requests.BaseRequest
	Document string `json:"document" validate:"required,min=1,max=255"`
	Password string `json:"password" validate:"required,min=1,max=255"`
}
