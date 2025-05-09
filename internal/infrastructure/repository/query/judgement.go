package query

import (
	"STUOJ/internal/infrastructure/repository/entity/field"
)

var (
	JudgementAllField = field.NewJudgementField().SelectAll()
)
