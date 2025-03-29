package yuki

import (
	"STUOJ/external/yuki"
	"STUOJ/internal/domain/image/valueobject"
	"io"
)

type YukiImage struct{}

func (YukiImage) UploadImage(reader io.Reader, filename string, album valueobject.Album) (string, error) {
	image, err := yuki.UploadImage(reader, filename, uint8(album))
	return image.Url, err
}

func (YukiImage) DeleteImage(url string) error {
	return yuki.DeleteImageByUrl(url)
}
