package queryfield

import (
	"STUOJ/internal/infrastructure/persistence/entity/field"
)

var (
	ProblemTagAllField    = field.NewProblemTagField().SelectAll()
	ProblemTagSimpleField = field.NewProblemTagField().
				SelectProblemId().
				SelectTagId()
)
