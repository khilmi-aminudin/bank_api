package handlers

import (
	"mime/multipart"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/google/uuid"

	m "github.com/khilmi-aminudin/bank_api/repositories"
	"github.com/khilmi-aminudin/bank_api/services"
	"github.com/khilmi-aminudin/bank_api/utils"
)

type CustomerHandler interface {
	CreateCustomer(c *gin.Context)
	UpdateCustomer(c *gin.Context)
	GetAllCustomers(c *gin.Context)
	GetCustomerById(c *gin.Context)
}

type customerHandler struct {
	service services.CustomerService
	s3      utils.AWSS3Client
	config  utils.Config
}

func NewCustomerHandler(service services.CustomerService, config utils.Config) CustomerHandler {
	return &customerHandler{
		service: service,
		s3:      utils.NewAWSS3Client(config),
		config:  config,
	}
}

type customerCreateRequest struct {
	IDCardType   string `json:"id_card_type,omitempty"`
	IDCardNumber string `json:"id_card_number,omitempty"`
	IDCardFile   string `json:"id_card_file,omitempty"`
	FirstName    string `json:"first_name" binding:"required"`
	LastName     string `json:"last_name" binding:"required"`
	PhoneNumber  string `json:"phone_number" binding:"required,numeric,min=8,max=13"`
	Email        string `json:"email" binding:"required,email"`
	Username     string `json:"username" binding:"required,min=3"`
	Password     string `json:"password" binding:"required,min=8"`
}

// CreateCustomer implements CustomerHandler.
func (h *customerHandler) CreateCustomer(c *gin.Context) {
	var req customerCreateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(responseBadRequest(err.Error()))
		return
	}

	hashedPassword, _ := utils.HashPassword(req.Password)

	data, err := h.service.CreateCustomer(c, m.CreateCustomerParams{
		IDCardNumber: req.IDCardNumber,
		IDCardFile:   req.IDCardFile,
		FirstName:    req.FirstName,
		LastName:     req.LastName,
		PhoneNumber:  req.PhoneNumber,
		Email:        req.Email,
		Username:     req.Username,
		Password:     hashedPassword,
	})

	if err != nil {
		c.JSON(responseBadRequest(err.Error()))
		return
	}

	c.JSON(responseCreated("success created", gin.H{
		"username":     data.Username,
		"id_card_type": data.IDCardType,
		"status":       data.Status,
	}))
}

type listCustomerRequest struct {
	PageID   int32 `form:"page_id" binding:"required,min=1"`
	PageSize int32 `form:"page_size" binding:"required,min=5,max=100"`
}

// GetAllCustomers implements CustomerHandler.
func (h *customerHandler) GetAllCustomers(c *gin.Context) {
	var req listCustomerRequest
	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(responseBadRequest(err.Error()))
		return
	}

	data, err := h.service.GetAllCustomers(c, m.GetAllCustomersParams{
		Limit:  req.PageSize,
		Offset: (req.PageID - 1) * req.PageSize,
	})

	if err != nil {
		c.JSON(responseBadRequest(err.Error()))
		return
	}

	c.JSON(responseOK("success", data))
}

type getCustomer struct {
	ID string `json:"id" binding:"required"`
}

// GetCustomerById implements CustomerHandler.
func (h *customerHandler) GetCustomerById(c *gin.Context) {
	var req getCustomer
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(responseBadRequest(err.Error()))
		return
	}
	parsedID, err := uuid.Parse(req.ID)
	if err != nil {
		c.JSON(responseBadRequest(err.Error()))
		return
	}

	data, err := h.service.GetCustomerById(c, parsedID)
	if err != nil {
		c.JSON(responseBadRequest(err.Error()))
		return
	}

	c.JSON(responseOK("success", data))
}

type updateCustomerrequest struct {
	ID           string                `form:"id" binding:"required"`
	IDCardType   string                `form:"id_card_type" binding:"required"`
	IDCardNumber string                `form:"id_card_number" binding:"required"`
	File         *multipart.FileHeader `form:"file" binding:"required"`
}

// UpdateCustomer implements CustomerHandler.
func (h *customerHandler) UpdateCustomer(c *gin.Context) {
	var req updateCustomerrequest
	if err := c.ShouldBindWith(req, binding.FormMultipart); err != nil {
		c.JSON(responseBadRequest(err.Error()))
		return
	}

	idCustomer, err := uuid.Parse(req.ID)
	if err != nil {
		c.JSON(responseBadRequest("invalid id format"))
		return
	}

	uploadedFilename, err := h.s3.Upload(c, req.File, "id-cards")
	if err != nil {
		c.JSON(responseInternalServerError("error uploading id card"))
		return
	}
	arg := m.UpdateCustomerParams{
		ID:           idCustomer,
		IDCardType:   m.IDCardType(req.IDCardType),
		IDCardNumber: req.IDCardNumber,
		IDCardFile:   uploadedFilename,
	}

	data, err := h.service.UpdateCustomer(c, arg)

	if err != nil {
		c.JSON(responseBadRequest(err.Error()))
		return
	}

	c.JSON(responseCreated("success created", gin.H{
		"username":     data.Username,
		"id_card_type": data.IDCardType,
		"status":       data.Status,
	}))
}
