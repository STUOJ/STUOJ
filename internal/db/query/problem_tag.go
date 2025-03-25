package query

import "STUOJ/internal/db/entity/field"

var (
	ProblemTagAllField    = field.NewProblemTagField().SelectAll()
	ProblemTagSimpleField = field.NewProblemTagField().
				SelectProblemId().
				SelectTagId()
)
