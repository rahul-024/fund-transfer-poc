package config

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator"
	_ "github.com/rahul-024/fund-transfer-poc/docs"
	"github.com/rahul-024/fund-transfer-poc/handlers"
	"github.com/rahul-024/fund-transfer-poc/util"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// Server serves HTTP requests for our banking service.
type Server struct {
	config util.ExtConfig
	router *gin.Engine
}

func NewServer(config util.ExtConfig) (*Server, error) {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("currency", util.ValidCurrency)
	}

	server := &Server{
		config: config,
	}

	server.setupRouter()
	return server, nil
}

func (server *Server) setupRouter() {
	router := gin.Default()
	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	v1 := router.Group("/api/v1")
	{
		v1.POST("/accounts", handlers.CreateAccount)

	}
	server.router = router
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}
