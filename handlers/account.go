package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	db "github.com/rahul-024/fund-transfer-poc/config"
	"github.com/rahul-024/fund-transfer-poc/logger"
	"github.com/rahul-024/fund-transfer-poc/models"
)

type CreateAccountInput struct {
	Currency string `json:"currency" binding:"required"`
	Owner    string `json:"owner" binding:"required"`
} // @name CreateAccountInput

type UpdateAccountInput struct {
	Currency string `json:"currency"`
	Owner    string `json:"owner"`
} // @name UpdateAccountInput

// PostAccount             godoc
//
//	@Summary		Create a new account
//	@Description	Takes a account JSON and store in DB. Return saved JSON.
//	@Tags			accounts
//	@Accept			json
//	@Produce		json
//	@Param			account	body		CreateAccountInput	true	"Account JSON"
//	@Success		201		{object}	CreateAccountInput
//	@Failure		400		{string}	string	"Bad/Invalid request"
//	@Failure		500		{string}	string	"Resource not found"
//	@Failure		500		{string}	string	"Internal server error"
//	@Router			/accounts [post]
func CreateAccount(ctx *gin.Context) {
	var input CreateAccountInput
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	account := models.Account{Currency: input.Currency, Owner: input.Owner}
	db.DB.Create(&account)

	ctx.JSON(http.StatusOK, account)
}

type listAccountRequest struct {
	PageID   int `form:"page_id" binding:"required,min=1"`
	PageSize int `form:"page_size" binding:"required,min=5,max=10"`
} // @name ListAccountRequest

// GetAccounts             godoc
//
//	@Summary		Get Accounts based on pageId and size
//	@Description	Responds with the list of all accounts as JSON.
//	@Tags			accounts
//	@Produce		json
//	@Param			page_id	query	int	true	"Provide the pageId from where the records needs to be returned"
//	@Param			page_size query	int	true	"provide the size of the page"
//	@Success		200	{object}	listAccountRequest
//	@Failure		400	{string}	string	"Bad/Invalid request"
//	@Failure		500	{string}	string	"Resource not found"
//	@Failure		500	{string}	string	"Internal server error"
//	@Router			/accounts [get]
func ListAccounts(ctx *gin.Context) {
	logger.Log.Info("In func ListAccounts()")
	var req listAccountRequest
	var accounts []models.Account
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	db.DB.Limit(req.PageSize).Offset((req.PageID - 1) * req.PageSize).Find(&accounts)
	ctx.JSON(http.StatusOK, accounts)
}

// GetAccountById             godoc
//
//	@Summary		Get single account by id
//	@Description	Returns the account whose id value matches the isbn.
//	@Tags			accounts
//	@Produce		json
//	@Param			id	path		int	true	"search account by id"
//	@Success		200	{object}	models.Account
//	@Failure		400	{string}	string	"Bad/Invalid request"
//	@Failure		500	{string}	string	"Resource not found"
//	@Failure		500	{string}	string	"Internal server error"
//	@Router			/accounts/{id} [get]
func GetAccountById(ctx *gin.Context) {
	var account models.Account
	if err := db.DB.Where("id=?", ctx.Param("id")).First(&account).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"data": account})
}

// DeleteAccountById             godoc
//
//	@Summary		Delete account by id
//	@Description	Delete an account with the given id
//	@Tags			accounts
//	@Produce		json
//	@Param			id	path		int	true	"delete account by id"
//	@Success		200	{string}	string
//	@Failure		400	{string}	string	"Bad/Invalid request"
//	@Failure		500	{string}	string	"Resource not found"
//	@Failure		500	{string}	string	"Internal server error"
//	@Router			/accounts/{id} [delete]
func DeleteAccountById(ctx *gin.Context) {
	var account models.Account
	if err := db.DB.Where("id=?", ctx.Param("id")).First(&account).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}
	db.DB.Delete(&account)
	ctx.JSON(http.StatusOK, gin.H{"data": "Account with id " + ctx.Param("id") + " deleted successfully"})
}

// UpdateAccountById             godoc
//
//		@Summary		Update account by id
//		@Description	Update an account with the given id
//		@Tags			accounts
//		@Produce		json
//	    @Consume		json
//		@Param			id	path		int	true	"update account by id"
//		@Success		200	{object}	models.Account
//		@Failure		400	{string}	string	"Bad/Invalid request"
//		@Failure		500	{string}	string	"Resource not found"
//		@Failure		500	{string}	string	"Internal server error"
//		@Router			/accounts/{id} [put]
func UpdateAccountById(ctx *gin.Context) {
	var account models.Account
	if err := db.DB.Where("id=?", ctx.Param("id")).First(&account).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	// Validate input
	var input UpdateAccountInput
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updateAccount := models.Account{Currency: input.Currency, Owner: input.Owner}
	db.DB.Model(&account).Updates(&updateAccount)
	ctx.JSON(http.StatusOK, gin.H{"data": account})
}
