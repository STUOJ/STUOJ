package query

import "STUOJ/internal/db/entity/field"

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
)
