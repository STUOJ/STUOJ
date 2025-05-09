package queryfield

import (
	"STUOJ/internal/infrastructure/persistence/entity/field"
)

var (
	LanguageAllField    = field.NewLanguageField().SelectAll()
	LanguageSimpleField = field.NewLanguageField().
				SelectId().
				SelectName().
				SelectStatus().
				SelectSerial()
)
