package query

import (
	"STUOJ/internal/infrastructure/repository/entity/field"
)

var (
	TeamSubmissionAllField    = field.NewTeamSubmissionField().SelectAll()
	TeamSubmissionSimpleField = field.NewTeamSubmissionField().
					SelectTeamId().
					SelectSubmissionId()
)
