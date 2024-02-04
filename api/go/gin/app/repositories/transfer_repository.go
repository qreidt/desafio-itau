package repositories

import (
	"gorm.io/gorm"
	"itau-api/app/models"
)

type TransferRepository struct {
	db *gorm.DB
}

func NewTransferRepository(db *gorm.DB) *TransferRepository {
	return &TransferRepository{db: db}
}

func selectUserData(db *gorm.DB) *gorm.DB {
	return db.Select("id, document, name")
}

func selectAccountData(db *gorm.DB) *gorm.DB {
	return db.Select("id, document, name")
}

func (r *TransferRepository) GetBySenderOrReceiverUserId(transfers *[]models.Transfer, userId uint64) error {
	isSenderUserQuery := r.db.Where("sender_user_id = ?", userId)
	isReceiverUserQuery := r.db.Where("receiver_user_id = ?", userId)

	return r.db.Where(isSenderUserQuery.Or(isReceiverUserQuery)).Preload("SenderUser", selectUserData).Preload("ReceiverUser", selectUserData).Preload("SenderAccount", selectAccountData).Preload("ReceiverAccount", selectAccountData).Find(transfers).Error
}

func (r *TransferRepository) FindTransfer(transfer *models.Transfer, transferId uint64) error {
	return r.db.Model(&models.Transfer{}).Preload("SenderUser", selectUserData).Preload("ReceiverUser", selectUserData).Preload("SenderAccount", selectAccountData).Preload("ReceiverAccount", selectAccountData).First(transfer, transferId).Error
}

func (r *TransferRepository) Create(transfer *models.Transfer) error {
	return r.db.Create(transfer).Error
}
