package controllers

import (
	"net/http"
	"strings"

	"kasir-api/common"
	"kasir-api/dtos"
	"kasir-api/services"

	"github.com/gin-gonic/gin"
)

var transactionService *services.TransactionService

// SetTransactionService initializes the transaction service with repository
func SetTransactionService(service *services.TransactionService) {
	transactionService = service
}

// Checkout godoc
// @Summary Complete transaction checkout
// @Description Process a checkout transaction with items and calculate total amount
// @Tags Transactions
// @Accept json
// @Produce json
// @Param request body dtos.CheckoutRequest true "Items to checkout"
// @Success 200 {object} dtos.CheckoutSuccessResponse
// @Failure 400 {object} dtos.ErrorResponse
// @Failure 500 {object} dtos.ErrorResponse
// @Router /transactions/checkout [post]
func Checkout(c *gin.Context) {
	var req dtos.CheckoutRequest

	// Bind and validate request body
	if err := c.ShouldBindJSON(&req); err != nil {
		errMsg := common.GenerateErrorMessage(err)
		common.ThrowExceptionOnValidation(http.StatusBadRequest, errMsg)
	}

	// Call service to process checkout
	transaction, err := transactionService.ProcessCheckout(req.Items)
	if err != nil {
		if strings.Contains(err.Error(), "not found") || strings.Contains(err.Error(), "insufficient") || strings.Contains(err.Error(), "no items") {
			common.ThrowException(http.StatusBadRequest, err.Error())
		}
		common.ThrowException(http.StatusInternalServerError, err.Error())
	}

	// Convert model to response DTO
	checkoutItems := make([]dtos.CheckoutItemResponse, 0)
	totalAmount := 0
	for _, detail := range transaction.Details {
		subtotal := detail.Subtotal
		checkoutItems = append(checkoutItems, dtos.CheckoutItemResponse{
			ProductID:   detail.ProductID,
			ProductName: detail.ProductName,
			Quantity:    detail.Quantity,
			Price:       subtotal / detail.Quantity,
			Subtotal:    subtotal,
		})
		totalAmount += subtotal
	}

	response := dtos.CheckoutSuccessResponse{
		Message: "Checkout completed successfully",
		Data: dtos.CheckoutResponse{
			TransactionID: transaction.ID,
			Items:         checkoutItems,
			TotalAmount:   totalAmount,
			CreatedAt:     transaction.CreatedAt,
		},
	}

	c.JSON(http.StatusOK, response)
}
