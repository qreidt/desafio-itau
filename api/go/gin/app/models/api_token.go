package models

import (
	"time"
)

type ApiToken struct {
	ID          uint64    `json:"id"`
	UserId      uint64    `json:"user_id"`
	PublicToken string    `json:"token"`
	CreatedAt   time.Time `json:"created_at"`
}
