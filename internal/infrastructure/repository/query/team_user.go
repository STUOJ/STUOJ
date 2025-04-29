package query

import (
	"STUOJ/internal/infrastructure/repository/entity/field"
)

var (
	TeamUserAllField    = field.NewTeamUserField().SelectAll()
	TeamUserSimpleField = field.NewTeamUserField().
				SelectTeamId().
				SelectUserId()
)
