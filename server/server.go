package server

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/khilmi-aminudin/bank_api/db"
	"github.com/khilmi-aminudin/bank_api/handlers"
	"github.com/khilmi-aminudin/bank_api/middleware"
	"github.com/khilmi-aminudin/bank_api/repositories"
	"github.com/khilmi-aminudin/bank_api/services"
	"github.com/khilmi-aminudin/bank_api/token"
	"github.com/khilmi-aminudin/bank_api/utils"
)

type Server struct {
	config     utils.Config
	tokenMaker token.Maker
	router     *gin.Engine
}

func NewServer(config utils.Config) (*Server, error) {
	tokenMaker, err := token.NewJWTMaker(config.TokenSymetricKey)
	if err != nil {
		return nil, fmt.Errorf("cannot create token maker: %v", err)
	}

	server := &Server{
		config:     config,
		tokenMaker: tokenMaker,
	}

	server.setupRouter()
	return server, nil
}

func (server *Server) setupRouter() {
	tMaker, _ := token.NewJWTMaker(server.config.TokenSymetricKey)
	dbConn := db.Connect(server.config)
	repository := repositories.NewRepo(dbConn)

	accountService := services.NewAccountService(repository)
	customerService := services.NewCustomerService(repository)

	accountHandler := handlers.NewAccountHandler(accountService, customerService)
	customerHandler := handlers.NewCustomerHandler(customerService, server.config)
	authHandler := handlers.NewAUthHandler(server.config, tMaker, customerService)

	// init router
	router := gin.Default()
	router.GET("/health", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "Hello I am okay",
		})
	})

	router.POST("/api/v1/register", customerHandler.CreateCustomer)
	router.POST("/api/v1/login", authHandler.Login)

	account := router.Group("/api/v1/accounts")
	{
		account.POST("/", accountHandler.CreateAccount)
		account.GET("/", accountHandler.GetAccountByNumber)
	}

	customer := router.Group("/api/v1/customers", middleware.AuthMiddleware(server.tokenMaker))
	{
		customer.PATCH("/", customerHandler.UpdateCustomer)
		customer.GET("/", customerHandler.GetAllCustomers)
		customer.GET("/:id", customerHandler.GetCustomerById)
	}

	server.router = router

}

func (server *Server) Start(addres string) error {
	return server.router.Run(addres)
}
