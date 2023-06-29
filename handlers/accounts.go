package handlers

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	"github.com/khilmi-aminudin/bank_api/middleware"
	m "github.com/khilmi-aminudin/bank_api/repositories"
	"github.com/khilmi-aminudin/bank_api/services"
)

type AccountHandler interface {
	CreateAccount(c *gin.Context)
	GetAccountByNumber(c *gin.Context)
}

type accountHandler struct {
	service         services.AccoountService
	customerservice services.CustomerService
}

func NewAccountHandler(service services.AccoountService, customerservice services.CustomerService) AccountHandler {
	return &accountHandler{
		service:         service,
		customerservice: customerservice,
	}
}

type createAccountRequest struct {
	Balance float64 `json:"balance" binding:"required,min=50000"`
}

// CreateAccount implements AccountHandler.
func (h *accountHandler) CreateAccount(c *gin.Context) {
	payload, err := middleware.GetPayload(c)
	if err != nil {
		c.JSON(responseBadRequest(err.Error()))
		return
	}

	cstData, err := h.customerservice.GetCustomerByUsername(c, payload.Username)
	if err != nil {
		c.JSON(responseNotFound(err.Error()))
		return
	}

	var req createAccountRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(responseBadRequest(err.Error()))
		return
	}

	account, err := h.service.CreateAccount(c, m.CreateAccountParams{
		CustomerID: cstData.ID,
		Balance:    req.Balance,
	})
	if err != nil {
		c.JSON(responseBadRequest(err.Error()))
		return
	}
	c.JSON(responseOK("success created", gin.H{
		"account_number":  account.Number,
		"account_balance": account.Balance,
	}))
}

type getAccountRequest struct {
	AccountNumber string `json:"account_number" binding:"required,min=50000"`
}

// GetAccountByNumber implements AccountHandler.
func (h *accountHandler) GetAccountByNumber(c *gin.Context) {
	payload, err := middleware.GetPayload(c)
	if err != nil {
		c.JSON(responseBadRequest(err.Error()))
		return
	}

	cstData, err := h.customerservice.GetCustomerByUsername(c, payload.Username)
	if err != nil {
		c.JSON(responseNotFound(err.Error()))
		return
	}

	var req getAccountRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(responseBadRequest(err.Error()))
		return
	}

	data, err := h.service.GetAccountByNumber(c, req.AccountNumber)
	if err != nil {
		c.JSON(responseBadRequest(err.Error()))
		return
	}

	if data.CustomerID != cstData.ID {
		c.JSON(responseOK("success", gin.H{
			"account_number": req.AccountNumber,
			"account_owner":  cstData.Username,
		}))
		return
	}

	rsp := gin.H{
		"account_info": struct {
			CustomerID uuid.UUID `json:"customer_id"`
			Number     string    `json:"account_number"`
			Balance    float64   `json:"balance"`
			CreatedAt  time.Time `json:"register_date"`
		}{
			CustomerID: data.CustomerID,
			Number:     data.Number,
			Balance:    data.Balance,
			CreatedAt:  data.CreatedAt,
		},
		"user_info": struct {
			Username string `json:"username"`
			Email    string `json:"email"`
			Status   string `json:"status"`
		}{
			Username: cstData.Username,
			Email:    cstData.Email,
			Status:   string(cstData.Status),
		},
	}
	c.JSON(responseOK("success", rsp))
}
