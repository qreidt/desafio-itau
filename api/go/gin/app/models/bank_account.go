package models

import "time"

type BankAccount struct {
	ID        uint64    `json:"id"`
	UserId    uint64    `json:"user_id"`
	Number    string    `json:"number"`
	Balance   int64     `json:"balance"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
