package queryfield

import (
	"STUOJ/internal/infrastructure/persistence/entity/field"
)

var (
	TeamUserAllField    = field.NewTeamUserField().SelectAll()
	TeamUserSimpleField = field.NewTeamUserField().
				SelectTeamId().
				SelectUserId()
)
