package yuki

import (
	"STUOJ/internal/domain/image/valueobject"
	"STUOJ/internal/infrastructure/yuki"
	"io"
)

type YukiImage struct{}

func (YukiImage) UploadImage(reader io.Reader, filename string, album valueobject.Album) (string, error) {
	image, err := yuki.UploadImage(reader, filename, album.Value())
	return image.Url, err
}

func (YukiImage) DeleteImage(url string) error {
	return yuki.DeleteImageByUrl(url)
}
