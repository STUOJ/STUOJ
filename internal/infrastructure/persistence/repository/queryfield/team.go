package queryfield

import (
	"STUOJ/internal/infrastructure/persistence/entity/field"
)

var (
	TeamAllField    = field.NewTeamField().SelectAll()
	TeamSimpleField = field.NewTeamField().
			SelectId().
			SelectUserId().
			SelectContestId().
			SelectName().
			SelectStatus()
)
