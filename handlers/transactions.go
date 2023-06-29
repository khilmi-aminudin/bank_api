package handlers

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	"github.com/khilmi-aminudin/bank_api/middleware"
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
	trxservice      services.TransactionsService
	accountService  services.AccoountService
	customerService services.CustomerService
}

func NewTransactionHandler(trxservice services.TransactionsService, accountService services.AccoountService, customerService services.CustomerService) TransactionHandler {
	return &transactionHandler{
		trxservice:      trxservice,
		accountService:  accountService,
		customerService: customerService,
	}
}

type accountId struct {
	AccountId string `uri:"id" binding:"required"`
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
	data, err := h.trxservice.GetTransactionHistory(c, idParsed)

	if err != nil {
		c.JSON(responseBadRequest(err.Error()))
		return
	}

	c.JSON(responseOK("success", data))
}

type getTransactionHistoryByTypeRequest struct {
	TransactionType string `form:"transaction_type"`
	AccountId       string `form:"account_id"`
}

// GetTransactionHistoryByType implements TransactionHandler.
func (h *transactionHandler) GetTransactionHistoryByType(c *gin.Context) {
	var req getTransactionHistoryByTypeRequest
	if err := c.ShouldBindQuery(&req); err != nil {
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

	data, err := h.trxservice.GetTransactionHistoryByType(c, args)
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

	data, err := h.trxservice.PaymentTx(c, args)
	if err != nil {
		c.JSON(responseBadRequest(err.Error()))
		return
	}
	c.JSON(responseOK("success", data))
}

type createTopupRequest struct {
	AccountNumber int64   `json:"account_number"`
	Amount        float64 `json:"amount"`
	Description   string  `json:"description"`
}

// TopupTx implements TransactionHandler.
func (h *transactionHandler) TopupTx(c *gin.Context) {
	var req createTopupRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(responseBadRequest(err.Error()))
		return
	}

	if req.Amount <= 0 {
		c.JSON(responseBadRequest("amount must be greater than zero"))
		return
	}

	account, err := h.accountService.GetAccountByNumber(c, fmt.Sprintf("%d", req.AccountNumber))
	if err != nil {
		c.JSON(responseNotFound(err.Error()))
		return
	}

	args := m.TopupParams{
		ToAccountId: account.ID,
		Amount:      req.Amount,
		Description: req.Description,
	}

	data, err := h.trxservice.TopupTx(c, args)
	if err != nil {
		c.JSON(responseBadRequest(err.Error()))
		return
	}

	c.JSON(responseOK("success", data))
}

type createTransferRequest struct {
	ToAccountNumber int64   `json:"to_account_number"`
	Amount          float64 `json:"amount"`
	Description     string  `json:"description"`
}

// TransferTx implements TransactionHandler.
func (h *transactionHandler) TransferTx(c *gin.Context) {
	payload, err := middleware.GetPayload(c)
	if err != nil {
		c.JSON(responseBadRequest(err.Error()))
		return
	}

	cstData, err := h.customerService.GetCustomerByUsername(c, payload.Username)
	if err != nil {
		c.JSON(responseNotFound(err.Error()))
		return
	}

	fromAccount, err := h.accountService.GetAccountByCustomerId(c, cstData.ID)
	if err != nil {
		c.JSON(responseBadRequest(err.Error()))
		return
	}

	var req createTransferRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(responseBadRequest(err.Error()))
		return
	}

	if req.Amount > fromAccount.Balance {
		c.JSON(responseBadRequest("balance not enough"))
		return
	}

	if req.Amount <= 0 {
		c.JSON(responseBadRequest("invalid amount"))
		return
	}

	toAccount, err := h.accountService.GetAccountByNumber(c, fmt.Sprintf("%d", req.ToAccountNumber))
	if err != nil {
		c.JSON(responseBadRequest(err.Error()))
		return
	}

	args := m.TranferTxParams{
		FromAccountID: fromAccount.ID,
		ToAccountID:   toAccount.ID,
		Amount:        req.Amount,
		Description:   req.Description,
	}

	data, err := h.trxservice.TransferTx(c, args)
	if err != nil {
		c.JSON(responseBadRequest(err.Error()))
		return
	}

	c.JSON(responseOK("success", fmt.Sprintf("transfer successfully transferred, transaction number : %v", data.TransactionID)))
}

type createWithdrawalRequest struct {
	AccountNumber int64   `json:"account_number"`
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

	account, err := h.accountService.GetAccountByNumber(c, fmt.Sprintf("%d", req.AccountNumber))
	if err != nil {
		c.JSON(responseNotFound(err.Error()))
		return
	}

	if req.Amount > account.Balance {
		c.JSON(responseBadRequest("balance not enough"))
		return
	}

	if req.Amount <= 0 {
		c.JSON(responseBadRequest("invalid amount"))
		return
	}

	args := m.WithdrawalParams{
		FromAccountID: account.ID,
		Amount:        req.Amount,
		Description:   req.Description,
	}

	data, err := h.trxservice.WithdrawalTx(c, args)
	if err != nil {
		c.JSON(responseBadRequest(err.Error()))
		return
	}

	c.JSON(responseOK("success", data))
}
