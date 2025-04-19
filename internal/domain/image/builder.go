package image

import (
	"STUOJ/internal/domain/image/valueobject"
	"io"
)

type Option func(*Image)

func NewImage(options ...Option) Image {
	image := Image{}
	for _, option := range options {
		option(&image)
	}
	return image
}

func WithReader(reader io.Reader) Option {
	return func(i *Image) {
		i.Reader = reader
	}
}

func WithKey(key string) Option {
	return func(i *Image) {
		i.Key = key
	}
}

func WithAlbum(album uint8) Option {
	return func(i *Image) {
		i.Album = valueobject.Album(album)
	}
}

func WithUrl(url string) Option {
	return func(i *Image) {
		i.Url = url
	}
}
