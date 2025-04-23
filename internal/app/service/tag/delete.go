package tag

import (
	"STUOJ/internal/domain/tag"
	"STUOJ/internal/model"
)

func Delete(id int64, reqUser model.ReqUser) error {
	tag := tag.NewTag(
		tag.WithId(id),
	)
	return tag.Delete()
}
