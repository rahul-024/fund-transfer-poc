package repository_test

import (
	"database/sql"
	"regexp"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/golang/mock/gomock"
	"github.com/rahul-024/fund-transfer-poc/logger"
	mockI "github.com/rahul-024/fund-transfer-poc/mocks"
	"github.com/rahul-024/fund-transfer-poc/models"
	"github.com/rahul-024/fund-transfer-poc/repository"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var gdb *gorm.DB
var mock sqlmock.Sqlmock

func mockDbConnection() (*gorm.DB, sqlmock.Sqlmock) {
	var db *sql.DB
	db, mock, _ := sqlmock.New()
	gdb, _ = gorm.Open(postgres.New(postgres.Config{Conn: db}), &gorm.Config{})
	return gdb, mock
}

func TestSaveAccount(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	mockLogger := mockI.NewMockLogger(mockCtrl)
	logger.SetLogger(mockLogger)
	mockLogger.EXPECT().Info("In func() SaveAccount :: REPO LAYER")

	account := models.Account{
		Currency:  "USD",
		Owner:     "John",
		Balance:   24.0,
		CreatedAt: time.Now(),
	}
	gdb, mock = mockDbConnection()
	accountRepositoryImpl := repository.NewAccountRepository(gdb)
	const sqlInsertAccount = `INSERT INTO "accounts" ("currency","owner","balance","created_at") 
						VALUES ($1,$2,$3,$4) RETURNING "id"`
	const newId = 1
	mock.ExpectBegin() // start transaction
	mock.ExpectQuery(regexp.QuoteMeta(sqlInsertAccount)).
		WithArgs(account.Currency, account.Owner, account.Balance, account.CreatedAt).
		WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(newId))
	mock.ExpectCommit() // commit transaction
	accountRepositoryImpl.SaveAccount(account)
	err := mock.ExpectationsWereMet()
	if err != nil {
		t.Errorf("Failed to meet expectations, got error: %v", err)
	}
}

func TestGetAll(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	mockLogger := mockI.NewMockLogger(mockCtrl)
	logger.SetLogger(mockLogger)
	mockLogger.EXPECT().Info("In func() GetAll :: REPO LAYER")
	gdb, mock = mockDbConnection()
	rows := sqlmock.
		NewRows([]string{"id", "currency", "owner", "balance", "created_at"}).
		AddRow(1, "USD", "John", 24, time.Now()).
		AddRow(2, "EUR", "Mike", 30, time.Now())

	accountRepositoryImpl := repository.NewAccountRepository(gdb)
	const sqlSelectFirst5 = `SELECT * FROM "accounts" LIMIT 5 OFFSET 1`
	mock.ExpectQuery(regexp.QuoteMeta(sqlSelectFirst5)).WillReturnRows(rows)
	accountRepositoryImpl.GetAll(1, 5)
	err := mock.ExpectationsWereMet()
	if err != nil {
		t.Errorf("Failed to meet expectations, got error: %v", err)
	}
}

func TestGetAccountById(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	mockLogger := mockI.NewMockLogger(mockCtrl)
	logger.SetLogger(mockLogger)
	mockLogger.EXPECT().Info("In func() GetAccountById :: REPO LAYER")
	gdb, mock = mockDbConnection()
	rows := sqlmock.
		NewRows([]string{"id", "currency", "owner", "balance", "created_at"}).
		AddRow(1, "USD", "John", 24, time.Now())

	accountRepositoryImpl := repository.NewAccountRepository(gdb)
	const sqlSelectByAccountId = `SELECT * FROM "accounts" WHERE id=$1 ORDER BY "accounts"."id" LIMIT 1`
	mock.ExpectQuery(regexp.QuoteMeta(sqlSelectByAccountId)).
		WithArgs(1).WillReturnRows(rows)
	accountRepositoryImpl.GetAccountById(1)
	err := mock.ExpectationsWereMet()
	if err != nil {
		t.Errorf("Failed to meet expectations, got error: %v", err)
	}
}

func TestDeleteAccountById(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	mockLogger := mockI.NewMockLogger(mockCtrl)
	logger.SetLogger(mockLogger)
	mockLogger.EXPECT().Info("In func() DeleteAccountById :: REPO LAYER")
	gdb, mock = mockDbConnection()
	accountRepositoryImpl := repository.NewAccountRepository(gdb)

	rows := sqlmock.
		NewRows([]string{"id", "currency", "owner", "balance", "created_at"}).
		AddRow(1, "USD", "John", 24, time.Now())
	const sqlSelectByAccountId = `SELECT * FROM "accounts" WHERE id=$1 ORDER BY "accounts"."id" LIMIT 1`
	mock.ExpectQuery(regexp.QuoteMeta(sqlSelectByAccountId)).
		WithArgs(1).WillReturnRows(rows)

	const sqlDeleteByAccountId = `DELETE FROM "accounts" WHERE "accounts"."id" = $1`
	mock.ExpectBegin()
	mock.ExpectExec(regexp.QuoteMeta(sqlDeleteByAccountId)).WithArgs(1).WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()
	accountRepositoryImpl.DeleteAccountById(1)

	err := mock.ExpectationsWereMet()
	if err != nil {
		t.Errorf("Failed to meet expectations, got error: %v", err)
	}
}

