package query

import (
	"STUOJ/internal/infrastructure/repository/entity/field"
)

var (
	ContestAllField    = field.NewContestField().SelectAll()
	ContestSimpleField = field.NewContestField().
				SelectId().
				SelectUserId().
				SelectTitle().
				SelectStatus().
				SelectFormat().
				SelectTeamSize().
				SelectStartTime().
				SelectEndTime()
	ContestListItemField = field.NewContestField().
				SelectId().
				SelectTitle().
				SelectStatus().
				SelectFormat().
				SelectTeamSize().
				SelectStartTime().
				SelectEndTime()
)
