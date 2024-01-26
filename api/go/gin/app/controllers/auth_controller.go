package controllers

import (
	"github.com/gin-gonic/gin"
	"itau-api/app/exceptions"
	"itau-api/app/models"
	"itau-api/app/requests"
	"itau-api/app/services"
)

type AuthController struct {
	userService *services.UserService
	authService *services.AuthService
}

func NewAuthController(
	userService *services.UserService,
	authService *services.AuthService,
) *AuthController {
	return &AuthController{
		userService: userService,
		authService: authService,
	}
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
	user, err := c.userService.CreateUser(&models.User{
		Name:     form.Name,
		Document: form.Document,
		Password: form.Password,
		Balance:  0,
	})

	// Something Happened
	if err != nil {
		exceptions.NewUnexpectedError(err, ctx)
		return
	}

	// Return created user
	ctx.JSON(201, user)
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
