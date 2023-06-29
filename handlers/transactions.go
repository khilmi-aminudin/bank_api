package handlers

import (
	"fmt"

	"github.com/gin-gonic/gin"

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
	merchantService services.MerchantService
}

func NewTransactionHandler(
	trxservice services.TransactionsService,
	accountService services.AccoountService,
	customerService services.CustomerService,
	merchantService services.MerchantService,
) TransactionHandler {
	return &transactionHandler{
		trxservice:      trxservice,
		accountService:  accountService,
		customerService: customerService,
		merchantService: merchantService,
	}
}

// GetTransactionHistory implements TransactionHandler.
func (h *transactionHandler) GetTransactionHistory(c *gin.Context) {
	payload, err := middleware.GetPayload(c)
	if err != nil {
		logger.Errorln(fmt.Sprintf("CALLED : middleware.GetPayload(c), Error: %v", err))
		c.JSON(responseBadRequest(err.Error()))
		return
	}

	cstData, err := h.customerService.GetCustomerByUsername(c, payload.Username)
	if err != nil {
		logger.Errorln(fmt.Sprintf("CALLED : h.customerService.GetCustomerByUsername(c, payload.Username), Error: %v", err))
		c.JSON(responseNotFound(err.Error()))
		return
	}

	fromAccount, err := h.accountService.GetAccountByCustomerId(c, cstData.ID)
	if err != nil {
		logger.Errorln(fmt.Sprintf("CALLED : h.accountService.GetAccountByCustomerId(c, cstData.ID), Error: %v", err))
		c.JSON(responseNotFound(err.Error()))
		return
	}

	data, err := h.trxservice.GetTransactionHistory(c, fromAccount.ID)
	if err != nil {
		logger.Errorln(fmt.Sprintf("CALLED : h.trxservice.GetTransactionHistory(c, fromAccount.ID), Error: %v", err))
		c.JSON(responseBadRequest(err.Error()))
		return
	}

	c.JSON(responseOK("success", data))
}

type getTransactionHistoryByTypeRequest struct {
	TransactionType string `uri:"type"`
}

// GetTransactionHistoryByType implements TransactionHandler.
func (h *transactionHandler) GetTransactionHistoryByType(c *gin.Context) {
	fmt.Println("CALLED")
	payload, err := middleware.GetPayload(c)
	if err != nil {
		logger.Errorln(fmt.Sprintf("CALLED : middleware.GetPayload(c), Error: %v", err))
		c.JSON(responseBadRequest(err.Error()))
		return
	}

	cstData, err := h.customerService.GetCustomerByUsername(c, payload.Username)
	if err != nil {
		logger.Errorln(fmt.Sprintf("CALLED : h.customerService.GetCustomerByUsername(c, payload.Username), Error: %v", err))
		c.JSON(responseNotFound(err.Error()))
		return
	}

	fromAccount, err := h.accountService.GetAccountByCustomerId(c, cstData.ID)
	if err != nil {
		logger.Errorln(fmt.Sprintf("CALLED : h.accountService.GetAccountByCustomerId(c, cstData.ID), Error: %v", err))
		c.JSON(responseNotFound(err.Error()))
		return
	}

	var req getTransactionHistoryByTypeRequest
	if err := c.ShouldBindUri(&req); err != nil {
		logger.Errorln(fmt.Sprintf("CALLED :  c.ShouldBindUri(&req), Error: %v", err))
		c.JSON(responseBadRequest(err.Error()))
		return
	}

	args := m.GetTransactionHistoryByTypeParams{
		TransactionType: m.TransactionType(req.TransactionType),
		FromAccountID:   fromAccount.ID,
	}

	data, err := h.trxservice.GetTransactionHistoryByType(c, args)
	if err != nil {
		logger.Errorln(fmt.Sprintf("CALLED : h.trxservice.GetTransactionHistoryByType(c, args), Error: %v", err))
		c.JSON(responseBadRequest(err.Error()))
		return
	}

	c.JSON(responseOK("success", data))
}

type createPaymentRequest struct {
	FromAccountID string  `json:"-"`
	ToMerchant    string  `json:"to_merchant" binding:"required"`
	Amount        float64 `json:"amount" binding:"required,min=1"`
	Description   string  `json:"description"`
}

