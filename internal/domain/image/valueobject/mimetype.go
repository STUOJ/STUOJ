package valueobject

import "fmt"

type MimeType uint64

const (
	JPEG MimeType = 1
	PNG  MimeType = 2
	GIF  MimeType = 3
)

func (m MimeType) String() string {
	switch m {
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
	switch m {
	case JPEG, PNG, GIF:
		return nil
	default:
		return ErrInvalidMimeType
	}
}

var (
	ErrInvalidMimeType = fmt.Errorf("invalid mime type")
)
