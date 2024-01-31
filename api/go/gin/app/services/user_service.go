package services

import (
	"gorm.io/gorm"
	"itau-api/app/models"
	"itau-api/app/repositories"
	"itau-api/app/util"
)

type UserService struct {
	db             *gorm.DB
	userRepository *repositories.UserRepository
}

func NewUserService(db *gorm.DB, userRepository *repositories.UserRepository) *UserService {
	return &UserService{db: db, userRepository: userRepository}
}

// Store

func (s *UserService) CreateUser(user *models.User) error {

	// Hash user passowrd
	if hashedPassword, err := util.HashPassword(user.Password); err != nil {
		return err
	} else {
		user.Password = hashedPassword
	}

	// Start Transaction
	return s.db.Transaction(func(tx *gorm.DB) error {
		userRepository := repositories.NewUserRepository(tx)
		bankRepository := repositories.NewBankAccountRepository(tx)

		// Save user
		if err := userRepository.Create(user); err != nil {
			return err
		}

		bankAccount := models.BankAccount{
			UserId:  user.ID,
			Balance: 0,
		}

		// Create Bank Account
		if err := bankRepository.Create(&bankAccount); err != nil {
			return err
		}

		user.BankAccounts = []models.BankAccount{bankAccount}

		return nil

	})
}

// Find

func (s *UserService) FindUserByDocument(document string) (*models.User, error) {
	var user models.User
	if err := s.userRepository.FindByDocument(document, &user); err != nil {
		return nil, err
	}

	return &user, nil
}
