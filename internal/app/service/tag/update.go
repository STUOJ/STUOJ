package tag

import (
	"STUOJ/internal/app/dto/request"
	"STUOJ/internal/domain/tag"
	"STUOJ/internal/model"
)

func Update(req request.UpdateTagReq, reqUser model.ReqUser) error {
	tag := tag.NewTag(
		tag.WithId(req.Id),
		tag.WithName(req.Name),
	)
	return tag.Update()
}
