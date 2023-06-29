package handlers

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	"github.com/khilmi-aminudin/bank_api/middleware"
	m "github.com/khilmi-aminudin/bank_api/repositories"
	"github.com/khilmi-aminudin/bank_api/services"
)

type MerchantHandler interface {
	CreateMerchant(c *gin.Context)
	UpdateMerchant(c *gin.Context)
	GetAllMerchants(c *gin.Context)
	GetMerchantByName(c *gin.Context)
}

type merchantHandler struct {
	service services.MerchantService
}

func NewMerchantHandler(service services.MerchantService) MerchantHandler {
	return &merchantHandler{
		service: service,
	}
}

type createMerchantRequest struct {
	Name    string `json:"name" binding:"required"`
	Address string `json:"address" binding:"required"`
	Website string `json:"website" binding:"required"`
	Email   string `json:"email" binding:"required"`
}

// CreateMerchant implements MerchantHandler.
func (h *merchantHandler) CreateMerchant(c *gin.Context) {
	payload, err := middleware.GetPayload(c)
	if err != nil {
		c.JSON(responseBadRequest(err.Error()))
		return
	}

	if payload.Role != string(m.RoleAdmin) {
		c.JSON(responseUnauthorized("Unauthorized"))
		return
	}

	var req createMerchantRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(responseBadRequest(err.Error()))
		return
	}
	arg := m.CreateMerchantParams{
		Name:    req.Name,
		Address: req.Address,
		Website: req.Website,
		Email:   req.Email,
	}
	merchant, err := h.service.CreateMerchant(c, arg)
	if err != nil {
		c.JSON(responseBadRequest(err.Error()))
		return
	}

	data := struct {
		Name    string `json:"name"`
		Address string `json:"address"`
		Website string `json:"website"`
	}{
		Name:    merchant.Name,
		Address: merchant.Address,
		Website: merchant.Website,
	}

	c.JSON(responseCreated("success created", data))
}

// GetAllMerchants implements MerchantHandler.
func (h *merchantHandler) GetAllMerchants(c *gin.Context) {
	payload, err := middleware.GetPayload(c)
	if err != nil {
		c.JSON(responseBadRequest(err.Error()))
		return
	}

	data, err := h.service.GetAllMerchants(c)
	if err != nil {
		c.JSON(responseBadRequest(err.Error()))
		return
	}

	if payload.Role == string(m.RoleAdmin) {
		c.JSON(responseOK("success", data))
		return
	}

	type newMerchant struct {
		Name    string `json:"name"`
		Address string `json:"address"`
		Website string `json:"website"`
	}

	var newData []newMerchant

	for _, d := range data {
		nm := newMerchant{
			Name:    d.Name,
			Address: d.Address,
			Website: d.Website,
		}

		newData = append(newData, nm)
	}

	c.JSON(responseOK("success", newData))

}

type getMerchantByNameRequest struct {
	MerchantName string `uri:"merchant_name" binding:"required"`
}

// GetMerchantByName implements MerchantHandler.
func (h *merchantHandler) GetMerchantByName(c *gin.Context) {
	payload, err := middleware.GetPayload(c)
	if err != nil {
		c.JSON(responseBadRequest(err.Error()))
		return
	}

	var req getMerchantByNameRequest
	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(responseBadRequest(err.Error()))
		return
	}

	merchant, err := h.service.GetMerchantByName(c, req.MerchantName)
	if err != nil {
		c.JSON(responseBadRequest(err.Error()))
		return
	}

	if payload.Role == string(m.RoleAdmin) {
		fmt.Println("CALLED")
		c.JSON(responseOK("success", merchant))
		return
	}
	data := struct {
		Name    string `json:"name"`
		Address string `json:"address"`
		Website string `json:"website"`
	}{
		Name:    merchant.Name,
		Address: merchant.Address,
		Website: merchant.Website,
	}
	c.JSON(responseOK("success", data))

}

type updateMerchantRequest struct {
	ID      uuid.UUID `json:"-"`
	Name    string    `json:"name" binding:"required,min=3"`
	Address string    `json:"address" binding:"required,min=3"`
	Website string    `json:"website" binding:"required,min=3"`
}

// UpdateMerchant implements MerchantHandler.
func (h *merchantHandler) UpdateMerchant(c *gin.Context) {
	payload, err := middleware.GetPayload(c)
	if err != nil {
		c.JSON(responseBadRequest(err.Error()))
		return
	}

	if payload.Role != string(m.RoleAdmin) {
		c.JSON(responseUnauthorized("Unauthorized"))
		return
	}

	var req updateMerchantRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(responseBadRequest(err.Error()))
		return
	}

	merchant, err := h.service.GetMerchantByName(c, req.Name)
	if err != nil {
		c.JSON(responseNotFound(err.Error()))
		return
	}

	args := m.UpdateMerchantParams{
		ID:      merchant.ID,
		Name:    req.Name,
		Address: req.Address,
		Website: req.Website,
	}

	if err := h.service.UpdateMerchant(c, args); err != nil {
		c.JSON(responseBadRequest(err.Error()))
		return
	}

	c.JSON(responseOK("success", "merchant updated"))
}
