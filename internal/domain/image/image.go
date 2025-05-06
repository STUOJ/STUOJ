package image

//go:generate go run ../../../dev/gen/builder.go image

import (
	"STUOJ/internal/domain/image/valueobject"
	"STUOJ/internal/domain/image/yuki"
	"STUOJ/pkg/errors"
	"io"
	"time"
)

type Image struct {
	Key        valueobject.Key
	Url        string
	Album      valueobject.Album
	Reader     io.Reader
	CreateTime time.Time
}

type ImageHandler interface {
	UploadImage(io.Reader, string, valueobject.Album) (string, error)
	DeleteImage(string) error
}

var handler ImageHandler = new(yuki.YukiImage)

func (i Image) Upload() (string, error) {
	if err := i.verify(); err != nil {
		return "", errors.ErrValidation.WithError(err)
	}
	url, err := handler.UploadImage(i.Reader, i.Key.Value(), i.Album)
	if err != nil {
		return "", errors.ErrInternalServer.WithError(err)
	}
	i.Url = url
	return url, nil
}

func (i Image) Delete() error {
	err := handler.DeleteImage(i.Url)
	if err != nil {
		return errors.ErrInternalServer.WithError(err)
	}
	return nil
}
