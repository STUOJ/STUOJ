package valueobject

import (
	"STUOJ/internal/model"
	"errors"
)

type Key struct {
	model.Valueobject[string]
}

func NewKey(key string) Key {
	var k Key
	k.Set(key)
	return k
}

func (k *Key) Verify() error {
	if len(k.Value()) == 0 {
		return errors.New("Key为空")
	}
	return nil
}
