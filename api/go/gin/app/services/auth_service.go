package services

import (
	"errors"
	"itau-api/app/models"
	"itau-api/app/repositories"
	"itau-api/app/util"
)

type AuthService struct {
	UserService     *UserService
	UserRepository  *repositories.UserRepository
	TokenRepository *repositories.TokenRepository
}

func NewAuthService(
	userService *UserService,
	userRepository *repositories.UserRepository,
	tokenRepository *repositories.TokenRepository,
) *AuthService {
	return &AuthService{
		UserService:     userService,
		UserRepository:  userRepository,
		TokenRepository: tokenRepository,
	}
}

func (s *AuthService) AuthenticateUser(document string, password string) (*models.User, string, error) {
	// Find User By Document
	var authUser *models.User
	if user, err := s.UserService.FindUserByDocument(document); err == nil {
		authUser = user
	} else {
		return nil, "", err
	}

	// Check password
	if same := util.ComparePassword(authUser.Password, password); !same {
		return nil, "", errors.New("senhas não coincidem")
	}

	// Simplify exposed token
	_, publicToken := s.TokenRepository.Create(authUser.ID)
	return authUser, publicToken, nil
}
