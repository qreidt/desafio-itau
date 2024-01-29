package middlewares

import (
	"errors"
	"github.com/gin-gonic/gin"
	"itau-api/app/exceptions"
	"itau-api/app/models"
	"itau-api/app/repositories"
	"strings"
)

type AuthMiddleware struct {
	TokenRepository *repositories.TokenRepository
	UserRepository  *repositories.UserRepository
}

func NewAuthMiddeware(
	userRepository *repositories.UserRepository,
	tokenRepository *repositories.TokenRepository,
) *AuthMiddleware {
	return &AuthMiddleware{
		UserRepository:  userRepository,
		TokenRepository: tokenRepository,
	}
}

func (mid *AuthMiddleware) UseAuth(ctx *gin.Context) {

	// ex: ["Bearer", "er08ryw0r8yw9er8y239r8yq"]
	// ex: ["Bearer", "1|er08ryw0r8yw9er8y239r8yq"]
	bearer := strings.Split(ctx.GetHeader("Authorization"), " ")
	if len(bearer) != 2 {
		exceptions.NewUnauthorizedRequestBody(errors.New("token não autorizado"), ctx)
		return
	}

	if bearer[1] == "" {
		exceptions.NewUnauthorizedRequestBody(errors.New("token não autorizado"), ctx)
		return
	}

	tokenId, tokenString := getTokenFromBearer(strings.Split(bearer[1], "|"))
	if tokenString == "" && tokenId == "" {
		exceptions.NewUnauthorizedRequestBody(errors.New("token não autorizado"), ctx)
		return
	}

	var token models.ApiToken
	if err := mid.TokenRepository.FindByIdAndToken(&token, tokenId, tokenString); err != nil {
		exceptions.NewUnauthorizedRequestBody(errors.New("token não autorizado"), ctx)
		return
	}

	var user models.User
	if err := mid.UserRepository.FindById(&user, token.UserId); err != nil {
		exceptions.NewUnauthorizedRequestBody(errors.New("token não autorizado"), ctx)
		return
	}

	ctx.Set("user", user)
	ctx.Next()
}

func getTokenFromBearer(bearerToken []string) (string, string) {
	switch len(bearerToken) {
	case 1:
		return "", bearerToken[0]

	case 2:
		return bearerToken[0], bearerToken[1]

	default:
		return "", ""
	}
}

func GetAuthUser(ctx *gin.Context) models.User {
	return ctx.MustGet("user").(models.User)
}
