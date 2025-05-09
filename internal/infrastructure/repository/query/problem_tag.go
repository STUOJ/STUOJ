package query

import (
	"STUOJ/internal/infrastructure/repository/entity/field"
)

var (
	ProblemTagAllField    = field.NewProblemTagField().SelectAll()
	ProblemTagSimpleField = field.NewProblemTagField().
				SelectProblemId().
				SelectTagId()
)
