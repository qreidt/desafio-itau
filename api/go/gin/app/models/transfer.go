package models

import (
	"time"
)

type Transfer struct {
	ID                uint64    `json:"id"`
	SenderUserId      uint64    `json:"sender_user_id"`
	ReceiverUserId    uint64    `json:"receiver_user_id"`
	SenderAccountId   uint64    `json:"sender_account_id"`
	ReceiverAccountId uint64    `json:"receiver_account_id"`
	Type              string    `json:"type"`
	Value             uint      `json:"value"`
	CreatedAt         time.Time `json:"created_at"`

	SenderUser      User        `json:"sender_user" gorm:"foreignKey:sender_user_id"`
	ReceiverUser    User        `json:"receiver_user" gorm:"foreignKey:receiver_user_id"`
	SenderAccount   BankAccount `json:"sender_account" gorm:"foreignKey:sender_account_id"`
	ReceiverAccount BankAccount `json:"receiver_account" gorm:"foreignKey:receiver_account_id"`
}
