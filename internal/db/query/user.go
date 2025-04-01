package query

import "STUOJ/internal/db/entity/field"

var (
	UserAllField    = field.NewUserField().SelectAll()
	UserSimpleField = field.NewUserField().
			SelectId().
			SelectUsername().
			SelectAvatar().
			SelectRole()
)
