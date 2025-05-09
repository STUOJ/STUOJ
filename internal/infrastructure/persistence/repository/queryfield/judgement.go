package queryfield

import (
	"STUOJ/internal/infrastructure/persistence/entity/field"
)

var (
	JudgementAllField = field.NewJudgementField().SelectAll()
)
