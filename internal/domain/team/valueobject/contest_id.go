package valueobject

import (
	"STUOJ/internal/domain/shared"
	"errors"
)

type ContestId struct {
	shared.Valueobject[int64]
}

func NewContestId(value int64) ContestId {
	var i ContestId
	i.Set(value)
	return i
}

func (i ContestId) Verify() error {
	if i.Value() <= 0 {
		return errors.New("contest id 不合法")
	}
	return nil
}
