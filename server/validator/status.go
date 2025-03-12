package validator

import (
	"STUOJ/internal/model"

	"github.com/go-playground/validator/v10"
)

func StatusRangeValidator(fl validator.FieldLevel) bool {
	if status, ok := fl.Field().Interface().(model.ValidatableStatus); ok {
		return status.IsValid()
	}

	return false
}
