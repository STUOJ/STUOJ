package queryfield

import (
	"STUOJ/internal/infrastructure/persistence/entity/field"
)

var (
	TeamSubmissionAllField    = field.NewTeamSubmissionField().SelectAll()
	TeamSubmissionSimpleField = field.NewTeamSubmissionField().
					SelectTeamId().
					SelectSubmissionId()
)
