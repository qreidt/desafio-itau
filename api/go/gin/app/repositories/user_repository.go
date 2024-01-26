package repositories

import (
	"fmt"
	"gorm.io/gorm"
	"itau-api/app/models"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db}
}

// Create a User
func (repo *UserRepository) Create(user *models.User) error {
	return repo.db.Create(user).Error
}

// FindById Find an user from by ID
func (repo *UserRepository) FindById(id uint64) (*models.User, error) {
	var user models.User
	result := repo.db.First(&user, id)

	return &user, result.Error
}

// FindByDocument Find a user from a document
func (repo *UserRepository) FindByDocument(document string) (*models.User, error) {
	var user models.User
	result := repo.db.First(&models.User{}, "document = ?", document)

	fmt.Println(result)

	return &user, result.Error
}

// Update a user
func (repo *UserRepository) Update(user *models.User) error {
	return repo.db.Save(user).Error
}

// Delete a user
func (repo *UserRepository) Delete(user *models.User) error {
	return repo.db.Save(user).Error
}
