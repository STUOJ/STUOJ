package query

import (
	"STUOJ/internal/infrastructure/repository/entity/field"
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
