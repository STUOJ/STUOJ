package query

import "STUOJ/internal/db/entity/field"

var (
	TeamUserAllField    = field.NewTeamUserField().SelectAll()
	TeamUserSimpleField = field.NewTeamUserField().
				SelectTeamId().
				SelectUserId()
)
