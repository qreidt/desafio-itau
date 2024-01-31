package repositories

import (
	"fmt"
	"gorm.io/gorm"
	"itau-api/app/models"
	"itau-api/app/util"
)

type TokenRepository struct {
	db *gorm.DB
}

func NewTokenRepository(db *gorm.DB) *TokenRepository {
	return &TokenRepository{db: db}
}

// Store

func (repo *TokenRepository) Create(token *models.ApiToken, userId uint64) (string, error) {
	publicToken := util.RandStringBytes(48)

	token.UserId = userId
	token.PublicToken = publicToken

	if err := repo.db.Create(&token).Error; err != nil {
		return "", err
	}

	publicToken = fmt.Sprintf("%d|%s", token.ID, publicToken)
	return publicToken, nil
}

func (repo *TokenRepository) FindByIdAndToken(token *models.ApiToken, tokenId string, tokenString string) error {
	result := repo.db.First(&token, "id = ? and public_token = ?", tokenId, tokenString)
	return result.Error
}
