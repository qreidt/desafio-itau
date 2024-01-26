package services

import (
	"itau-api/app/models"
	"itau-api/app/repositories"
	"itau-api/app/util"
)

type UserService struct {
	repository *repositories.UserRepository
}

func NewUserService(repository *repositories.UserRepository) *UserService {
	return &UserService{repository: repository}
}

// Store

func (s *UserService) CreateUser(user *models.User) (*models.User, error) {

	// Hash user passowrd
	if hashedPassword, err := util.HashPassword(user.Password); err != nil {
		return nil, err
	} else {
		user.Password = hashedPassword
	}

	// Save user
	if err := s.repository.Create(user); err != nil {
		return nil, err
	}

	return user, nil
}

// Find

func (s *UserService) FindUserByDocument(document string) (*models.User, error) {
	var user models.User
	if err := s.repository.FindByDocument(document, &user); err != nil {
		return nil, err
	}

	return &user, nil
}