func TestUpdateAccountById(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	mockLogger := mockI.NewMockLogger(mockCtrl)
	logger.SetLogger(mockLogger)
	mockLogger.EXPECT().Info("In func() UpdateAccountById :: REPO LAYER")
	gdb, mock = mockDbConnection()
	accountRepositoryImpl := repository.NewAccountRepository(gdb)

	originalAccount := models.Account{
		Id:        1,
		Currency:  "USD",
		Owner:     "John",
		Balance:   10.0,
		CreatedAt: time.Date(2021, time.Month(2), 21, 1, 10, 30, 0, time.UTC),
	}
	changedAccount := models.Account{
		Currency:  "USD",
		Owner:     "John",
		Balance:   24.0,
		CreatedAt: time.Date(2021, time.Month(2), 21, 1, 10, 30, 0, time.UTC),
	}

	const sqlUpdateByAccountId = `UPDATE "accounts" SET "currency"=$1,"owner"=$2,"balance"=$3,"created_at"=$4 WHERE "id" = $5`
	mock.ExpectBegin()
	mock.ExpectExec(regexp.QuoteMeta(sqlUpdateByAccountId)).
		WithArgs(changedAccount.Currency, changedAccount.Owner, changedAccount.Balance, originalAccount.CreatedAt, 1).
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()
	accountRepositoryImpl.UpdateAccountById(originalAccount, changedAccount)

	err := mock.ExpectationsWereMet()
	if err != nil {
		t.Errorf("Failed to meet expectations, got error: %v", err)
	}
}

func TestSaveTransfer(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	mockLogger := mockI.NewMockLogger(mockCtrl)
	logger.SetLogger(mockLogger)
	mockLogger.EXPECT().Info("In func() SaveTransfer :: REPO LAYER")
	gdb, mock = mockDbConnection()
	accountRepositoryImpl := repository.NewAccountRepository(gdb)

	transfer := models.Transfer{
		FromAccountID: 1,
		ToAccountID:   2,
		Amount:        20.0,
		CreatedAt:     time.Now(),
	}

	const sqlInsertTransfer = `INSERT INTO "transfers" ("from_account_id","to_account_id","amount","created_at") 
						VALUES ($1,$2,$3,$4) RETURNING "id"`

	const newId = 1
	mock.ExpectBegin() // start transaction
	mock.ExpectQuery(regexp.QuoteMeta(sqlInsertTransfer)).
		WithArgs(transfer.FromAccountID, transfer.ToAccountID, transfer.Amount, transfer.CreatedAt).
		WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(newId))
	mock.ExpectCommit() // commit transaction
	accountRepositoryImpl.SaveTransfer(&transfer)
	err := mock.ExpectationsWereMet()
	if err != nil {
		t.Errorf("Failed to meet expectations, got error: %v", err)
	}
}

func TestSaveEntry(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	mockLogger := mockI.NewMockLogger(mockCtrl)
	logger.SetLogger(mockLogger)
	mockLogger.EXPECT().Info("In func() SaveEntry :: REPO LAYER")
	gdb, mock = mockDbConnection()
	accountRepositoryImpl := repository.NewAccountRepository(gdb)

	entry := models.Entry{
		AccountID: 1,
		Amount:    20.0,
		CreatedAt: time.Now(),
	}

	const sqlInsertEntry = `INSERT INTO "entries" ("account_id","amount","created_at") 
						VALUES ($1,$2,$3) RETURNING "id"`

	const newId = 1
	mock.ExpectBegin() // start transaction
	mock.ExpectQuery(regexp.QuoteMeta(sqlInsertEntry)).
		WithArgs(entry.AccountID, entry.Amount, entry.CreatedAt).
		WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(newId))
	mock.ExpectCommit() // commit transaction
	accountRepositoryImpl.SaveEntry(&entry)
	err := mock.ExpectationsWereMet()
	if err != nil {
		t.Errorf("Failed to meet expectations, got error: %v", err)
	}
}

func TestIncrementBalance(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	mockLogger := mockI.NewMockLogger(mockCtrl)
	logger.SetLogger(mockLogger)
	mockLogger.EXPECT().Info("In func() IncrementBalance :: REPO LAYER")
	gdb, mock = mockDbConnection()
	accountRepositoryImpl := repository.NewAccountRepository(gdb)

	account := models.Account{
		Id:        2,
		Currency:  "USD",
		Owner:     "John",
		Balance:   10.0,
		CreatedAt: time.Date(2021, time.Month(2), 21, 1, 10, 30, 0, time.UTC),
	}

	const sqlIncrementBalByAccountId = `UPDATE "accounts" SET "balance"=balance + $1 WHERE id=$2`
	mock.ExpectBegin() // start transaction
	mock.ExpectExec(regexp.QuoteMeta(sqlIncrementBalByAccountId)).
		WithArgs(14.0, account.Id).
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit() // commit transaction
	accountRepositoryImpl.IncrementBalance(2, 14.0)
	err := mock.ExpectationsWereMet()
	if err != nil {
		t.Errorf("Failed to meet expectations, got error: %v", err)
	}
}

func TestDecrementBalance(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	mockLogger := mockI.NewMockLogger(mockCtrl)
	logger.SetLogger(mockLogger)
	mockLogger.EXPECT().Info("In func() DecrementBalance :: REPO LAYER")
	gdb, mock = mockDbConnection()
	accountRepositoryImpl := repository.NewAccountRepository(gdb)

	account := models.Account{
		Id:        1,
		Currency:  "USD",
		Owner:     "John",
		Balance:   24.0,
		CreatedAt: time.Date(2021, time.Month(2), 21, 1, 10, 30, 0, time.UTC),
	}

	const sqlIncrementBalByAccountId = `UPDATE "accounts" SET "balance"=balance - $1 WHERE id=$2`
	mock.ExpectBegin() // start transaction
	mock.ExpectExec(regexp.QuoteMeta(sqlIncrementBalByAccountId)).
		WithArgs(10.0, account.Id).
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit() // commit transaction
	accountRepositoryImpl.DecrementBalance(1, 10.0)
	err := mock.ExpectationsWereMet()
	if err != nil {
		t.Errorf("Failed to meet expectations, got error: %v", err)
	}
}
