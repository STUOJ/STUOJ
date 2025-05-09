package valueobject

import (
	"STUOJ/internal/model"
)

type Stderr struct {
	model.Valueobject[string]
}

func NewStderr(content string) Stderr {
	var stderr Stderr
	stderr.Set(content)
	return stderr
}
