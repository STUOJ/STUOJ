package queryfield

import (
	"STUOJ/internal/infrastructure/persistence/entity/field"
)

var (
	HistoryAllField    = field.NewHistoryField().SelectAll()
	HistorySimpleField = field.NewHistoryField().
				SelectId().
				SelectUserId().
				SelectProblemId().
				SelectTitle().
				SelectDifficulty().
				SelectCreateTime().
				SelectOperation()
)
