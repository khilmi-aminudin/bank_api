package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	m "github.com/khilmi-aminudin/bank_api/repositories"
	"github.com/khilmi-aminudin/bank_api/services"
)

type TransactionHandler interface {
	TransferTx(c *gin.Context)
	PaymentTx(c *gin.Context)
	WithdrawalTx(c *gin.Context)
	TopupTx(c *gin.Context)
	GetTransactionHistory(c *gin.Context)
	GetTransactionHistoryByType(c *gin.Context)
}

type transactionHandler struct {
	service services.TransactionsService
}

func NewTransactionHandler(service services.TransactionsService) TransactionHandler {
	return &transactionHandler{
		service: service,
	}
}

type accountId struct {
	AccountId string `json:"account_id"`
}

// GetTransactionHistory implements TransactionHandler.
func (h *transactionHandler) GetTransactionHistory(c *gin.Context) {
	var req accountId
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(responseBadRequest(err.Error()))
		return
	}
	var idParsed uuid.UUID
	var err error
	if idParsed, err = uuid.Parse(req.AccountId); err != nil {
		c.JSON(responseBadRequest(err.Error()))
		return
	}
	data, err := h.service.GetTransactionHistory(c, idParsed)

	if err != nil {
		c.JSON(responseBadRequest(err.Error()))
		return
	}

	c.JSON(responseOK("success", data))
}

type getTransactionHistoryByTypeRequest struct {
	TransactionType string `json:"transaction_type"`
	AccountId       string `json:"account_id"`
}

// GetTransactionHistoryByType implements TransactionHandler.
func (h *transactionHandler) GetTransactionHistoryByType(c *gin.Context) {
	var req getTransactionHistoryByTypeRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(responseBadRequest(err.Error()))
		return
	}

	fromAccountID, err := uuid.Parse(req.AccountId)
	if err != nil {
		c.JSON(responseBadRequest(err.Error()))
		return
	}

	args := m.GetTransactionHistoryByTypeParams{
		TransactionType: m.TransactionType(req.TransactionType),
		FromAccountID:   fromAccountID,
	}

	data, err := h.service.GetTransactionHistoryByType(c, args)
	if err != nil {
		c.JSON(responseBadRequest(err.Error()))
		return
	}

	c.JSON(responseOK("success", data))
}

type createPaymentRequest struct {
	FromAccountID string  `json:"from_account_id"`
	ToMerchantID  string  `json:"to_merchant_id"`
	Amount        float64 `json:"amount"`
	Description   string  `json:"description"`
}

// PaymentTx implements TransactionHandler.
func (h *transactionHandler) PaymentTx(c *gin.Context) {
	var req createPaymentRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(responseBadRequest(err.Error()))
		return
	}
	accountId, err := uuid.Parse(req.FromAccountID)
	if err != nil {
		c.JSON(responseBadRequest(err.Error()))
		return
	}

	merchantId, err := uuid.Parse(req.ToMerchantID)
	if err != nil {
		c.JSON(responseBadRequest(err.Error()))
		return
	}

	args := m.PaymentTxParams{
		FromAccountID: accountId,
		ToMerchantID:  merchantId,
		Amount:        req.Amount,
		Description:   req.Description,
	}

	data, err := h.service.PaymentTx(c, args)
	if err != nil {
		c.JSON(responseBadRequest(err.Error()))
		return
	}
	c.JSON(responseOK("success", data))
}

type createTopupRequest struct {
	ToAccountId string  `json:"to_account_id"`
	Amount      float64 `json:"amount"`
	Description string  `json:"description"`
}

// TopupTx implements TransactionHandler.
func (h *transactionHandler) TopupTx(c *gin.Context) {
	var req createTopupRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(responseBadRequest(err.Error()))
		return
	}

	accountId, err := uuid.Parse(req.ToAccountId)
	if err != nil {
		c.JSON(responseBadRequest(err.Error()))
		return
	}

	args := m.TopupParams{
		ToAccountId: accountId,
		Amount:      req.Amount,
		Description: req.Description,
	}

	data, err := h.service.TopupTx(c, args)
	if err != nil {
		c.JSON(responseBadRequest(err.Error()))
		return
	}

	c.JSON(responseOK("success", data))
}

type createTransferRequest struct {
	FromAccountID string  `json:"from_account_id"`
	ToAccountID   string  `json:"to_account_id"`
	Amount        float64 `json:"amount"`
	Description   string  `json:"description"`
}

// TransferTx implements TransactionHandler.
func (h *transactionHandler) TransferTx(c *gin.Context) {
	var req createTransferRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(responseBadRequest(err.Error()))
		return
	}

	fromAccountId, err := uuid.Parse(req.FromAccountID)
	if err != nil {
		c.JSON(responseBadRequest(err.Error()))
		return
	}

	toAccountId, err := uuid.Parse(req.ToAccountID)
	if err != nil {
		c.JSON(responseBadRequest(err.Error()))
		return
	}
	args := m.TranferTxParams{
		FromAccountID: fromAccountId,
		ToAccountID:   toAccountId,
		Amount:        req.Amount,
		Description:   req.Description,
	}

	data, err := h.service.TransferTx(c, args)
	if err != nil {
		c.JSON(responseBadRequest(err.Error()))
		return
	}

	c.JSON(responseOK("success", data))
}

type createWithdrawalRequest struct {
	FromAccountID string  `json:"from_account_id"`
	Amount        float64 `json:"amount"`
	Description   string  `json:"description"`
}

// WithdrawalTx implements TransactionHandler.
func (h *transactionHandler) WithdrawalTx(c *gin.Context) {
	var req createWithdrawalRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(responseBadRequest(err.Error()))
		return
	}
	fromAccountId, err := uuid.Parse(req.FromAccountID)
	if err != nil {
		c.JSON(responseBadRequest(err.Error()))
		return
	}

	args := m.WithdrawalParams{
		FromAccountID: fromAccountId,
		Amount:        req.Amount,
		Description:   req.Description,
	}

	data, err := h.service.WithdrawalTx(c, args)
	if err != nil {
		c.JSON(responseBadRequest(err.Error()))
		return
	}

	c.JSON(responseOK("success", data))
}
