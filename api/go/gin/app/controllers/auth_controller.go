package controllers

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"itau-api/app/exceptions"
	"itau-api/app/models"
	"itau-api/app/repositories"
	"itau-api/app/requests"
)

type AuthController struct {
	userRepository *repositories.UserRepository
}

func NewUserController(db *gorm.DB) *AuthController {
	return &AuthController{userRepository: repositories.NewUserRepository(db)}
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
	var user models.User
	if _ = c.userRepository.FindByDocument(form.Document, &user); user != (models.User{}) {
		ctx.JSON(422, gin.H{
			"document.unique": "Este documento já está em uso.",
		})
		return
	}

	// Hash user passowrd
	if hashedPassword, err := models.HashPassword(form.Password); err != nil {
		exceptions.NewUnexpectedError(err, ctx)
		return
	} else {

		// Fill user model with request and default data
		user = models.User{
			Name:     form.Name,
			Document: form.Document,
			Password: hashedPassword,
			Balance:  0,
		}
	}

	// Save user
	if err := c.userRepository.Create(&user); err != nil {
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

	// Find User By Document
	var user models.User
	if err := c.userRepository.FindByDocument(form.Document, &user); err != nil {
		ctx.JSON(401, gin.H{
			"message": "Credenciais não encontradas",
		})
		return
	}

	// Check password
	if same := models.ComparePassword(user.Password, form.Password); !same {
		ctx.JSON(401, gin.H{
			"message": "Credenciais não encontradas",
		})
		return
	}

	ctx.JSON(201, gin.H{
		"user":  user,
		"token": "",
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
