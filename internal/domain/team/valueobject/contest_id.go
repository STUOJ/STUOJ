package valueobject

import (
	"STUOJ/internal/model"
	"errors"
)

type ContestId struct {
	model.Valueobject[int64]
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
