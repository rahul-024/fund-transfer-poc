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
	const sqlInsert = `
					INSERT INTO "accounts" ("currency","owner","balance","created_at") 
						VALUES ($1,$2,$3,$4) RETURNING "id"`
	const newId = 1
	mock.ExpectBegin() // start transaction
	mock.ExpectQuery(regexp.QuoteMeta(sqlInsert)).
		WithArgs(account.Currency, account.Owner, account.Balance, account.CreatedAt).
		WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(newId))
	mock.ExpectCommit() // commit transaction
	accountRepositoryImpl.SaveAccount(account)
	err := mock.ExpectationsWereMet()
	if err != nil {
		t.Errorf("Failed to meet expectations, got error: %v", err)
	}
}
