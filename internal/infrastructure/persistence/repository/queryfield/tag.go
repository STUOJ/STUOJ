package queryfield

import (
	"STUOJ/internal/infrastructure/persistence/entity/field"
)

var (
	TagAllField    = field.NewTagField().SelectAll()
	TagSimpleField = field.NewTagField().
			SelectId().
			SelectName()
)
