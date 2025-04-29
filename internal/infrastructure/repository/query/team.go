package query

import (
	"STUOJ/internal/infrastructure/repository/entity/field"
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
