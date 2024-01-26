package exceptions

import (
	"github.com/gin-gonic/gin"
)

func NewInvalidRequestBody(err error, ctx *gin.Context) {
	ctx.JSON(400, map[string]string{"error": err.Error()})
}

func NewUnexpectedError(err error, ctx *gin.Context) {
	ctx.JSON(500, map[string]string{"error": err.Error()})
}
