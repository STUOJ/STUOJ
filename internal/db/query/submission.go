package query

import "STUOJ/internal/db/entity/field"

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
