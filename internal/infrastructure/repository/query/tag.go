package query

import (
	"STUOJ/internal/infrastructure/repository/entity/field"
)

var (
	TagAllField    = field.NewTagField().SelectAll()
	TagSimpleField = field.NewTagField().
			SelectId().
			SelectName()
)
