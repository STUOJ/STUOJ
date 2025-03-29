package valueobject

import "fmt"

type Album uint8

const (
	Avatar  Album = 1
	Problem Album = 2
	Blog    Album = 3
)

func (a Album) String() string {
	switch a {
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
	switch a {
	case Avatar, Problem, Blog:
		return nil
	default:
		return ErrInvalidAlbum
	}
}

var (
	ErrInvalidAlbum = fmt.Errorf("invalid album")
)
