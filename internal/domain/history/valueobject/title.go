package valueobject

import (
	"STUOJ/internal/model"
)

type Title struct {
	model.Valueobject[string]
}

func NewTitle(title string) Title {
	var t Title
	t.Set(title)
	return t
}
