package services

import (
	"itau-api/app/models"
	"itau-api/app/repositories"
	"itau-api/app/util"
)

type BankAccountService struct {
	BankAccountRepository *repositories.BankAccountRepository
}

func NewBankAccountService(repository *repositories.BankAccountRepository) *BankAccountService {
	return &BankAccountService{BankAccountRepository: repository}
}

func (s *BankAccountService) FindBankAccountsByUserId(bankAccounts *[]models.BankAccount, userId uint64) error {
	return s.BankAccountRepository.GetUserBankAccounts(bankAccounts, userId)
}

func (s *BankAccountService) CreateBankAccount(bankAccount *models.BankAccount) error {
	bankAccount.Number = util.RandNumbers(10)
	return s.BankAccountRepository.Create(bankAccount)
}
