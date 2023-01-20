package service_test

import (
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/rahul-024/fund-transfer-poc/logger"
	mock "github.com/rahul-024/fund-transfer-poc/mocks"

	"github.com/rahul-024/fund-transfer-poc/models"
	"github.com/rahul-024/fund-transfer-poc/models/request"
	"github.com/rahul-024/fund-transfer-poc/repository"
	"github.com/rahul-024/fund-transfer-poc/service"
	"gorm.io/gorm"
)

func TestSaveAccount(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	mockAccountRepo := mock.NewMockAccountRepository(mockCtrl)
	mockLogger := mock.NewMockLogger(mockCtrl)
	logger.SetLogger(mockLogger)
	mockLogger.EXPECT().Info("In func() SaveAccount :: SERVICE LAYER")
	account := models.Account{Currency: "USD", Owner: "rahul", Balance: 24}
	mockAccountRepo.EXPECT().SaveAccount(account).Return(models.Account{Currency: "USD", Owner: "rahul", Balance: 24}, nil).Times(1)
	accountServiceImpl := service.NewAccountService(mockAccountRepo)
	accountServiceImpl.SaveAccount(account)
}

func TestGetAll(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	mockAccountRepo := mock.NewMockAccountRepository(mockCtrl)
	mockLogger := mock.NewMockLogger(mockCtrl)
	logger.SetLogger(mockLogger)
	mockLogger.EXPECT().Info("In func() GetAll :: SERVICE LAYER")
	mockAccountRepo.EXPECT().GetAll(0, 5).Return(nil, nil).Times(1)
	accountServiceImpl := service.NewAccountService(mockAccountRepo)
	accountServiceImpl.GetAll(0, 5)
}

func TestGetAccountById(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	mockAccountRepo := mock.NewMockAccountRepository(mockCtrl)
	mockLogger := mock.NewMockLogger(mockCtrl)
	logger.SetLogger(mockLogger)
	mockLogger.EXPECT().Info("In func() GetAccountById :: SERVICE LAYER")
	mockAccountRepo.EXPECT().GetAccountById(1).Return(models.Account{Id: 1}, nil).Times(1)
	accountServiceImpl := service.NewAccountService(mockAccountRepo)
	accountServiceImpl.GetAccountById(1)
}

func TestDeleteAccountById(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	mockAccountRepo := mock.NewMockAccountRepository(mockCtrl)
	mockLogger := mock.NewMockLogger(mockCtrl)
	logger.SetLogger(mockLogger)
	mockLogger.EXPECT().Info("In func() DeleteAccountById :: SERVICE LAYER")
	mockAccountRepo.EXPECT().DeleteAccountById(1).Return(nil).Times(1)
	accountServiceImpl := service.NewAccountService(mockAccountRepo)
	accountServiceImpl.DeleteAccountById(1)
}

func TestUpdateAccountById(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	mockAccountRepo := mock.NewMockAccountRepository(mockCtrl)
	mockLogger := mock.NewMockLogger(mockCtrl)
	logger.SetLogger(mockLogger)
	mockLogger.EXPECT().Info("In func() UpdateAccountById :: SERVICE LAYER")
	originalAccount := models.Account{Id: 1, Currency: "USD", Owner: "rahul"}
	changedAccount := models.Account{Id: 1, Currency: "USD", Owner: "mike"}
	mockAccountRepo.EXPECT().UpdateAccountById(originalAccount, changedAccount).
		Return(models.Account{Id: 1, Currency: "USD", Owner: "mike"}, nil).Times(1)
	accountServiceImpl := service.NewAccountService(mockAccountRepo)
	accountServiceImpl.UpdateAccountById(originalAccount, changedAccount)
}

func TestWithTrx(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	mockAccountRepo := mock.NewMockAccountRepository(mockCtrl)
	mockLogger := mock.NewMockLogger(mockCtrl)
	logger.SetLogger(mockLogger)
	mockLogger.EXPECT().Info("In func() WithTrx :: SERVICE LAYER")
	mockAccountRepo.EXPECT().WithTrx(&gorm.DB{}).
		Return(repository.AccountRepositoryImpl{}).Times(1)
	accountServiceImpl := service.NewAccountService(mockAccountRepo)
	accountServiceImpl.WithTrx(&gorm.DB{})
}

func TestSaveTransfer(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	mockAccountRepo := mock.NewMockAccountRepository(mockCtrl)
	mockLogger := mock.NewMockLogger(mockCtrl)
	logger.SetLogger(mockLogger)
	mockLogger.EXPECT().Info("In func() SaveTransfer :: SERVICE LAYER")
	transferRequest := request.TransferRequest{FromAccountID: 1, ToAccountID: 2, Amount: 20, Currency: "USD"}
	transfer := &models.Transfer{Id: 0, FromAccountID: 1, ToAccountID: 2, Amount: 20, CreatedAt: time.Time{}}
	mockAccountRepo.EXPECT().SaveTransfer(transfer).
		Return(*transfer, nil).Times(1)
	accountServiceImpl := service.NewAccountService(mockAccountRepo)
	accountServiceImpl.SaveTransfer(&transferRequest)
}

func TestSaveEntry(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	mockAccountRepo := mock.NewMockAccountRepository(mockCtrl)
	mockLogger := mock.NewMockLogger(mockCtrl)
	logger.SetLogger(mockLogger)
	mockLogger.EXPECT().Info("In func() SaveEntry :: SERVICE LAYER")
	transferRequest := request.TransferRequest{FromAccountID: 1, ToAccountID: 2, Amount: 20, Currency: "USD"}
	entry := &models.Entry{Id: 0, AccountID: 1, Amount: -20}
	mockAccountRepo.EXPECT().SaveEntry(entry).
		Return(nil).Times(1)
	accountServiceImpl := service.NewAccountService(mockAccountRepo)
	accountServiceImpl.SaveEntry(&transferRequest, "DEBIT")

	//test CREDIT entry
	(*entry).Amount = 20
	(*entry).AccountID = 2
	mockLogger.EXPECT().Info("In func() SaveEntry :: SERVICE LAYER")
	mockAccountRepo.EXPECT().SaveEntry(entry).Return(nil).Times(1)
	accountServiceImpl = service.NewAccountService(mockAccountRepo)
	accountServiceImpl.SaveEntry(&transferRequest, "CREDIT")

}

func TestIncrementBalance(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	mockAccountRepo := mock.NewMockAccountRepository(mockCtrl)
	mockLogger := mock.NewMockLogger(mockCtrl)
	logger.SetLogger(mockLogger)
	mockLogger.EXPECT().Info("In func() IncrementBalance :: SERVICE LAYER")
	mockAccountRepo.EXPECT().IncrementBalance(1, 24.0).Return(nil).Times(1)
	accountServiceImpl := service.NewAccountService(mockAccountRepo)
	accountServiceImpl.IncrementBalance(1, 24.0)
}

func TestDecrementBalance(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	mockAccountRepo := mock.NewMockAccountRepository(mockCtrl)
	mockLogger := mock.NewMockLogger(mockCtrl)
	logger.SetLogger(mockLogger)
	mockLogger.EXPECT().Info("In func() DecrementBalance :: SERVICE LAYER")
	mockAccountRepo.EXPECT().DecrementBalance(1, 24.0).Return(nil).Times(1)
	accountServiceImpl := service.NewAccountService(mockAccountRepo)
	accountServiceImpl.DecrementBalance(1, 24.0)
}
