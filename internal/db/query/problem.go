package query

import "STUOJ/internal/db/entity/field"

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
				SelectDifficulty().
				SelectCreateTime().
				SelectUpdateTime()
)
