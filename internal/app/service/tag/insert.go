package tag

import (
	"STUOJ/internal/app/dto/request"
	"STUOJ/internal/domain/tag"
	"STUOJ/internal/model"
)

func Insert(req request.CreateTagReq, reqUser model.ReqUser) (int64, error) {
	tag := tag.NewTag(
		tag.WithName(req.Name),
	)
	return tag.Create()
}
