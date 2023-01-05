package service

import (
	"github.com/rahul-024/fund-transfer-poc/logger"
	"github.com/rahul-024/fund-transfer-poc/models"
	"github.com/rahul-024/fund-transfer-poc/repository"
	"gorm.io/gorm"
)

type AccountService interface {
	Save(models.Account) (models.Account, error)
	GetAll(pageId int, pageSize int) ([]models.Account, error)
	GetAccountById(id int) (models.Account, error)
	DeleteAccountById(id int) error
	UpdateAccountById(models.Account, models.Account) (models.Account, error)
	WithTrx(*gorm.DB) accountService
	IncrementMoney(uint, float64) error
	DecrementMoney(uint, float64) error
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

func (a accountService) Save(user models.Account) (models.Account, error) {
	logger.Log.Info("In func() Save :: SERVICE LAYER")
	return a.accountRepository.Save(user)
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

func (a accountService) IncrementMoney(receiver uint, amount float64) error {

	return a.accountRepository.IncrementMoney(receiver, amount)
}

func (a accountService) DecrementMoney(giver uint, amount float64) error {

	return a.accountRepository.DecrementMoney(giver, amount)
}
