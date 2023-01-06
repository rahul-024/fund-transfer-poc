package service

import (
	"github.com/devfeel/mapper"
	"github.com/rahul-024/fund-transfer-poc/logger"
	"github.com/rahul-024/fund-transfer-poc/models"
	"github.com/rahul-024/fund-transfer-poc/models/request"
	"github.com/rahul-024/fund-transfer-poc/repository"
	"gorm.io/gorm"
)

func init() {

	mapper.Register(&request.TransferRequest{})
	mapper.Register(&models.Transfer{})
}

type AccountService interface {
	SaveAccount(models.Account) (models.Account, error)
	GetAll(pageId int, pageSize int) ([]models.Account, error)
	GetAccountById(id int) (models.Account, error)
	DeleteAccountById(id int) error
	UpdateAccountById(models.Account, models.Account) (models.Account, error)
	WithTrx(*gorm.DB) accountService
	SaveTransfer(req *request.TransferRequest) (models.Transfer, error)
	SaveEntry(req *request.TransferRequest, dc string) error
	IncrementBalance(int, float64) error
	DecrementBalance(int, float64) error
}

type accountService struct {
	accountRepository repository.AccountRepository
}

func NewAccountService(r repository.AccountRepository) AccountService {
	return accountService{
		accountRepository: r,
	}
}

// WithTrx enables repository with transaction
func (a accountService) WithTrx(trxHandle *gorm.DB) accountService {
	a.accountRepository = a.accountRepository.WithTrx(trxHandle)
	return a
}

func (a accountService) SaveAccount(user models.Account) (models.Account, error) {
	logger.Log.Info("In func() Save :: SERVICE LAYER")
	return a.accountRepository.SaveAccount(user)
}

func (a accountService) GetAll(pageId int, pageSize int) ([]models.Account, error) {
	logger.Log.Info("In func() GetAll :: SERVICE LAYER")
	return a.accountRepository.GetAll(pageId, pageSize)
}

func (a accountService) GetAccountById(id int) (models.Account, error) {
	logger.Log.Info("In func() GetAccountById :: SERVICE LAYER")
	return a.accountRepository.GetAccountById(id)
}

func (a accountService) DeleteAccountById(id int) error {
	logger.Log.Info("In func() DeleteAccountById :: SERVICE LAYER")
	return a.accountRepository.DeleteAccountById(id)
}

func (a accountService) UpdateAccountById(originalAccount models.Account, changedAccount models.Account) (models.Account, error) {
	logger.Log.Info("In func() UpdateAccountById :: SERVICE LAYER")
	return a.accountRepository.UpdateAccountById(originalAccount, changedAccount)
}

func (a accountService) SaveTransfer(req *request.TransferRequest) (models.Transfer, error) {
	logger.Log.Info("In func() SaveTransfer :: SERVICE LAYER")
	transfer := &models.Transfer{}
	mapper.Mapper(req, transfer)
	return a.accountRepository.SaveTransfer(transfer)
}

func (a accountService) SaveEntry(req *request.TransferRequest, dc string) error {
	logger.Log.Info("In func() SaveEntry :: SERVICE LAYER")
	entry := &models.Entry{}
	if dc == "DEBIT" {
		(*entry).AccountID = req.FromAccountID
		(*entry).Amount = -req.Amount
	} else if dc == "CREDIT" {
		(*entry).AccountID = req.ToAccountID
		(*entry).Amount = -req.Amount
	}
	err := a.accountRepository.SaveEntry(entry)
	return err
}

func (a accountService) IncrementBalance(receiver int, amount float64) error {

	return a.accountRepository.IncrementBalance(receiver, amount)
}

func (a accountService) DecrementBalance(giver int, amount float64) error {

	return a.accountRepository.DecrementBalance(giver, amount)
}
