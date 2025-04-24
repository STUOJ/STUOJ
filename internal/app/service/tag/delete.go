package tag

import (
	"STUOJ/internal/domain/tag"
	"STUOJ/internal/model"
)

func Delete(id int64, reqUser model.ReqUser) error {
	// 检查权限
	err := isPermission(reqUser)
	if err != nil {
		return err
	}

	tag := tag.NewTag(
		tag.WithId(id),
	)
	return tag.Delete()
}
