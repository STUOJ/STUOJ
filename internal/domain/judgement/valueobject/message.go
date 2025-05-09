package valueobject

import (
	"STUOJ/internal/domain/shared"
)

type Message struct {
	shared.Valueobject[string]
}

func NewMessage(content string) Message {
	var msg Message
	msg.Set(content)
	return msg
}
