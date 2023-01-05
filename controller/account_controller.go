package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/rahul-024/fund-transfer-poc/logger"
	"github.com/rahul-024/fund-transfer-poc/models"
	"github.com/rahul-024/fund-transfer-poc/service"
)

type AccountController interface {
	CreateAccount(*gin.Context)
	GetAccounts(*gin.Context)
	GetAccountById(*gin.Context)
	DeleteAccountById(*gin.Context)
	UpdateAccountById(*gin.Context)
	//TransferMoney(*gin.Context)
}

type accountController struct {
	accountService service.AccountService
}

type CreateAccountInput struct {
	Currency string  `json:"currency" binding:"required"`
	Owner    string  `json:"owner" binding:"required"`
	Balance  float64 `json:"balance"`
} // @name CreateAccountInput

type UpdateAccountInput struct {
	Currency string  `json:"currency"`
	Owner    string  `json:"owner"`
	Balance  float64 `json:"balance"`
} // @name UpdateAccountInput

type getAccountsRequest struct {
	PageID   int `form:"page_id" binding:"required,min=1"`
	PageSize int `form:"page_size" binding:"required,min=5,max=10"`
} // @name ListAccountRequest

func NewAccountController(s service.AccountService) AccountController {
	return accountController{
		accountService: s,
	}
}

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
func (a accountController) CreateAccount(c *gin.Context) {
	logger.Log.Info("In func() CreateAccount :: CONTROLLER LAYER")
	var input CreateAccountInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	account := models.Account{Currency: input.Currency, Owner: input.Owner}
	account, err := a.accountService.Save(account)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error while saving user"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": account})
}

// GetAccounts             godoc
//
//	@Summary		Get Accounts based on pageId and size
//	@Description	Responds with the list of all accounts as JSON.
//	@Tags			accounts
//	@Produce		json
//	@Param			page_id	query	int	true	"Provide the pageId from where the records needs to be returned"
//	@Param			page_size query	int	true	"provide the size of the page"
//	@Success		200	{object}	getAccountsRequest
//	@Failure		400	{string}	string	"Bad/Invalid request"
//	@Failure		500	{string}	string	"Resource not found"
//	@Failure		500	{string}	string	"Internal server error"
//	@Router			/accounts [get]
func (a accountController) GetAccounts(ctx *gin.Context) {
	logger.Log.Info("In func() GetAccounts :: CONTROLLER LAYER")
	var req getAccountsRequest
	var accounts []models.Account
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	accounts, error := a.accountService.GetAll(req.PageID, req.PageSize)
	if error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error occurred while fetching accounts"})
		return
	}
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
func (a accountController) GetAccountById(ctx *gin.Context) {
	logger.Log.Info("In func() GetAccountById :: CONTROLLER LAYER")
	var account models.Account
	intVar, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Path param is not an int"})
		return
	}
	account, err = a.accountService.GetAccountById(intVar)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
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
func (a accountController) DeleteAccountById(ctx *gin.Context) {
	logger.Log.Info("In func() DeleteAccountById :: CONTROLLER LAYER")
	intVar, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Path param is not an int"})
		return
	}
	err = a.accountService.DeleteAccountById(intVar)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
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
func (a accountController) UpdateAccountById(ctx *gin.Context) {
	logger.Log.Info("In func() UpdateAccountById :: CONTROLLER LAYER")
	intVar, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Path param is not an int"})
		return
	}
	account, err := a.accountService.GetAccountById(intVar)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Validate input
	var input UpdateAccountInput
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updatedAccount := models.Account{Currency: input.Currency, Owner: input.Owner, Balance: input.Balance}
	updatedAccount, err = a.accountService.UpdateAccountById(account, updatedAccount)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"data": updatedAccount})
}
