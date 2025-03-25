package query

import "STUOJ/internal/db/entity/field"

var (
	TeamSubmissionAllField    = field.NewTeamSubmissionField().SelectAll()
	TeamSubmissionSimpleField = field.NewTeamSubmissionField().
					SelectTeamId().
					SelectSubmissionId()
)
