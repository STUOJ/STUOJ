package validator

import (
	"STUOJ/internal/domain/shared"
	"github.com/go-playground/validator/v10"
)

func StatusRangeValidator(fl validator.FieldLevel) bool {
	if status, ok := fl.Field().Interface().(shared.ValidatableStatus); ok {
		return status.IsValid()
	}

	return false
}
