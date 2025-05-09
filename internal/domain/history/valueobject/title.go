package valueobject

import (
	"STUOJ/internal/domain/shared"
)

type Title struct {
	shared.Valueobject[string]
}

func NewTitle(title string) Title {
	var t Title
	t.Set(title)
	return t
}
