package query

import "STUOJ/internal/db/entity/field"

var (
	SolutionAllField    = field.NewSolutionField().SelectAll()
	SolutionSimpleField = field.NewSolutionField().
				SelectId().
				SelectLanguageId().
				SelectProblemId().
				SelectSourceCode()
)
