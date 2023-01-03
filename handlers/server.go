package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator"
	_ "github.com/rahul-024/fund-transfer-poc/docs"
	"github.com/rahul-024/fund-transfer-poc/util"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// Server serves HTTP requests for our banking service.
type Server struct {
	router *gin.Engine
}

func NewServer() (*Server, error) {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("currency", util.ValidCurrency)
	}

	server := &Server{}
	server.setupRouter()
	return server, nil
}

func (server *Server) setupRouter() {
	router := gin.Default()
	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	v1 := router.Group("/api/v1")
	{
		v1.POST("/accounts", CreateAccount)
		v1.GET("/accounts", ListAccounts)
		v1.GET("/accounts/:id", GetAccountById)
		v1.DELETE("/accounts/:id", DeleteAccountById)
		v1.PUT("/accounts/:id", UpdateAccountById)

	}
	server.router = router
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}
