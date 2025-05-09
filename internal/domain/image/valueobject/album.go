package valueobject

import (
	"STUOJ/internal/model"
	"fmt"
)

type Album struct {
	model.Valueobject[uint8]
}

const (
	Avatar  uint8 = 1
	Problem uint8 = 2
	Blog    uint8 = 3
)

func (a Album) String() string {
	switch a.Value() {
	case Avatar:
		return "avatar"
	case Problem:
		return "problem"
	case Blog:
		return "blog"
	default:
		return "unknown"
	}
}

func (a Album) Verify() error {
	switch a.Value() {
	case Avatar, Problem, Blog:
		return nil
	default:
		return ErrInvalidAlbum
	}
}

func NewAlbum(album uint8) Album {
	var a Album
	a.Set(album)
	return a
}

var (
	ErrInvalidAlbum = fmt.Errorf("invalid album")
)