// PaymentTx implements TransactionHandler.
func (h *transactionHandler) PaymentTx(c *gin.Context) {
	payload, err := middleware.GetPayload(c)
	if err != nil {
		logger.Errorln(fmt.Sprintf("CALLED : middleware.GetPayload(c), Error: %v", err))
		c.JSON(responseBadRequest(err.Error()))
		return
	}

	cstData, err := h.customerService.GetCustomerByUsername(c, payload.Username)
	if err != nil {
		logger.Errorln(fmt.Sprintf("CALLED : h.customerService.GetCustomerByUsername(c, payload.Username), Error: %v", err))
		c.JSON(responseNotFound(err.Error()))
		return
	}

	fromAccount, err := h.accountService.GetAccountByCustomerId(c, cstData.ID)
	if err != nil {
		logger.Errorln(fmt.Sprintf("CALLED : h.accountService.GetAccountByCustomerId(c, cstData.ID), Error: %v", err))
		c.JSON(responseNotFound(err.Error()))
		return
	}

	var req createPaymentRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		logger.Errorln(fmt.Sprintf("CALLED :  c.ShouldBindJSON(&req), Error: %v", err))
		c.JSON(responseBadRequest(err.Error()))
		return
	}

	if req.Amount > fromAccount.Balance {
		c.JSON(responseBadRequest("balance not enough"))
		return
	}

	merchant, err := h.merchantService.GetMerchantByName(c, req.ToMerchant)
	if err != nil {
		logger.Errorln(fmt.Sprintf("CALLED : h.merchantService.GetMerchantByName(c, req.ToMerchant), Error: %v", err))
		c.JSON(responseNotFound(err.Error()))
		return
	}

	args := m.PaymentTxParams{
		FromAccountID: fromAccount.ID,
		ToMerchantID:  merchant.ID,
		Amount:        req.Amount,
		Description:   req.Description,
	}

	data, err := h.trxservice.PaymentTx(c, args)
	if err != nil {
		logger.Errorln(fmt.Sprintf("CALLED : h.trxservice.PaymentTx(c, args), Error: %v", err))
		c.JSON(responseBadRequest(err.Error()))
		return
	}
	c.JSON(responseOK("success", "payment success, transaction id : ", data.TransactionID))
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
		logger.Errorln(fmt.Sprintf("CALLED : c.ShouldBindJSON(&req) , Error: %v", err))
		c.JSON(responseBadRequest(err.Error()))
		return
	}

	if req.Amount <= 0 {
		c.JSON(responseBadRequest("amount must be greater than zero"))
		return
	}

	account, err := h.accountService.GetAccountByNumber(c, fmt.Sprintf("%d", req.AccountNumber))
	if err != nil {
		logger.Errorln(fmt.Sprintf("CALLED : h.accountService.GetAccountByNumber(c, accountNumber), Error: %v", err))
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
		logger.Errorln(fmt.Sprintf("CALLED : h.trxservice.TopupTx(c, args), Error: %v", err))
		c.JSON(responseBadRequest(err.Error()))
		return
	}

	c.JSON(responseOK("success", fmt.Sprint("topup success, with transaction id: ", data.TransactionID)))
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
		logger.Errorln(fmt.Sprintf("CALLED : middleware.GetPayload(c), Error: %v", err))
		c.JSON(responseBadRequest(err.Error()))
		return
	}

	cstData, err := h.customerService.GetCustomerByUsername(c, payload.Username)
	if err != nil {
		logger.Errorln(fmt.Sprintf("CALLED : h.customerService.GetCustomerByUsername(c, payload.Username), Error: %v", err))
		c.JSON(responseNotFound(err.Error()))
		return
	}

	fromAccount, err := h.accountService.GetAccountByCustomerId(c, cstData.ID)
	if err != nil {
		logger.Errorln(fmt.Sprintf("CALLED : h.accountService.GetAccountByCustomerId(c, cstData.ID), Error: %v", err))
		c.JSON(responseBadRequest(err.Error()))
		return
	}

	var req createTransferRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		logger.Errorln(fmt.Sprintf("CALLED : c.ShouldBindJSON(&req) , Error: %v", err))
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
		logger.Errorln(fmt.Sprintf("CALLED : h.accountService.GetAccountByNumber(c, accountNumber), Error: %v", err))
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
		logger.Errorln(fmt.Sprintf("CALLED : h.trxservice.TransferTx(c, args), Error: %v", err))
		c.JSON(responseBadRequest(err.Error()))
		return
	}

	c.JSON(responseOK("success", fmt.Sprintf("transfer successfully transferred, transaction number : %v", data.TransactionID)))
}

type createWithdrawalRequest struct {
	Amount      float64 `json:"amount"`
	Description string  `json:"description"`
}

// WithdrawalTx implements TransactionHandler.
func (h *transactionHandler) WithdrawalTx(c *gin.Context) {
	payload, err := middleware.GetPayload(c)
	if err != nil {
		logger.Errorln(fmt.Sprintf("CALLED : middleware.GetPayload(c), Error: %v", err))
		c.JSON(responseBadRequest(err.Error()))
		return
	}

	cstData, err := h.customerService.GetCustomerByUsername(c, payload.Username)
	if err != nil {
		logger.Errorln(fmt.Sprintf("CALLED : h.customerService.GetCustomerByUsername(c, payload.Username), Error: %v", err))
		c.JSON(responseNotFound(err.Error()))
		return
	}

	account, err := h.accountService.GetAccountByCustomerId(c, cstData.ID)
	if err != nil {
		logger.Errorln(fmt.Sprintf("CALLED : h.accountService.GetAccountByCustomerId(c, cstData.ID), Error: %v", err))
		c.JSON(responseBadRequest(err.Error()))
		return
	}

	var req createWithdrawalRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		logger.Errorln(fmt.Sprintf("CALLED : c.ShouldBindJSON(&req) , Error: %v", err))
		c.JSON(responseBadRequest(err.Error()))
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
		logger.Errorln(fmt.Sprintf("CALLED : h.trxservice.WithdrawalTx(c, args) , Error: %v", err))
		c.JSON(responseBadRequest(err.Error()))
		return
	}

	c.JSON(responseOK("success", fmt.Sprintf("transfer successfully transferred, transaction number : %v", data.TransactionID)))
}
