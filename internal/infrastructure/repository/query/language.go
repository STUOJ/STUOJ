package query

import (
	"STUOJ/internal/infrastructure/repository/entity/field"
)

var (
	LanguageAllField    = field.NewLanguageField().SelectAll()
	LanguageSimpleField = field.NewLanguageField().
				SelectId().
				SelectName().
				SelectStatus().
				SelectSerial()
)
