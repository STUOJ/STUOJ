package queryfield

import (
	"STUOJ/internal/infrastructure/persistence/entity/field"
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
