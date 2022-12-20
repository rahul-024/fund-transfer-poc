package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rahul-024/fund-transfer-poc/config"
	"github.com/rahul-024/fund-transfer-poc/models"
)

type createAccountRequest struct {
	Currency string `json:"currency" binding:"required,currency"`
	Owner    string `json:"owner" binding:"required"`
}

func (server *Server) createAccount(ctx *gin.Context) {
	var input createAccountRequest
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	account := models.Account{Currency: input.Currency, Owner: input.Owner}
	config.DB.Create(&account)

	ctx.JSON(http.StatusOK, account)
}
