package config

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator"
	"github.com/rahul-024/fund-transfer-poc/controller"
	_ "github.com/rahul-024/fund-transfer-poc/docs"
	"github.com/rahul-024/fund-transfer-poc/repository"
	"github.com/rahul-024/fund-transfer-poc/service"
	"github.com/rahul-024/fund-transfer-poc/util"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"gorm.io/gorm"
)

// Server serves HTTP requests for our banking service.
type Server struct {
	router *gin.Engine
}

func NewServer(db *gorm.DB) (*Server, error) {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("currency", util.ValidCurrency)
	}

	server := &Server{}
	server.setupRouter(db)
	return server, nil
}

func (server *Server) setupRouter(db *gorm.DB) {
	router := gin.Default()
	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	var (
		accountRepository = repository.NewAccountRepository(db)
		accountService    = service.NewAccountService(accountRepository)
		accountController = controller.NewAccountController(accountService)
	)

	accounts := router.Group("/api/v1/accounts")
	{
		accounts.POST("/", accountController.CreateAccount)
		accounts.GET("/", accountController.GetAccounts)
		accounts.GET("/:id", accountController.GetAccountById)
		accounts.DELETE("/:id", accountController.DeleteAccountById)
		accounts.PUT("/:id", accountController.UpdateAccountById)
	}
	server.router = router
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}
