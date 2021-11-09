package api

import (
    "github.com/go-playground/validator/v10"
    "github.com/sakhaei-wd/banker/util"
)

var validCurrency validator.Func = func(fieldLevel validator.FieldLevel) bool {
	//fieldLevel.field() get the value of the field
    if currency, ok := fieldLevel.Field().Interface().(string); ok {
        return util.IsSupportedCurrency(currency)
    }
    return false
}
