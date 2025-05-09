package queryfield

import (
	"STUOJ/internal/infrastructure/persistence/entity/field"
)

var (
	BlogAllField    = field.NewBlogField().SelectAll()
	BlogSimpleField = field.NewBlogField().
			SelectId().
			SelectTitle().
			SelectProblemId().
			SelectUserId()
)
