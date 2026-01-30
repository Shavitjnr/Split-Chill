package validators

import (
	"github.com/go-playground/validator/v10"

	"github.com/Shavitjnr/split-chill-ai/pkg/core"
)


func ValidateFiscalYearStart(fl validator.FieldLevel) bool {
	date, ok := fl.Field().Interface().(core.FiscalYearStart)
	if !ok {
		return false
	}

	
	_, _, err := date.GetMonthDay()
	return err == nil
}
