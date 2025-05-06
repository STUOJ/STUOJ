package valueobject

import (
	"STUOJ/internal/model"
	"errors"
)

type Message struct {
	model.Valueobject[string]
}

func NewMessage(content string) Message {
	var msg Message
	msg.Set(content)
	return msg
}

func (m Message) Verify() error {
	if len(m.Value()) > 65535 {
		return errors.New("message too long")
	}
	return nil
}
