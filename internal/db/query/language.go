package query

import "STUOJ/internal/db/entity/field"

var (
	LanguageAllField    = field.NewLanguageField().SelectAll()
	LanguageSimpleField = field.NewLanguageField().
				SelectId().
				SelectName().
				SelectStatus().
				SelectSerial()
)
