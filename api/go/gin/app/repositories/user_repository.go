package repositories

import (
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
func (repo *UserRepository) FindById(user *models.User, id uint64) error {
	result := repo.db.First(&user, id)
	return result.Error
}

// FindByDocument Find an user from a document
func (repo *UserRepository) FindByDocument(document string, model *models.User) error {
	result := repo.db.First(model, "document = ?", document)
	return result.Error
}

// Update a user
func (repo *UserRepository) Update(user *models.User) error {
	return repo.db.Save(user).Error
}

// Delete a user
func (repo *UserRepository) Delete(user *models.User) error {
	return repo.db.Save(user).Error
}
