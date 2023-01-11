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

type AccountServiceImpl struct {
	accountRepository repository.AccountRepository
}

type AccountService interface {
	SaveAccount(models.Account) (models.Account, error)
	GetAll(pageId int, pageSize int) ([]models.Account, error)
	GetAccountById(id int) (models.Account, error)
	DeleteAccountById(id int) error
	UpdateAccountById(models.Account, models.Account) (models.Account, error)
	WithTrx(*gorm.DB) AccountServiceImpl
	SaveTransfer(req *request.TransferRequest) (models.Transfer, error)
	SaveEntry(req *request.TransferRequest, dc string) error
	IncrementBalance(int, float64) error
	DecrementBalance(int, float64) error
}

func NewAccountService(r repository.AccountRepository) AccountService {
	return AccountServiceImpl{
		accountRepository: r,
	}
}

// WithTrx enables repository with transaction
func (a AccountServiceImpl) WithTrx(trxHandle *gorm.DB) AccountServiceImpl {
	logger.Log.Info("In func() WithTrx :: SERVICE LAYER")
	a.accountRepository = a.accountRepository.WithTrx(trxHandle)
	return a
}

func (a AccountServiceImpl) SaveAccount(account models.Account) (models.Account, error) {
	logger.Log.Info("In func() SaveAccount :: SERVICE LAYER")
	return a.accountRepository.SaveAccount(account)
}

func (a AccountServiceImpl) GetAll(pageId int, pageSize int) ([]models.Account, error) {
	logger.Log.Info("In func() GetAll :: SERVICE LAYER")
	return a.accountRepository.GetAll(pageId, pageSize)
}

func (a AccountServiceImpl) GetAccountById(id int) (models.Account, error) {
	logger.Log.Info("In func() GetAccountById :: SERVICE LAYER")
	return a.accountRepository.GetAccountById(id)
}

func (a AccountServiceImpl) DeleteAccountById(id int) error {
	logger.Log.Info("In func() DeleteAccountById :: SERVICE LAYER")
	return a.accountRepository.DeleteAccountById(id)
}

func (a AccountServiceImpl) UpdateAccountById(originalAccount models.Account, changedAccount models.Account) (models.Account, error) {
	logger.Log.Info("In func() UpdateAccountById :: SERVICE LAYER")
	return a.accountRepository.UpdateAccountById(originalAccount, changedAccount)
}

func (a AccountServiceImpl) SaveTransfer(req *request.TransferRequest) (models.Transfer, error) {
	logger.Log.Info("In func() SaveTransfer :: SERVICE LAYER")
	transfer := &models.Transfer{}
	mapper.Mapper(req, transfer)
	return a.accountRepository.SaveTransfer(transfer)
}

func (a AccountServiceImpl) SaveEntry(req *request.TransferRequest, dc string) error {
	logger.Log.Info("In func() SaveEntry :: SERVICE LAYER")
	entry := &models.Entry{}
	if dc == "DEBIT" {
		(*entry).AccountID = req.FromAccountID
		(*entry).Amount = -req.Amount
	} else if dc == "CREDIT" {
		(*entry).AccountID = req.ToAccountID
		(*entry).Amount = req.Amount
	}
	err := a.accountRepository.SaveEntry(entry)
	return err
}

func (a AccountServiceImpl) IncrementBalance(receiver int, amount float64) error {
	logger.Log.Info("In func() IncrementBalance :: SERVICE LAYER")
	return a.accountRepository.IncrementBalance(receiver, amount)
}

func (a AccountServiceImpl) DecrementBalance(giver int, amount float64) error {
	logger.Log.Info("In func() DecrementBalance :: SERVICE LAYER")
	return a.accountRepository.DecrementBalance(giver, amount)
}
