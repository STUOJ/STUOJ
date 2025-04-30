package image

import (
	"STUOJ/internal/domain/image/valueobject"
	"STUOJ/internal/domain/image/yuki"
	"STUOJ/pkg/errors"
	"fmt"
	"io"
	"time"
)

type Image struct {
	Key        string
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

func (i Image) verify() error {
	if err := i.Album.Verify(); err != nil {
		return err
	}
	if i.Reader == nil {
		return fmt.Errorf("reader is nil")
	}
	return nil
}

func (i Image) Upload() (string, error) {
	if err := i.verify(); err != nil {
		return "", errors.ErrValidation.WithError(err)
	}
	url, err := handler.UploadImage(i.Reader, i.Key, i.Album)
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
