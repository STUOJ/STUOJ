package query

import (
	"STUOJ/internal/infrastructure/repository/entity/field"
)

var (
	TestcaseAllField    = field.NewTestcaseField().SelectAll()
	TestcaseSimpleField = field.NewTestcaseField().
				SelectId().
				SelectProblemId().
				SelectSerial().
				SelectTestInput().
				SelectTestOutput()
)
