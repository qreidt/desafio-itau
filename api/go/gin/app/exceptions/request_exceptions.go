package exceptions

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func NewInvalidRequestBody(err error, ctx *gin.Context) {
	ctx.JSON(400, map[string]string{"error": err.Error()})
}

func NewUnauthorizedRequestBody(err error, ctx *gin.Context) {
	ctx.JSON(401, map[string]string{"error": err.Error()})
}

func NewNotFoundException(ctx *gin.Context) {
	ctx.JSON(http.StatusNotFound, map[string]string{"error": "Not Found"})
	return
}

func NewUnexpectedError(err error, ctx *gin.Context) {
	ctx.JSON(500, map[string]string{"error": err.Error()})
}
