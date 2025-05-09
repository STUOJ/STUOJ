package query

import (
	"STUOJ/internal/infrastructure/repository/entity/field"
)

var (
	SubmissionAllField    = field.NewSubmissionField().SelectAll()
	SubmissionSimpleField = field.NewSubmissionField().
				SelectId().
				SelectUserId().
				SelectProblemId().
				SelectStatus().
				SelectScore().
				SelectMemory().
				SelectTime().
				SelectLength().
				SelectLanguageId().
				SelectCreateTime().
				SelectUpdateTime()
)
