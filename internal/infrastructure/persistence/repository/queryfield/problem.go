package queryfield

import (
	"STUOJ/internal/infrastructure/persistence/entity/field"
)

var (
	ProblemAllField    = field.NewProblemField().SelectAll()
	ProblemSimpleField = field.NewProblemField().
				SelectId().
				SelectTitle().
				SelectSource().
				SelectDifficulty()
	ProblemListItemField = field.NewProblemField().
				SelectId().
				SelectTitle().
				SelectSource().
				SelectStatus().
				SelectDifficulty().
				SelectCreateTime().
				SelectUpdateTime()
)
