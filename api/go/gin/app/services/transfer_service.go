package services

import (
	"itau-api/app/models"
	"itau-api/app/repositories"
)

type TransferService struct {
	transferRepository *repositories.TransferRepository
}

func NewTransferService(transferRepository *repositories.TransferRepository) *TransferService {
	return &TransferService{transferRepository: transferRepository}
}

func (s *TransferService) ListUserTransfers(transfers *[]models.Transfer, userId uint64) error {
	return s.transferRepository.GetBySenderOrReceiverUserId(transfers, userId)
}

func (s *TransferService) CreateTransfer(transfer *models.Transfer) error {
	return s.transferRepository.Create(transfer)
}

func (s *TransferService) FindTransfer(transfer *models.Transfer, transferId uint64) error {
	return s.transferRepository.FindTransfer(transfer, transferId)
}
