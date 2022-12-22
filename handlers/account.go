package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	db "github.com/rahul-024/fund-transfer-poc/db/config"
	"github.com/rahul-024/fund-transfer-poc/models"
)

type createAccountRequest struct {
	Currency string `json:"currency" binding:"required"`
	Owner    string `json:"owner" binding:"required"`
} //	@name	CreateAccountRequest

// PostAccount             godoc
//
//	@Summary		Create a new account
//	@Description	Takes a account JSON and store in DB. Return saved JSON.
//	@Tags			accounts
//	@Accept			json
//	@Produce		json
//	@Param			account	body		createAccountRequest	true	"Account JSON"
//	@Success		201		{object}	createAccountRequest
//
// @Failure 400 {string} string "Bad/Invalid request"
//
//	@Router			/accounts [post]
func CreateAccount(ctx *gin.Context) {
	var input createAccountRequest
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	account := models.Account{Currency: input.Currency, Owner: input.Owner}
	db.DB.Create(&account)

	ctx.JSON(http.StatusOK, account)
}
