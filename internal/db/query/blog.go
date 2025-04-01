package query

import (
	"STUOJ/internal/db/entity/field"
)

var (
	BlogAllField    = field.NewBlogField().SelectAll()
	BlogSimpleField = field.NewBlogField().
			SelectId().
			SelectTitle().
			SelectProblemId().
			SelectUserId()
)
