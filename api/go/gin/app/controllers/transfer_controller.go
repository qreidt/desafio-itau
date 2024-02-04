package controllers

import (
	"github.com/gin-gonic/gin"
	"itau-api/app/exceptions"
	"itau-api/app/models"
	"itau-api/app/services"
	"net/http"
	"strconv"
)

type TransferController struct {
	transferService *services.TransferService
}

func NewTransferController(transferService *services.TransferService) *TransferController {
	return &TransferController{
		transferService: transferService,
	}
}

func (c *TransferController) Index(ctx *gin.Context) {
	user := ctx.MustGet("user").(models.User)

	var transfers []models.Transfer
	if err := c.transferService.ListUserTransfers(&transfers, user.ID); err != nil {
		exceptions.NewUnexpectedError(err, ctx)
		return
	}

	ctx.JSON(http.StatusOK, transfers)
}

func (c *TransferController) Store(ctx *gin.Context) {
	// Todo Start With Validation
}

func (c *TransferController) Show(ctx *gin.Context) {
	// Get the Id from the url
	var transactionId uint64
	if r, err := strconv.ParseInt(ctx.Param("transaction_id"), 10, 64); err != nil {
		transactionId = uint64(r)
	} else {
		exceptions.NewNotFoundException(ctx)
		return
	}

	// Find the transfer in the database
	var transfer models.Transfer
	if err := c.transferService.FindTransfer(&transfer, transactionId); err != nil {
		exceptions.NewNotFoundException(ctx)
		return
	}

	// Return
	ctx.JSON(200, transfer)
}

type TransferStoreRequest struct {
	SenderAccountId   uint64 `json:"sender_account_id" validate:"required,numeric"`
	ReceiverAccountId uint64 `json:"receiver_account_id" validate:"required,numeric"`
}
