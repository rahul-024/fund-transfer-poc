package util

import (
	"github.com/go-playground/validator"
)

var ValidCurrency validator.Func = func(fieldLevel validator.FieldLevel) bool {
	if currency, ok := fieldLevel.Field().Interface().(string); ok {
		return IsSupportedCurrency(currency)
	}
	return false
}
