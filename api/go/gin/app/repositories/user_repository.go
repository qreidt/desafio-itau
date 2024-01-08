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

func (repo *UserRepository) Create(user *models.User) error {
	return repo.db.Create(user).Error
}

func (repo *UserRepository) FindById(id uint64) (*models.User, error) {
	var user models.User
	result := repo.db.First(&user, id)

	return &user, result.Error
}

func (repo *UserRepository) FindByDocument(document string) (*models.User, error) {
	var user models.User
	result := repo.db.First(&user, "document = ?", document)

	return &user, result.Error
}

func (repo *UserRepository) Update(user *models.User) error {
	return repo.db.Save(user).Error
}

func (repo *UserRepository) Delete(user *models.User) error {
	return repo.db.Save(user).Error
}
