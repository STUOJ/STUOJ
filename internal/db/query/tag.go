package query

import "STUOJ/internal/db/entity/field"

var (
	TagAllField    = field.NewTagField().SelectAll()
	TagSimpleField = field.NewTagField().
			SelectId().
			SelectName()
)
