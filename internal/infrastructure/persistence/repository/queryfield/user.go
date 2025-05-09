package queryfield

import (
	"STUOJ/internal/infrastructure/persistence/entity/field"
)

var (
	UserAllField    = field.NewUserField().SelectAll()
	UserSimpleField = field.NewUserField().
			SelectId().
			SelectUsername().
			SelectAvatar().
			SelectRole()
)
