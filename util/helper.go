package util

import (
	"fmt"

	"github.com/rahul-024/fund-transfer-poc/models"
)

func FormDsn(runtimeConfig *models.RuntimeConfig) string {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%v sslmode=disable",
		runtimeConfig.Datasource.Host, runtimeConfig.Datasource.UserName, runtimeConfig.Datasource.Password,
		runtimeConfig.Datasource.DbName, runtimeConfig.Datasource.Port)
	return dsn
}
