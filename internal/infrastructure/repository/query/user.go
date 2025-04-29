package query

import (
	"STUOJ/internal/infrastructure/repository/entity/field"
)

var (
	UserAllField    = field.NewUserField().SelectAll()
	UserSimpleField = field.NewUserField().
			SelectId().
			SelectUsername().
			SelectAvatar().
			SelectRole()
)
