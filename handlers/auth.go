package handlers

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	m "github.com/khilmi-aminudin/bank_api/repositories"
	"github.com/khilmi-aminudin/bank_api/services"
	"github.com/khilmi-aminudin/bank_api/token"
	"github.com/khilmi-aminudin/bank_api/utils"
)

type AuthHandler interface {
	// RenewAccessToken(ctx *gin.Context)
	Login(ctx *gin.Context)
}

type authHandler struct {
	config      utils.Config
	tokenMaker  token.Maker
	customerSvc services.CustomerService
}

func NewAUthHandler(config utils.Config, tokenMaker token.Maker, customerSvc services.CustomerService) AuthHandler {
	return &authHandler{
		config:      config,
		tokenMaker:  tokenMaker,
		customerSvc: customerSvc,
	}
}

type loginRequest struct {
	Username string `json:"username" binding:"required,alphanum"`
	Password string `json:"password" binding:"required,min=6"`
}

type customerResponse struct {
	Role         m.Role         `json:"role"`
	Username     string         `json:"username"`
	RegisteredAt time.Time      `json:"registered_at"`
	Status       m.CustomerEnum `json:"status"`
}

type loginUserResponse struct {
	AccessToken           string           `json:"access_token"`
	AccessTokenExpiresAt  time.Time        `json:"access_token_expires_at"`
	RefreshToken          string           `json:"refresh_token"`
	RefreshTokenExpiresAt time.Time        `json:"refresh_token_expires_at"`
	Customer              customerResponse `json:"user"`
}

func newCustomerResponse(cst m.GetCustomerByUsernameRow) customerResponse {
	return customerResponse{
		Role:         cst.Role,
		Username:     cst.Username,
		RegisteredAt: cst.CreatedAt,
		Status:       cst.Status,
	}
}

// Login implements AuthHandler.
func (h *authHandler) Login(c *gin.Context) {
	var req loginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(responseBadRequest(err.Error()))
		return
	}

	cstData, err := h.customerSvc.GetCustomerByUsername(c, req.Username)
	if err != nil {
		c.JSON(responseNotFound(fmt.Sprintf("user %s not found", req.Username)))
		return
	}

	if cstData.Status != m.CustomerEnumActive && cstData.Status != m.CustomerEnumPending {
		c.JSON(responseForbidden(fmt.Sprintf("user %s is %s", req.Username, string(cstData.Status))))
		return
	}

	err = utils.CheckPassword(req.Password, cstData.Password)
	if err != nil {
		c.JSON(responseUnauthorized(err.Error()))
		return
	}
	accessToken, accessPayload, err := h.tokenMaker.CreateToken(
		req.Username,
		string(cstData.Role),
		h.config.AccessTokenDuration,
	)
	if err != nil {
		c.JSON(responseInternalServerError(err.Error()))
		return
	}
	refreshToken, refreshPayload, err := h.tokenMaker.CreateToken(
		req.Username,
		string(cstData.Role),
		h.config.RefreshTokenDuration,
	)
	if err != nil {
		c.JSON(responseInternalServerError(err.Error()))
		return
	}
	rsp := loginUserResponse{
		AccessToken:           accessToken,
		AccessTokenExpiresAt:  accessPayload.ExpiredAt,
		RefreshToken:          refreshToken,
		RefreshTokenExpiresAt: refreshPayload.ExpiredAt,
		Customer:              newCustomerResponse(cstData),
	}

	c.JSON(http.StatusOK, rsp)
}

// // RenewAccessToken implements AuthHandler.
// func (h *authHandler) RenewAccessToken(c *gin.Context) {
// 	panic("unimplemented")
// }
