package controllers

import (
	"github.com/gin-gonic/gin"
)

type TransactionController struct{}

func NewTransactionController() *TransactionController {
	return &TransactionController{}
}

func (c *TransactionController) Index(ctx *gin.Context) {

}

func (c *TransactionController) Store(ctx *gin.Context) {

}

func (c *TransactionController) Show(ctx *gin.Context) {

}
