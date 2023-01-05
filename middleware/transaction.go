package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rahul-024/fund-transfer-poc/logger"
	"gorm.io/gorm"
)

// StatusInList -> checks if the given status is in the list
func StatusInList(status int, statusList []int) bool {
	for _, i := range statusList {
		if i == status {
			return true
		}
	}
	return false
}

// DBTransactionMiddleware : to setup the database transaction middleware
func DBTransactionMiddleware(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		txHandle := db.Begin()
		logger.Log.Info("beginning database transaction")

		defer func() {
			if r := recover(); r != nil {
				txHandle.Rollback()
			}
		}()

		c.Set("db_trx", txHandle)
		c.Next()

		if StatusInList(c.Writer.Status(), []int{http.StatusOK, http.StatusCreated}) {
			logger.Log.Info("committing transactions")
			if err := txHandle.Commit().Error; err != nil {
				logger.Log.Info("trx commit error: ", err)
			}
		} else {
			logger.Log.Info("rolling back transaction due to status code: ", c.Writer.Status())
			txHandle.Rollback()
		}
	}
}
