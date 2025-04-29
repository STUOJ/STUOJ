package query

import (
	"STUOJ/internal/infrastructure/repository/entity/field"
)

var (
	BlogAllField    = field.NewBlogField().SelectAll()
	BlogSimpleField = field.NewBlogField().
			SelectId().
			SelectTitle().
			SelectProblemId().
			SelectUserId()
)
