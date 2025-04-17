package image

import (
	"STUOJ/internal/app/dto/request"
	"STUOJ/internal/domain/image"
	"STUOJ/internal/errors"
	"STUOJ/internal/model"
)

func Insert(req request.UploadImageReq, reqUser model.ReqUser) (string, error) {
	ioReader, err := req.File.Open()
	if err != nil {
		return "", errors.ErrValidation.WithError(err)
	}
	defer ioReader.Close()
	image := image.NewImage(image.WithReader(ioReader))
	return image.Upload()
}
