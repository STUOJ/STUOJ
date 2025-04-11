package request

import "mime/multipart"

type UploadImageReq struct {
	Role string                `form:"role"`
	File *multipart.FileHeader `form:"file"`
}
