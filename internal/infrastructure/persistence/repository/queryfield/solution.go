package queryfield

import (
	"STUOJ/internal/infrastructure/persistence/entity/field"
)

var (
	SolutionAllField    = field.NewSolutionField().SelectAll()
	SolutionSimpleField = field.NewSolutionField().
				SelectId().
				SelectLanguageId().
				SelectProblemId().
				SelectSourceCode()
)
