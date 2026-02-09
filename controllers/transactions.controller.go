package controllers

import (
	"net/http"
	"strings"
	"time"

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

// SearchTransactions godoc
// @Summary Search transactions by date or date range
// @Description Search transactions. Default is today. Format: YYYY/MM/DD or YYYY/MM/DD - YYYY/MM/DD
// @Tags Transactions
// @Produce json
// @Param date query string false "Date or date range (e.g. 2025/12/01 - 2026/02/05)"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} dtos.ErrorResponse
// @Failure 500 {object} dtos.ErrorResponse
// @Router /transactions/search [get]
func SearchTransactions(c *gin.Context) {
	dateParam := c.DefaultQuery("date", "")

	loc := time.Local
	var start time.Time
	var end time.Time
	var err error

	if dateParam == "" {
		now := time.Now().In(loc)
		start = time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, loc)
		end = start.AddDate(0, 0, 1).Add(-time.Nanosecond)
	} else {
		parts := strings.Split(dateParam, "-")
		if len(parts) == 1 {
			d := strings.TrimSpace(parts[0])
			parsed, perr := time.ParseInLocation("2006/01/02", d, loc)
			if perr != nil {
				// try dash-separated date
				parsed, perr = time.ParseInLocation("2006-01-02", d, loc)
			}
			if perr != nil {
				common.ThrowException(http.StatusBadRequest, "invalid date format, expected YYYY/MM/DD or YYYY-MM-DD")
				return
			}
			start = time.Date(parsed.Year(), parsed.Month(), parsed.Day(), 0, 0, 0, 0, loc)
			end = start.AddDate(0, 0, 1).Add(-time.Nanosecond)
		} else {
			s := strings.TrimSpace(parts[0])
			e := strings.TrimSpace(parts[1])
			sParsed, serr := time.ParseInLocation("2006/01/02", s, loc)
			if serr != nil {
				sParsed, serr = time.ParseInLocation("2006-01-02", s, loc)
			}
			eParsed, eerr := time.ParseInLocation("2006/01/02", e, loc)
			if eerr != nil {
				eParsed, eerr = time.ParseInLocation("2006-01-02", e, loc)
			}
			if serr != nil || eerr != nil {
				common.ThrowException(http.StatusBadRequest, "invalid date range format, expected YYYY/MM/DD - YYYY/MM/DD")
				return
			}
			start = time.Date(sParsed.Year(), sParsed.Month(), sParsed.Day(), 0, 0, 0, 0, loc)
			end = time.Date(eParsed.Year(), eParsed.Month(), eParsed.Day(), 23, 59, 59, int(time.Second-time.Nanosecond), loc)
			if end.Before(start) {
				common.ThrowException(http.StatusBadRequest, "end date must be after or equal to start date")
				return
			}
		}
	}

	// Call service to get transactions
	transactions, err := transactionService.SearchTransactionsByDateRange(start, end)
	if err != nil {
		common.ThrowException(http.StatusInternalServerError, err.Error())
		return
	}

	// Convert to DTOs
	respList := make([]dtos.TransactionResponse, len(transactions))
	for i, t := range transactions {
		details := make([]dtos.TransactionDetailResponse, len(t.Details))
		for j, d := range t.Details {
			details[j] = dtos.TransactionDetailResponse{
				ProductID:   d.ProductID,
				ProductName: d.ProductName,
				Quantity:    d.Quantity,
				Subtotal:    d.Subtotal,
				CreatedAt:   d.CreatedAt,
			}
		}

		respList[i] = dtos.TransactionResponse{
			TransactionID: t.ID,
			TotalAmount:   t.TotalAmount,
			Details:       details,
			CreatedAt:     t.CreatedAt,
		}
	}

	response := common.GetListSuccessResponse(respList)
	c.JSON(http.StatusOK, response)
}

// GetTransactionReport godoc
// @Summary Get selling report by date range
// @Description Generate a selling report with total revenue, transaction count, and best-selling products
// @Tags Transactions
// @Produce json
// @Param start_date query string true "Start date (format: YYYY-MM-DD, e.g. 2026-01-01)"
// @Param end_date query string true "End date (format: YYYY-MM-DD, e.g. 2026-02-01)"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} dtos.ErrorResponse
// @Failure 500 {object} dtos.ErrorResponse
// @Router /transactions/report [get]
func GetTransactionReport(c *gin.Context) {
	startDateParam := c.Query("start_date")
	endDateParam := c.Query("end_date")

	if startDateParam == "" || endDateParam == "" {
		common.ThrowException(http.StatusBadRequest, "start_date and end_date query parameters are required")
		return
	}

	loc := time.Local

	// Parse start date
	startParsed, startErr := time.ParseInLocation("2006-01-02", startDateParam, loc)
	if startErr != nil {
		common.ThrowException(http.StatusBadRequest, "invalid start_date format, expected YYYY-MM-DD")
		return
	}

	// Parse end date
	endParsed, endErr := time.ParseInLocation("2006-01-02", endDateParam, loc)
	if endErr != nil {
		common.ThrowException(http.StatusBadRequest, "invalid end_date format, expected YYYY-MM-DD")
		return
	}

	// Create time ranges
	start := time.Date(startParsed.Year(), startParsed.Month(), startParsed.Day(), 0, 0, 0, 0, loc)
	end := time.Date(endParsed.Year(), endParsed.Month(), endParsed.Day(), 23, 59, 59, int(time.Second-time.Nanosecond), loc)

	if end.Before(start) {
		common.ThrowException(http.StatusBadRequest, "end_date must be after or equal to start_date")
		return
	}

	// Call service to generate report
	report, err := transactionService.GenerateReport(start, end)
	if err != nil {
		common.ThrowException(http.StatusInternalServerError, err.Error())
		return
	}

	response := common.SuccessResponseWithData(report)
	c.JSON(http.StatusOK, response)
}
