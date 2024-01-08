package models

import (
	"time"
)

type User struct {
	ID        uint64
	Name      string
	Document  string
	Password  string
	Balance   int64
	CreatedAt time.Time
	UpdatedAt time.Time
}
