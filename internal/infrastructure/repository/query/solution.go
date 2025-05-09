package query

import (
	"STUOJ/internal/infrastructure/repository/entity/field"
)

var (
	SolutionAllField    = field.NewSolutionField().SelectAll()
	SolutionSimpleField = field.NewSolutionField().
				SelectId().
				SelectLanguageId().
				SelectProblemId().
				SelectSourceCode()
)
