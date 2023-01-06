package config

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator"
	_ "github.com/rahul-024/fund-transfer-poc/docs"
	controller "github.com/rahul-024/fund-transfer-poc/handler"
	"github.com/rahul-024/fund-transfer-poc/middleware"
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
		accountHandler    = controller.NewAccountHandler(accountService)
	)

	accounts := router.Group("/api/v1/accounts")
	{
		accounts.POST("/", accountHandler.CreateAccount)
		accounts.GET("/", accountHandler.GetAccounts)
		accounts.GET("/:id", accountHandler.GetAccountById)
		accounts.DELETE("/:id", accountHandler.DeleteAccountById)
		accounts.PUT("/:id", accountHandler.UpdateAccountById)
	}

	transfers := router.Group("/api/v1/transfers")
	{
		transfers.POST("/", middleware.DBTransactionMiddleware(db), accountHandler.SaveTransfer)
	}
	server.router = router
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}
