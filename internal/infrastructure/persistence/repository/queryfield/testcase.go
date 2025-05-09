package queryfield

import (
	"STUOJ/internal/infrastructure/persistence/entity/field"
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
