package validators

import (
	"github.com/go-playground/validator/v10"

	"github.com/Shavitjnr/split-chill-ai/pkg/utils"
)


func ValidEmail(fl validator.FieldLevel) bool {
	if value, ok := fl.Field().Interface().(string); ok {
		if utils.IsValidEmail(value) {
			return true
		}
	}

	return false
}
