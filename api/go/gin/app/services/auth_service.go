package services

import (
	"errors"
	"itau-api/app/models"
	"itau-api/app/repositories"
	"itau-api/app/util"
)

type AuthService struct {
	UserRepository  *repositories.UserRepository
	TokenRepository *repositories.TokenRepository
}

func NewAuthService(
	userRepository *repositories.UserRepository,
	tokenRepository *repositories.TokenRepository,
) *AuthService {
	return &AuthService{
		UserRepository:  userRepository,
		TokenRepository: tokenRepository,
	}
}

func (s *AuthService) AuthenticateUser(document string, password string) (*models.User, string, error) {
	// Find User By Document
	var authUser models.User
	if err := s.UserRepository.FindByDocument(document, &authUser); err != nil {
		return nil, "", err
	}

	// Check password
	if same := util.ComparePassword(authUser.Password, password); !same {
		return nil, "", errors.New("senhas não coincidem")
	}

	// Simplify exposed token
	publicToken, _ := s.TokenRepository.Create(&models.ApiToken{}, authUser.ID)
	return &authUser, publicToken, nil
}
