package repositories

import (
	"fmt"
	"gorm.io/gorm"
	"itau-api/app/models"
	"math/rand"
)

type TokenRepository struct {
	db *gorm.DB
}

func NewTokenRepository(db *gorm.DB) *TokenRepository {
	return &TokenRepository{db: db}
}

const l = 62
const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

// Store

func (repo *TokenRepository) Create(userId uint64) (*models.ApiToken, string) {
	publicToken := randStringBytes(48)

	token := models.ApiToken{
		UserId:      userId,
		PublicToken: publicToken,
	}

	repo.db.Create(&token)

	publicToken = fmt.Sprintf("%d|%s", token.ID, publicToken)
	return &token, publicToken
}

func randStringBytes(n uint8) string {
	b := make([]byte, n)

	for i := range b {
		b[i] = letters[rand.Intn(l)]
	}

	return string(b)
}
