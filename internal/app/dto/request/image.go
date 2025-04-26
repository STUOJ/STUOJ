package request

import "mime/multipart"

type UploadImageReq struct {
	Role int64                 `form:"role"`
	File *multipart.FileHeader `form:"file"`
}
