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

func (c *AuthController) Register(context *gin.Context) {
	// Start Request Struct
	var form RegisterRequest

	// Validate JSON and fill request data
	if err := context.ShouldBindJSON(&form); err != nil {
		exceptions.NewInvalidRequestBody(err, context)
		return
	}

	// Validate Request Struct
	if err := requests.Validate(form); err != nil {
		requests.NewUnprocessableEntityException(err, form, context)
		return
	}

	// Validate if user document is already in use
	if user, _ := c.userRepository.FindByDocument(form.Document); user != nil {
		context.JSON(422, gin.H{
			"document.unique": "Este documento já está em uso.",
		})
		return
	}

	var user models.User
	// Hash user passowrd
	if hashedPassword, err := models.HashPassword(form.Password); err != nil {
		exceptions.NewUnexpectedError(err, context)
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
		exceptions.NewUnexpectedError(err, context)
		return
	}

	// Return created user
	context.JSON(201, user)
}

type RegisterRequest struct {
	requests.BaseRequest
	Name                 string `json:"name" validate:"required,min=1,max=255"`
	Document             string `json:"document" validate:"required,min=1,max=255"`
	Password             string `json:"password" validate:"required,min=8,max=72"`
	PasswordConfirmation string `json:"password_confirmation" validate:"required,eqfield=Password"`
}
