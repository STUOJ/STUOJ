package valueobject

import (
	"STUOJ/internal/domain/shared"
	"fmt"
)

type MimeType struct {
	shared.Valueobject[int64]
}

const (
	JPEG int64 = 1
	PNG  int64 = 2
	GIF  int64 = 3
)

func (m MimeType) String() string {
	switch m.Value() {
	case JPEG:
		return "image/jpeg"
	case PNG:
		return "image/png"
	case GIF:
		return "image/gif"
	default:
		return "unknown"
	}
}

func (m MimeType) Verify() error {
	switch m.Value() {
	case JPEG, PNG, GIF:
		return nil
	default:
		return ErrInvalidMimeType
	}
}

func NewMimeType(mimetype int64) MimeType {
	var m MimeType
	m.Set(mimetype)
	return m
}

var (
	ErrInvalidMimeType = fmt.Errorf("invalid mime type")
)
