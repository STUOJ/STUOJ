package valueobject

import (
	"STUOJ/internal/domain/shared"
	"errors"
)

type Key struct {
	shared.Valueobject[string]
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
