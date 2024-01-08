package api

import (
	"github.com/Silverpoision/simple_bank/util"
	"github.com/go-playground/validator/v10"
)

var validCurrency validator.Func = func(fieldLevel validator.FieldLevel) bool {
	if currency, ok := fieldLevel.Field().Interface().(string); ok {
		return util.ValidateCurrency(currency)
	}

	return false
}
