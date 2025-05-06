package valueobject

import (
	"STUOJ/internal/model"
	"errors"
)

type Stderr struct {
	model.Valueobject[string]
}

func NewStderr(content string) Stderr {
	var stderr Stderr
	stderr.Set(content)
	return stderr
}

func (s Stderr) Verify() error {
	if len(s.Value()) > 65535 {
		return errors.New("stderr 太长")
	}
	return nil
}
