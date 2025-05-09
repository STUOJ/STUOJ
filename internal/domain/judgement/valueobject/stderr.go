package valueobject

import (
	"STUOJ/internal/domain/shared"
)

type Stderr struct {
	shared.Valueobject[string]
}

func NewStderr(content string) Stderr {
	var stderr Stderr
	stderr.Set(content)
	return stderr
}
