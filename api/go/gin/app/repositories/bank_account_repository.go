package repositories

import (
	"gorm.io/gorm"
	"itau-api/app/models"
)

type BankAccountRepository struct {
	db *gorm.DB
}

func NewBankAccountRepository(db *gorm.DB) *BankAccountRepository {
	return &BankAccountRepository{db: db}
}

func (r *BankAccountRepository) GetUserBankAccounts(bankAccounts *[]models.BankAccount, userId uint64) error {
	result := r.db.Find(bankAccounts, "user_id = ?", userId)
	return result.Error
}

func (r *BankAccountRepository) Create(bankAccount *models.BankAccount) error {
	return r.db.Create(bankAccount).Error
}
