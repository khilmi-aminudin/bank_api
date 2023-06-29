package handlers

import (
	"fmt"
	"strings"
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
	customerService services.CustomerService
}

func NewAccountHandler(service services.AccoountService, customerService services.CustomerService) AccountHandler {
	return &accountHandler{
		service:         service,
		customerService: customerService,
	}
}

type createAccountRequest struct {
	Balance float64 `json:"balance" binding:"required,min=50000"`
}

// CreateAccount implements AccountHandler.
func (h *accountHandler) CreateAccount(c *gin.Context) {
	logger.Info("CALLED : CreateAccount(c *gin.Context)")
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

	if cstData.Status != m.CustomerEnumActive {
		c.JSON(responseForbidden("your user account is not active"))
		return
	}

	var req createAccountRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		logger.Errorln(fmt.Sprintf("CALLED : c.ShouldBindJSON(&req) , Error: %v", err))
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
	AccountNumber string `uri:"account_number" binding:"required"`
}

// GetAccountByNumber implements AccountHandler.
func (h *accountHandler) GetAccountByNumber(c *gin.Context) {
	logger.Info("CALLED : GetAccountByNumber(c *gin.Context)")
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

	var req getAccountRequest
	if err := c.ShouldBindUri(&req); err != nil {
		logger.Errorln(fmt.Sprintf("CALLED :  c.ShouldBindUri(&req), Error: %v", err))
		c.JSON(responseBadRequest(err.Error()))
		return
	}

	data, err := h.service.GetAccountByNumber(c, req.AccountNumber)
	if err != nil {
		logger.Errorln(fmt.Sprintf("CALLED : h.service.GetAccountByNumber(c, req.AccountNumber), Error: %v", err))
		c.JSON(responseBadRequest(err.Error()))
		return
	}

	if data.CustomerID != cstData.ID {
		cst, err := h.customerService.GetCustomerById(c, data.CustomerID)
		if err != nil {
			logger.Errorln(fmt.Sprintf("CALLED : h.customerService.GetCustomerById(c, data.CustomerID), Error: %v", err))
			c.JSON(responseBadRequest(err.Error()))
			return
		}

		c.JSON(responseOK("success", gin.H{
			"account_number": req.AccountNumber,
			"account_owner":  strings.ToUpper(fmt.Sprintf("%s %s", cst.FirstName, cst.LastName)),
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
