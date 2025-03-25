package query

import "STUOJ/internal/db/entity/field"

var (
	TestcaseAllField    = field.NewTestcaseField().SelectAll()
	TestcaseSimpleField = field.NewTestcaseField().
				SelectId().
				SelectProblemId().
				SelectSerial().
				SelectTestInput().
				SelectTestOutput()
)
