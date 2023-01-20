package handler_test

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/rahul-024/fund-transfer-poc/handler"
	"github.com/rahul-024/fund-transfer-poc/logger"
	mock "github.com/rahul-024/fund-transfer-poc/mocks"
	"github.com/rahul-024/fund-transfer-poc/models"
	"gopkg.in/go-playground/assert.v1"
)

func TestCreateAccount(t *testing.T) {

	mockCtrl := gomock.NewController(t)
	mockAccountService := mock.NewMockAccountService(mockCtrl)
	mockLogger := mock.NewMockLogger(mockCtrl)
	logger.SetLogger(mockLogger)
	//Success case
	mockLogger.EXPECT().Info("In func() CreateAccount :: HANDLER LAYER")
	recorder := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(recorder)
	jsonParam := `{"Currency":"USD","Owner":"rahul","Balance": 0.0}`
	req := httptest.NewRequest(http.MethodGet, "/", strings.NewReader(string(jsonParam)))
	c.Request = req
	account := models.Account{Currency: "USD", Owner: "rahul", Balance: 0.0}
	mockAccountService.EXPECT().SaveAccount(account).Return(models.Account{Currency: "USD", Owner: "rahul", Balance: 24}, nil).Times(1)
	accountHandlerImpl := handler.NewAccountHandler(mockAccountService)
	accountHandlerImpl.CreateAccount(c)

	//Failure case(1)
	jsonParam = `{}`
	recorder = httptest.NewRecorder()
	c, _ = gin.CreateTestContext(recorder)
	mockLogger.EXPECT().Info("In func() CreateAccount :: HANDLER LAYER")
	req = httptest.NewRequest(http.MethodGet, "/", strings.NewReader(string(jsonParam)))
	c.Request = req
	accountHandlerImpl.CreateAccount(c)
	assert.Equal(t, 400, recorder.Code)

	//Failure case(2)
	jsonParam = `{"Currency":"USD","Owner":"rahul","Balance": 0.0}`
	recorder = httptest.NewRecorder()
	c, _ = gin.CreateTestContext(recorder)
	mockLogger.EXPECT().Info("In func() CreateAccount :: HANDLER LAYER")
	req = httptest.NewRequest(http.MethodGet, "/", strings.NewReader(string(jsonParam)))
	c.Request = req
	mockAccountService.EXPECT().SaveAccount(account).
		Return(models.Account{Currency: "USD", Owner: "rahul", Balance: 24}, nil)
	accountHandlerImpl.CreateAccount(c)
	assert.Equal(t, 400, recorder.Code)
}
