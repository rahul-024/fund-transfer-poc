package repository

import (
	"errors"

	"github.com/rahul-024/fund-transfer-poc/logger"
	"github.com/rahul-024/fund-transfer-poc/models"
	"gorm.io/gorm"
)

type AccountRepositoryImpl struct {
	DB *gorm.DB
}

type AccountRepository interface {
	SaveAccount(models.Account) (models.Account, error)
	GetAll(pageId int, pageSize int) ([]models.Account, error)
	GetAccountById(id int) (models.Account, error)
	DeleteAccountById(id int) error
	UpdateAccountById(models.Account, models.Account) (models.Account, error)
	SaveTransfer(*models.Transfer) (models.Transfer, error)
	SaveEntry(*models.Entry) error
	IncrementBalance(int, float64) error
	DecrementBalance(int, float64) error
	WithTrx(*gorm.DB) AccountRepositoryImpl
}

func NewAccountRepository(db *gorm.DB) AccountRepository {
	return AccountRepositoryImpl{
		DB: db,
	}
}

func (a AccountRepositoryImpl) SaveAccount(account models.Account) (models.Account, error) {
	logger.Log.Info("In func() SaveAccount :: REPO LAYER")
	err := a.DB.Create(&account).Error
	return account, err
}

func (a AccountRepositoryImpl) GetAll(pageId int, pageSize int) (accounts []models.Account, err error) {
	logger.Log.Info("In func() GetAll :: REPO LAYER")
	err = a.DB.Limit(pageSize).Offset((pageId - 1) * pageSize).Find(&accounts).Error
	return accounts, err
}

func (a AccountRepositoryImpl) GetAccountById(id int) (account models.Account, err error) {
	logger.Log.Info("In func() GetAccountById :: REPO LAYER")
	err = a.DB.Where("id=?", id).First(&account).Error
	return account, err
}

func (a AccountRepositoryImpl) DeleteAccountById(id int) error {
	logger.Log.Info("In func() DeleteAccountById :: REPO LAYER")
	var account models.Account
	err := a.DB.Where("id=?", id).First(&account).Error
	if err != nil {
		return err
	}
	err = a.DB.Delete(&account).Error
	return err
}

func (a AccountRepositoryImpl) UpdateAccountById(originalAccount models.Account, changedAccount models.Account) (models.Account, error) {
	logger.Log.Info("In func() UpdateAccountById :: REPO LAYER")
	err := a.DB.Model(&originalAccount).Updates(&changedAccount).Error
	return changedAccount, err
}

func (a AccountRepositoryImpl) SaveTransfer(transfer *models.Transfer) (models.Transfer, error) {
	logger.Log.Info("In func() SaveTransfer :: REPO LAYER")
	err := a.DB.Create(&transfer).Error
	return *transfer, err
}

func (a AccountRepositoryImpl) SaveEntry(entry *models.Entry) error {
	logger.Log.Info("In func() SaveEntry :: REPO LAYER")
	err := a.DB.Create(&entry).Error
	return err
}

func (a AccountRepositoryImpl) IncrementBalance(receiver int, amount float64) error {
	logger.Log.Info("In func() IncrementBalance :: REPO LAYER")
	return a.DB.Model(&models.Account{}).Where("id=?", receiver).Update("balance", gorm.Expr("balance + ?", amount)).Error
}

func (a AccountRepositoryImpl) DecrementBalance(giver int, amount float64) error {
	logger.Log.Info("In func() DecrementBalance :: REPO LAYER")
	return errors.New("something")
	//return a.DB.Model(&models.Account{}).Where("id=?", giver).Update("balance", gorm.Expr("balance - ?", amount)).Error
}

func (a AccountRepositoryImpl) WithTrx(trxHandle *gorm.DB) AccountRepositoryImpl {
	logger.Log.Info("In func() WithTrx :: REPO LAYER")
	if trxHandle == nil {
		logger.Log.Info("Transaction Database not found")
		return a
	}
	a.DB = trxHandle
	return a
}
