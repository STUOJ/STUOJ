package valueobject

import (
	"STUOJ/internal/model"
)

type Message struct {
	model.Valueobject[string]
}

func NewMessage(content string) Message {
	var msg Message
	msg.Set(content)
	return msg
}
